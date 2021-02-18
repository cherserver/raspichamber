package devices

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/assert"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel/pins"
	"github.com/cherserver/raspichamber/service/software"
)

var _ software.DryerHatch = &DryerHatch{}

type DryerHatch struct {
	pin lowlevel.PwmPin
}

func NewDryerServo(subsystems *PinSubsystems) *DryerHatch {
	return &DryerHatch{
		pin: pins.NewPWMSoftwarePin(subsystems.ExternalPinSubsystem, lowlevel.DryerServoPin),
	}
}

func (s *DryerHatch) Init() error {
	err := s.pin.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize 'dryer servo': %w", err)
	}

	return nil
}

func (s *DryerHatch) Stop() {
	s.pin.Stop()
}

func (s *DryerHatch) SetAngle(angle uint8) error {
	if err := assert.IsAngleFrom0To90(angle); err != nil {
		return err
	}

	return s.pin.SetPwmPercent(angle)
}
