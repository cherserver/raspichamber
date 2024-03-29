package devices

import (
	"github.com/cristalhq/atomix"
	"github.com/d2r2/go-logger"

	"github.com/d2r2/go-dht"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
	"github.com/cherserver/raspichamber/service/software"
)

const (
	numRetries = 10
)

var _ software.Thermometer = &Thermometer{}

func NewInnerThermometer(subsystems *PinSubsystems) *Thermometer {
	return newThermometer(subsystems, lowlevel.InnerTempPin)
}

func NewOuterThermometer(subsystems *PinSubsystems) *Thermometer {
	return newThermometer(subsystems, lowlevel.OuterTempPin)
}

func NewDryerThermometer(subsystems *PinSubsystems) *Thermometer {
	return newThermometer(subsystems, lowlevel.DryerTempPin)
}

type Thermometer struct {
	pin lowlevel.Pin

	temperature *atomix.Float32
	humidity    *atomix.Float32
}

func newThermometer(subsystems *PinSubsystems, pin lowlevel.Pin) *Thermometer {
	_ = subsystems // not used
	return &Thermometer{
		pin:         pin,
		temperature: atomix.NewFloat32(0),
		humidity:    atomix.NewFloat32(0),
	}
}

func (t *Thermometer) Init() error {
	_ = logger.ChangePackageLogLevel("dht", logger.ErrorLevel)
	go t.workCycle()
	return nil
}

func (t *Thermometer) workCycle() {
	// TODO: exit after experiment end
	for {
		temperature, humidity, _, err :=
			dht.ReadDHTxxWithRetry(dht.DHT22, t.pin.GPIOIndex(), false, numRetries)

		_ = err

		t.temperature.Store(temperature)
		t.humidity.Store(humidity)
	}
}

func (t *Thermometer) Stop() {}

func (t *Thermometer) Temperature() float32 {
	return t.temperature.Load()
}

func (t *Thermometer) Humidity() float32 {
	return t.humidity.Load()
}
