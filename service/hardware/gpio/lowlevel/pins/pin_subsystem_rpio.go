package pins

import (
	"log"

	"github.com/stianeikeland/go-rpio"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
)

var subsystemTypeRPIO = subsystemTypeRPIOType{}

type subsystemTypeRPIOType struct {
}

type pinRpioSubsystem struct {
	isInitialized bool
}

var _ lowlevel.PinSubsystem = &pinRpioSubsystem{}

func (f *pinRpioSubsystem) IsInitialized() bool {
	return f.isInitialized
}

func (f *pinRpioSubsystem) Type() interface{} {
	return subsystemTypeRPIO
}

func NewRPIOPinSubsystem() *pinRpioSubsystem {
	return &pinRpioSubsystem{
		isInitialized: false,
	}
}

func (f *pinRpioSubsystem) Init() error {
	err := rpio.Open()
	if err != nil {
		log.Printf("RPIO subsystem init error: %v", err)
		return err
	}

	f.isInitialized = true

	return nil
}

func (f *pinRpioSubsystem) Stop() {
	err := rpio.Close()

	if err != nil {
		log.Printf("RPIO subsystem stop error: %v", err)
	}
}

func requireIsInitializedRPIO(subsystem lowlevel.PinSubsystem, pin lowlevel.Pin) {
	if subsystem.Type() != subsystemTypeRPIO {
		log.Panicf("RPIO pin subsystem is required to initialize pin '%v'", pin)
	}

	if !subsystem.IsInitialized() {
		log.Panicf("Try to initialize pin '%v' before RPIO pin subsystem", pin)
	}
}
