package fan

import (
	"log"

	"github.com/stianeikeland/go-rpio"
)

const (
	// TODO: move to config
	fanPin          = 12
	fanSpeedPercent = 25

	pwmFreq     = 25000 // 25kHz for fan pm pin
	pwmCycleLen = 100   // so dutyLen will be percent
)

type fan struct {
	pin rpio.Pin
}

func New() *fan {
	return &fan{
		pin: rpio.Pin(fanPin),
	}
}

func (f *fan) Init() error {
	err := rpio.Open()
	if err != nil {
		log.Printf("PWM pin init error: %v", err)
		return err
	}

	f.pin.Pwm()
	f.pin.Freq(pwmFreq * pwmCycleLen)
	f.pin.DutyCycle(fanSpeedPercent, pwmCycleLen)

	return nil
}

func (f *fan) Stop() {
	err := rpio.Close()

	if err != nil {
		log.Printf("PWM pin stop error: %v", err)
	}
}
