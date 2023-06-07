package devices

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/uart"
	"github.com/cherserver/raspichamber/service/software"
)

var _ software.Fan = &Fan{}
var _ uart.StatusDataProcessor = &Fan{}

type Fan struct {
	uart *uart.UART

	name         string
	speedPercent uint8
	rpm          uint32
}

func NewFan(name string, uart *uart.UART) *Fan {
	return &Fan{
		uart: uart,

		name:         name,
		speedPercent: 0,
		rpm:          0,
	}
}

func (f *Fan) Name() string {
	return f.name
}

func (f *Fan) Register() error {
	return f.uart.RegisterStatusDataProcessor(f)
}

func (f *Fan) RPM() uint32 {
	return f.rpm
}

func (f *Fan) SetSpeedPercent(value uint8) error {
	err := f.uart.Send(fmt.Sprintf("set-fan-power -name %s -percent %d", f.name, value))
	if err != nil {
		return err
	}

	f.speedPercent = value

	return nil
}

func (f *Fan) SpeedPercent() uint8 {
	return f.speedPercent
}

func (f *Fan) ProcessStatusData(data map[string]interface{}) {
	for key, val := range data {
		switch key {
		case "p":
			floatVal, ok := val.(float64)
			if ok {
				f.speedPercent = uint8(floatVal)
			}
		case "rpm":
			floatVal, ok := val.(float64)
			if ok {
				f.rpm = uint32(floatVal)
			}
		}
	}
}
