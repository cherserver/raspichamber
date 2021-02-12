package pins

import (
	"github.com/stianeikeland/go-rpio"

	"github.com/cherserver/raspichamber/service/hardware/assert"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
)

const (
	pwmFreq     = 25000 // 25kHz for fan pwm pin
	pwmCycleLen = 100   // so dutyLen will be percent
)

type pwmHardwarePin struct {
	pinSubsystem lowlevel.PinSubsystem
	pin          lowlevel.Pin
	hwPin        rpio.Pin
}

var _ lowlevel.PwmPin = &pwmHardwarePin{}

func NewPWMHardwarePin(pinSubsystem lowlevel.PinSubsystem, pin lowlevel.Pin) *pwmHardwarePin {
	return &pwmHardwarePin{
		pinSubsystem: pinSubsystem,
		pin:          pin,
		hwPin:        rpio.Pin(pin.GPIOIndex()),
	}
}

func (f *pwmHardwarePin) Init() error {
	requireIsInitializedRPIO(f.pinSubsystem, f.pin)

	f.hwPin.Pwm()
	f.hwPin.Freq(pwmFreq * pwmCycleLen)
	f.hwPin.DutyCycle(0, pwmCycleLen) // set 0 'percent'

	return nil
}

func (f *pwmHardwarePin) Stop() {}

func (f *pwmHardwarePin) SetPwmPercent(value uint8) error {
	if err := assert.IsPercentFrom0To100(value); err != nil {
		return err
	}

	f.hwPin.DutyCycle(uint32(value), pwmCycleLen)
	return nil
}
