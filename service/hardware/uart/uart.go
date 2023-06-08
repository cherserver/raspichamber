package uart

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"go.bug.st/serial"
)

type StatusDataProcessor interface {
	Name() string
	ProcessStatusData(map[string]interface{})
}

type UART struct {
	portName string
	baudRate int

	port serial.Port

	receivers map[string]StatusDataProcessor
	sendMutex sync.Mutex
}

func New(port string, baudRate int) *UART {
	return &UART{
		portName: port,
		baudRate: baudRate,
		port:     nil,

		receivers: make(map[string]StatusDataProcessor),
	}
}

func (u *UART) Init() error {
	var err error
	u.port, err = serial.Open(u.portName, &serial.Mode{
		BaudRate:          u.baudRate,
		DataBits:          8,
		Parity:            serial.NoParity,
		StopBits:          serial.OneStopBit,
		InitialStatusBits: nil,
	})
	if err != nil {
		return fmt.Errorf("failed to open serial port '%s': %w", u.portName, err)
	}

	_ = u.port.ResetInputBuffer()
	_ = u.port.ResetOutputBuffer()

	err = u.Send("set-auto-report on")
	if err != nil {
		return fmt.Errorf("failed to set pico autoreport 'on': %w", err)
	}

	go u.reader()

	return nil
}

func (u *UART) RegisterStatusDataProcessor(processor StatusDataProcessor) error {
	_, fnd := u.receivers[processor.Name()]
	if fnd {
		return fmt.Errorf("failed to register '%s': name is already in use", processor.Name())
	}

	u.receivers[processor.Name()] = processor
	return nil
}

func (u *UART) Send(data string) error {
	u.sendMutex.Lock()
	defer u.sendMutex.Unlock()

	if len(data) == 0 {
		return nil
	}

	if data[len(data)-1] != '\n' {
		data += "\n"
	}

	_, err := u.port.Write([]byte(data))
	if err != nil {
		return fmt.Errorf("failed to write data to UART: %w", err)
	}

	return nil
}

func (u *UART) reader() {
	reader := bufio.NewReader(u.port)
	for {
		readStr, err := reader.ReadString('\n')
		if err != nil {
			//TODO: restore state?
			log.Fatalf("Error reading data from UART: %v", err)
		}

		go u.processOutput(readStr)
	}
}

func (u *UART) processOutput(data string) {
	if len(data) == 0 {
		return
	}

	if data[0] == '{' {
		u.processStatus(data)
		return
	}

	u.processResponse(data)
}

func (u *UART) processResponse(data string) {
	log.Printf("Response: %s", data)
}

func (u *UART) processStatus(data string) {
	message := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(data), &message)

	if err != nil {
		log.Printf("Failed to unmarshall status message: %v. Msg: '%s'", err, data)
		return
	}

	for name, status := range message {
		if status == nil {
			continue
		}

		statusStorage, ok := status.(map[string]interface{})
		if !ok {
			continue
		}

		receiver, fnd := u.receivers[name]
		if !fnd {
			continue
		}

		receiver.ProcessStatusData(statusStorage)
	}
}

func (u *UART) Stop() {
	if u.port == nil {
		return
	}

	_ = u.Send("set-auto-report off")

	_ = u.port.Close()
}
