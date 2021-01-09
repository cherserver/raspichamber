package servo

import (
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

// TODO: move to config
const servoPin = "7"

type servo struct {
	robot *gobot.Robot
}

func New() *servo {
	return &servo{
		robot: nil,
	}
}

func (s *servo) Init() error {
	r := raspi.NewAdaptor()
	hardwareServo := gpio.NewServoDriver(r, servoPin)

	work := func() {
		gobot.Every(1*time.Second, func() {
			err := hardwareServo.Center()
			if err != nil {
				log.Printf("Servo error: %v", err)
			}
		})
	}

	robot := gobot.NewRobot("servo",
		[]gobot.Connection{r},
		[]gobot.Device{hardwareServo},
		work,
	)

	return robot.Start()
}

func (s *servo) Stop() {

}
