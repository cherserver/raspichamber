package fan

import (
	"log"

	"github.com/stianeikeland/go-rpio"
)

const (
	// TODO: move to config
	fanSpeedPercent = 10

	pwmFreq     = 25000 // 25kHz for fan pm pin
	pwmCycleLen = 100   // so dutyLen will be percent
)

type fan struct {
	pwmPin        rpio.Pin
	tachometerPin rpio.Pin
}

func New(pwmPin rpio.Pin, tachometerPin rpio.Pin) *fan {
	return &fan{
		pwmPin:        pwmPin,
		tachometerPin: tachometerPin,
	}
}

func (f *fan) Init() error {
	err := rpio.Open()
	if err != nil {
		log.Printf("PWM pin init error: %v", err)
		return err
	}

	f.pwmPin.Pwm()
	f.pwmPin.Freq(pwmFreq * pwmCycleLen)
	f.pwmPin.DutyCycle(fanSpeedPercent, pwmCycleLen)

	return nil
}

func (f *fan) Stop() {
	err := rpio.Close()

	if err != nil {
		log.Printf("PWM pin stop error: %v", err)
	}
}

func (f *fan) SetSpeedPercent(percent uint32) {
	f.pwmPin.DutyCycle(percent, pwmCycleLen)
}

func (f *fan) RPM() (uint32, error) {
	return 0, nil
}
