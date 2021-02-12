package devices

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel/pins"
)

func NewOuterFan(subsystems *PinSubsystems) *Fan {
	return newFan(
		pins.NewPWMHardwarePin(subsystems.NativePinSubsystem, lowlevel.OuterFanPwmPin),
		pins.NewRPMSensorPin(subsystems.NativePinSubsystem, lowlevel.OuterFanTachoPin),
	)
}

func NewInnerFan(subsystems *PinSubsystems) *Fan {
	return newFan(
		pins.NewPWMHardwarePin(subsystems.NativePinSubsystem, lowlevel.InnerFanPwmPin),
		pins.NewRPMSensorPin(subsystems.NativePinSubsystem, lowlevel.InnerFanTachoPin),
	)
}

func NewRPiFan(subsystems *PinSubsystems) *Fan {
	return newFan(
		pins.NewPWMSoftwarePin(subsystems.ExternalPinSubsystem, lowlevel.RPiFanPwmPin),
		pins.NewRPMSensorPin(subsystems.NativePinSubsystem, lowlevel.RPiFanTachoPin),
	)
}

type Fan struct {
	pwmPin   lowlevel.PwmPin
	tachoPin lowlevel.RpmSensorPin
}

func newFan(
	pwmPin lowlevel.PwmPin,
	tachoPin lowlevel.RpmSensorPin) *Fan {
	return &Fan{
		pwmPin:   pwmPin,
		tachoPin: tachoPin,
	}
}

func (f *Fan) Init() error {
	err := f.pwmPin.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize fan pwm pin: %w", err)
	}

	err = f.tachoPin.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize fan tacho pin: %w", err)
	}

	return nil
}

func (f *Fan) Stop() {
	f.tachoPin.Stop()
	f.pwmPin.Stop()
}
