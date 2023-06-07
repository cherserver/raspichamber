package pins

import (
	"github.com/stianeikeland/go-rpio"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
)

type switchPin struct {
	pinSubsystem lowlevel.PinSubsystem
	pin          lowlevel.Pin
	hwPin        rpio.Pin
}

var _ lowlevel.SwitchPin = &switchPin{}

func NewSwitchPin(pinSubsystem lowlevel.PinSubsystem, pin lowlevel.Pin) *switchPin {
	return &switchPin{
		pinSubsystem: pinSubsystem,
		pin:          pin,
		hwPin:        rpio.Pin(pin.GPIOIndex()),
	}
}

func (f *switchPin) Init() error {
	requireIsInitializedRPIO(f.pinSubsystem, f.pin)

	f.hwPin.Output()
	f.hwPin.Low()

	return nil
}

func (f *switchPin) Stop() {}

func (f *switchPin) RPM() (uint32, error) {
	return 0, nil
}

func (f *switchPin) TurnOn() {
	f.hwPin.High()
}

func (f *switchPin) TurnOff() {
	f.hwPin.Low()
}
