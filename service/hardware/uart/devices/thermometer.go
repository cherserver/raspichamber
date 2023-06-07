package devices

import (
	"github.com/cherserver/raspichamber/service/hardware/uart"
	"github.com/cherserver/raspichamber/service/software"
)

var _ software.Thermometer = &Thermometer{}
var _ uart.StatusDataProcessor = &Thermometer{}

type Thermometer struct {
	uart *uart.UART

	name        string
	temperature float32
	humidity    float32
}

func NewThermometer(name string, uart *uart.UART) *Thermometer {
	return &Thermometer{
		uart: uart,

		name:        name,
		temperature: 0,
		humidity:    0,
	}
}

func (t *Thermometer) Name() string {
	return t.name
}

func (t *Thermometer) Register() error {
	return t.uart.RegisterStatusDataProcessor(t)
}

func (t *Thermometer) Temperature() float32 {
	return t.temperature
}

func (t *Thermometer) Humidity() float32 {
	return t.humidity
}

func (t *Thermometer) ProcessStatusData(data map[string]interface{}) {
	for key, val := range data {
		switch key {
		case "hum":
			floatVal, ok := val.(float64)
			if ok {
				t.humidity = float32(floatVal)
			}
		case "temp":
			floatVal, ok := val.(float64)
			if ok {
				t.temperature = float32(floatVal)
			}
		}
	}
}
