package pins

import (
	"github.com/stianeikeland/go-rpio"

	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
)

type rpmSensorPin struct {
	pinSubsystem lowlevel.PinSubsystem
	pin          lowlevel.Pin
	hwPin        rpio.Pin
}

var _ lowlevel.RpmSensorPin = &rpmSensorPin{}

func NewRPMSensorPin(pinSubsystem lowlevel.PinSubsystem, pin lowlevel.Pin) *rpmSensorPin {
	return &rpmSensorPin{
		pinSubsystem: pinSubsystem,
		pin:          pin,
		hwPin:        rpio.Pin(pin.GPIOIndex()),
	}
}

func (f *rpmSensorPin) Init() error {
	requireIsInitializedRPIO(f.pinSubsystem, f.pin)

	f.hwPin.Input()
	f.hwPin.PullDown()
	f.hwPin.Detect(rpio.NoEdge)

	return nil
}

func (f *rpmSensorPin) Stop() {}

func (f *rpmSensorPin) RPM() (uint32, error) {
	return 0, nil
}
