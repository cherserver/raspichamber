package devices

import (
	"fmt"
	"sync"

	"github.com/cherserver/raspichamber/service/hardware/assert"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel/pins"
	"github.com/cherserver/raspichamber/service/software"
)

var _ software.DryerHatch = &DryerHatch{}

const (
	angle0PWMValue  = 15 // lower
	angle90PWMValue = 58 // upper
	angleStep       = float32(angle90PWMValue-angle0PWMValue) / 90
)

func angleToPWM(angle uint8) uint8 {
	return angle0PWMValue + uint8(float32(angle)*angleStep)
}

type DryerHatch struct {
	pin lowlevel.PwmPin

	angle      uint8
	angleMutex sync.RWMutex
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

	s.angleMutex.Lock()
	defer s.angleMutex.Unlock()

	err := s.pin.SetPwmPercent(angleToPWM(angle))
	if err != nil {
		return err
	}

	s.angle = angle

	return nil
}

func (s *DryerHatch) Angle() uint8 {
	s.angleMutex.RLock()
	defer s.angleMutex.RUnlock()

	return s.angle
}
