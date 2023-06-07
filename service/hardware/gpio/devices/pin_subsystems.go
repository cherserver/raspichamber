package devices

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel/pins"
)

type PinSubsystems struct {
	NativePinSubsystem   lowlevel.PinSubsystem
	ExternalPinSubsystem lowlevel.PinSubsystem
}

func NewPinSubsystems() *PinSubsystems {
	return &PinSubsystems{
		NativePinSubsystem:   pins.NewRPIOPinSubsystem(),
		ExternalPinSubsystem: pins.NewPiBlasterPinSubsystem(),
	}
}

func (s *PinSubsystems) Init() error {
	err := s.ExternalPinSubsystem.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize external pin subsystem: %w", err)
	}

	err = s.NativePinSubsystem.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize native pin subsystem: %w", err)
	}

	return nil
}

func (s *PinSubsystems) Stop() {
	s.NativePinSubsystem.Stop()
	s.ExternalPinSubsystem.Stop()
}
