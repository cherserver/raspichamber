package pins

import (
	"log"
	"os"

	"github.com/stianeikeland/go-rpio"

	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
)

var subsystemTypePiBlaster = subsystemTypePiBlasterType{}

type subsystemTypePiBlasterType struct {
}

type pinPiBlasterSubsystem struct {
	isInitialized bool
}

var _ lowlevel.PinSubsystem = &pinPiBlasterSubsystem{}

func (f *pinPiBlasterSubsystem) IsInitialized() bool {
	return f.isInitialized
}

func (f *pinPiBlasterSubsystem) Type() interface{} {
	return subsystemTypeRPIO
}

func NewPiBlasterPinSubsystem() *pinPiBlasterSubsystem {
	return &pinPiBlasterSubsystem{
		isInitialized: false,
	}
}

func (f *pinPiBlasterSubsystem) Init() error {
	file, err := os.OpenFile("/dev/pi-blaster", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	f.isInitialized = true

	return nil
}

func (f *pinPiBlasterSubsystem) Stop() {
	err := rpio.Close()

	if err != nil {
		log.Printf("PWM pin stop error: %v", err)
	}
}

func requireIsInitializedPiBlaster(subsystem lowlevel.PinSubsystem, pin lowlevel.Pin) {
	if subsystem.Type() != subsystemTypePiBlaster {
		log.Panicf("pi-blaster subsystem is required to initialize pin '%v'", pin)
	}

	if !subsystem.IsInitialized() {
		log.Panicf("Pin '%v' is initialized before pi-blaster subsystem", pin)
	}
}
