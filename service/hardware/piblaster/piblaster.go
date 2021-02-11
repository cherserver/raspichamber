package piblaster

import (
	"fmt"
	"os"

	"github.com/cherserver/raspichamber/service/hardware"
)

type piBlaster struct {
}

var _ hardware.ExternalPwmSubsystem = &piBlaster{}

func New() *piBlaster {
	return &piBlaster{}
}

func (p *piBlaster) Init() error {
	file, err := os.OpenFile("/dev/pi-blaster", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	return nil
}

func (p *piBlaster) Stop() {}

func (p *piBlaster) SetPinPwmPercent(pin int, value uint8) error {
	pinVal := 1 / float32(value)
	data := fmt.Sprintf("%v=%v\n", pin, pinVal)
	file, err := os.OpenFile("/dev/pi-blaster", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
