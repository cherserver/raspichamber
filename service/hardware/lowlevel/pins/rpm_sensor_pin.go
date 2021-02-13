package pins

import (
	"fmt"
	"log"
	"time"

	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"

	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
)

type rpmSensorPin struct {
	pinSubsystem lowlevel.PinSubsystem
	pin          lowlevel.Pin
	hwPin        int
	chip         *gpiod.Chip
	line         *gpiod.Line
}

var _ lowlevel.RpmSensorPin = &rpmSensorPin{}

func NewRPMSensorPin(pinSubsystem lowlevel.PinSubsystem, pin lowlevel.Pin) *rpmSensorPin {
	return &rpmSensorPin{
		pinSubsystem: pinSubsystem,
		pin:          pin,
		hwPin:        pin.GPIOIndex(),
	}
}

func (f *rpmSensorPin) Init() error {
	// requireIsInitializedRPIO(f.pinSubsystem, f.pin)
	var err error
	f.chip, err = gpiod.NewChip("gpiochip0")
	if err != nil {
		return fmt.Errorf("failed to initialize pin '%v', failed to init GPIOD chip: %w", f.pin, err)
	}

	offset := rpi.J8p7
	f.line, err = f.chip.RequestLine(offset,
		gpiod.WithPullUp,
		gpiod.WithRisingEdge,
		gpiod.WithEventHandler(f.edgeEventHandler))
	if err != nil {
		return fmt.Errorf("failed to initialize pin '%v', failed to init GPIOD line: %w", f.pin, err)
	}

	return nil
}

func (f *rpmSensorPin) Stop() {
	err := f.line.Close()

	if err != nil {
		log.Printf("GPIO line close error: %v", err)
	}

	err = f.chip.Close()

	if err != nil {
		log.Printf("GPIO chip close error: %v", err)
	}
}

func (f *rpmSensorPin) edgeEventHandler(evt gpiod.LineEvent) {
	t := time.Now()
	edge := "rising"
	if evt.Type == gpiod.LineEventFallingEdge {
		edge = "falling"
	}

	fmt.Printf("event:%3d %-7s %s (%s)\n",
		evt.Offset,
		edge,
		t.Format(time.RFC3339Nano),
		evt.Timestamp)
}

func (f *rpmSensorPin) RPM() (uint32, error) {
	return 0, nil
}
