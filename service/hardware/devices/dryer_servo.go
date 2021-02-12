package devices

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel/pins"
)

type DryerServo struct {
	pin lowlevel.PwmPin
}

func NewDryerServo(subsystems *PinSubsystems) *DryerServo {
	return &DryerServo{
		pin: pins.NewPWMSoftwarePin(subsystems.ExternalPinSubsystem, lowlevel.DryerServoPin),
	}
}

func (s *DryerServo) Init() error {
	err := s.pin.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize 'dryer servo': %w", err)
	}

	return nil
}

func (s *DryerServo) Stop() {
	s.pin.Stop()
}

func (s *DryerServo) SetAngle(angle uint8) error {
	return s.pin.SetPwmPercent(angle)
}
