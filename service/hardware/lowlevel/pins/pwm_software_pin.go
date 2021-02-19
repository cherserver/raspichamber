package pins

import (
	"fmt"
	"os"

	"github.com/cherserver/raspichamber/service/hardware/assert"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
)

type pwmSoftwarePin struct {
	subsystem lowlevel.PinSubsystem
	pin       lowlevel.Pin
}

var _ lowlevel.PwmPin = &pwmSoftwarePin{}

func NewPWMSoftwarePin(subsystem lowlevel.PinSubsystem, pin lowlevel.Pin) *pwmSoftwarePin {
	return &pwmSoftwarePin{
		subsystem: subsystem,
		pin:       pin,
	}
}

func (p *pwmSoftwarePin) Init() error {
	requireIsInitializedPiBlaster(p.subsystem, p.pin)

	return nil
}

func (p *pwmSoftwarePin) Stop() {}

func (p *pwmSoftwarePin) SetPwmPercent(value uint8) error {
	if err := assert.IsPercentFrom0To100(value); err != nil {
		return err
	}

	pinVal := float32(value) / 100
	data := fmt.Sprintf("%v=%v\n", p.pin.GPIOIndex(), pinVal)
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
