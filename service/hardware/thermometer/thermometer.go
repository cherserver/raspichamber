package thermometer

import (
	"log"

	"github.com/cristalhq/atomix"
	"github.com/d2r2/go-logger"

	"github.com/d2r2/go-dht"
)

const (
	numRetries = 10
)

type thermometer struct {
	pin int

	temperature *atomix.Float32
	humidity    *atomix.Float32
}

func New(pin int) *thermometer {
	return &thermometer{
		pin:         pin,
		temperature: atomix.NewFloat32(0),
		humidity:    atomix.NewFloat32(0),
	}
}

func (t *thermometer) Init() error {
	_ = logger.ChangePackageLogLevel("dht", logger.ErrorLevel)
	go t.workCycle()
	return nil
}

func (t *thermometer) workCycle() {
	// TODO: exit after experiment end
	for {
		log.Printf("DHT read cycle: %v", t.pin)
		temperature, humidity, _, err :=
			dht.ReadDHTxxWithRetry(dht.DHT22, t.pin, false, numRetries)

		if err != nil {
			log.Printf("DHT read error: %v", err)
		}

		t.temperature.Store(temperature)
		t.humidity.Store(humidity)

		log.Printf("DHT read temperature: %v, humidity: %v", temperature, humidity)
	}
}

func (t *thermometer) Stop() {}

func (t *thermometer) Temperature() float32 {
	return t.temperature.Load()
}

func (t *thermometer) Humidity() float32 {
	return t.humidity.Load()
}
