package pins

import (
	"log"
	"os"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
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
	return subsystemTypePiBlaster
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

func (f *pinPiBlasterSubsystem) Stop() {}

func requireIsInitializedPiBlaster(subsystem lowlevel.PinSubsystem, pin lowlevel.Pin) {
	if subsystem.Type() != subsystemTypePiBlaster {
		log.Panicf("pi-blaster subsystem is required to initialize pin '%v'", pin)
	}

	if !subsystem.IsInitialized() {
		log.Panicf("Try to initialize pin '%v' before pi-blaster subsystem", pin)
	}
}
