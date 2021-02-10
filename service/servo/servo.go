package servo

import (
	"log"
	"sync/atomic"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type servo struct {
	pin           string
	robot         *gobot.Robot
	requiredAngle uint32
}

func New(pin string) *servo {
	return &servo{
		pin:           pin,
		robot:         nil,
		requiredAngle: 30,
	}
}

func (s *servo) Init() error {
	r := raspi.NewAdaptor()
	hardwareServo := gpio.NewServoDriver(r, s.pin)

	work := func() {
		gobot.Every(1*time.Second, func() {
			angle := atomic.LoadUint32(&s.requiredAngle)
			err := hardwareServo.Move(uint8(angle))
			if err != nil {
				log.Printf("Servo cycle error: %v", err)
			}
		})
	}

	s.robot = gobot.NewRobot("servo",
		[]gobot.Connection{r},
		[]gobot.Device{hardwareServo},
		work,
	)

	return s.robot.Start()
}

func (s *servo) Stop() {
	err := s.robot.Stop()
	if err != nil {
		log.Printf("Servo stop error: %v", err)
	}
}

func (s *servo) SetAngle(angle uint8) {
	atomic.StoreUint32(&s.requiredAngle, uint32(angle))
}
