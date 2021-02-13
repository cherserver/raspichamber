package lowlevel

import (
	"fmt"

	"github.com/warthog618/gpiod/device/rpi"
)

const (
	InnerFanPwmGPIO   = 12
	InnerFanTachoGPIO = 17

	OuterFanPwmGPIO   = 13
	OuterFanTachoGPIO = 4

	RPiFanPwmGPIO   = 9
	RPiFanTachoGPIO = 1

	OuterTempGPIO = 23
	InnerTempGPIO = 22
	DryerTempGPIO = 2

	DryerServoGPIO = 3

	HeaterButton1GPIO = 26
	HeaterButton2GPIO = 16
	HeaterButton3GPIO = 20
	HeaterButton4GPIO = 21
)

type pin struct {
	gpioIndex int
	j8Index   int
	name      string
}

var _ Pin = pin{}

func (p pin) GPIOIndex() int {
	return p.gpioIndex
}

func (p pin) J8Index() int {
	return p.j8Index
}

func (p pin) Name() string {
	return p.name
}

func (p pin) String() string {
	return fmt.Sprintf("%s(GPIO%v)", p.name, p.gpioIndex)
}

var (
	InnerFanPwmPin = pin{
		gpioIndex: InnerFanPwmGPIO,
		name:      "internal_fan_pwm_pin",
	}

	InnerFanTachoPin = pin{
		gpioIndex: InnerFanTachoGPIO,
		j8Index:   rpi.J8p11,
		name:      "internal_fan_tacho_pin",
	}

	OuterFanPwmPin = pin{
		gpioIndex: OuterFanPwmGPIO,
		name:      "external_fan_pwm_pin",
	}

	OuterFanTachoPin = pin{
		gpioIndex: OuterFanTachoGPIO,
		j8Index:   rpi.J8p7,
		name:      "external_fan_tacho_pin",
	}

	RPiFanPwmPin = pin{
		gpioIndex: RPiFanPwmGPIO,
		name:      "rpi_fan_pwm_pin",
	}

	RPiFanTachoPin = pin{
		gpioIndex: RPiFanTachoGPIO,
		j8Index:   rpi.J8p28,
		name:      "rpi_fan_tacho_pin",
	}

	OuterTempPin = pin{
		gpioIndex: OuterTempGPIO,
		name:      "outer_temp_pin",
	}

	InnerTempPin = pin{
		gpioIndex: InnerTempGPIO,
		name:      "inner_temp_pin",
	}

	DryerTempPin = pin{
		gpioIndex: DryerTempGPIO,
		name:      "dryer_temp_pin",
	}

	DryerServoPin = pin{
		gpioIndex: DryerServoGPIO,
		name:      "dryer_servo_pin",
	}

	HeaterButton1Pin = pin{
		gpioIndex: HeaterButton1GPIO,
		name:      "heater_button_1_pin",
	}

	HeaterButton2Pin = pin{
		gpioIndex: HeaterButton2GPIO,
		name:      "heater_button_2_pin",
	}

	HeaterButton3Pin = pin{
		gpioIndex: HeaterButton3GPIO,
		name:      "heater_button_3_pin",
	}

	HeaterButton4Pin = pin{
		gpioIndex: HeaterButton4GPIO,
		name:      "heater_button_4_pin",
	}
)
