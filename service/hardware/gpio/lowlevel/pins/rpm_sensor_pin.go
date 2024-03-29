package pins

import (
	"fmt"
	"log"
	"time"

	"github.com/paulbellamy/ratecounter"
	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
)

const (
	secondsPerMinute   = 60
	ticksPerRevolution = 2
)

type rpmSensorPin struct {
	pinSubsystem lowlevel.PinSubsystem
	pin          lowlevel.Pin
	hwPin        int
	chip         *gpiod.Chip
	line         *gpiod.Line
	counter      *ratecounter.RateCounter
	lastTime     time.Duration
}

var _ lowlevel.RpmSensorPin = &rpmSensorPin{}

func NewRPMSensorPin(pinSubsystem lowlevel.PinSubsystem, pin lowlevel.Pin) *rpmSensorPin {
	return &rpmSensorPin{
		pinSubsystem: pinSubsystem,
		pin:          pin,
		hwPin:        rpi.MustPin(fmt.Sprintf("j8p%v", pin.J8Index())),

		counter: ratecounter.NewRateCounter(1 * time.Second),
	}
}

func (f *rpmSensorPin) Init() error {
	var err error
	f.chip, err = gpiod.NewChip("gpiochip0")
	if err != nil {
		return fmt.Errorf("failed to initialize pin '%v', failed to init GPIOD chip: %w", f.pin, err)
	}

	f.line, err = f.chip.RequestLine(f.hwPin,
		gpiod.WithPullUp,
		gpiod.WithFallingEdge,
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
	go f.countEvent(evt.Timestamp)
}

func (f *rpmSensorPin) countEvent(eventTime time.Duration) {
	currentDuration := eventTime - f.lastTime
	f.lastTime = eventTime
	if currentDuration < 50*time.Microsecond {
		return
	}

	f.counter.Incr(1)
	/*log.Printf("event:%3d %-7s %s (%s)\n",
	evt.Offset,
	edge,
	t.Format(time.RFC3339Nano),
	evt.Timestamp)*/
}

func (f *rpmSensorPin) RPM() (uint32, error) {
	herz := f.counter.Rate() / 2
	return uint32(herz * secondsPerMinute / ticksPerRevolution), nil
}
