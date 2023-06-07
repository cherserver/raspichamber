package devices

import (
	"fmt"
	"sync"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel/pins"
	"github.com/cherserver/raspichamber/service/software"
)

var _ software.Fan = &Fan{}

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

	speed      uint8
	speedMutex sync.RWMutex
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

func (f *Fan) RPM() uint32 {
	val, err := f.tachoPin.RPM()
	if err != nil {
		return 0
	}

	return val
}

func (f *Fan) SetSpeedPercent(value uint8) error {
	f.speedMutex.Lock()
	defer f.speedMutex.Unlock()

	err := f.pwmPin.SetPwmPercent(value)
	if err != nil {
		return err
	}

	f.speed = value

	return nil
}

func (f *Fan) SpeedPercent() uint8 {
	f.speedMutex.RLock()
	defer f.speedMutex.RUnlock()

	return f.speed
}
