package lowlevel

import (
	"fmt"
)

const (
	// Inner fan
	InnerFanPwmGPIO = 12
	InnerFanPwmNum  = 32

	InnerFanTachoGPIO = 17
	InnerFanTachoNum  = 11

	// Outer fan
	OuterFanPwmGPIO = 13
	OuterFanPwmNum  = 33

	OuterFanTachoGPIO = 4
	OuterFanTachoNum  = 7

	// RPi fan
	RPiFanPwmGPIO = 9
	RPiFanPwmNum  = 21

	RPiFanTachoGPIO = 1
	RPiFanTachoNum  = 28

	// Temp
	OuterTempGPIO = 23
	OuterTempNum  = 16

	InnerTempGPIO = 22
	InnerTempNum  = 15

	DryerTempGPIO = 2
	DryerTempNum  = 3

	// Servo
	DryerServoGPIO = 3
	DryerServoNum  = 5

	// Heater buttons
	HeaterButton1GPIO = 26
	HeaterButton1Num  = 37

	HeaterButton2GPIO = 16
	HeaterButton2Num  = 36

	HeaterButton3GPIO = 20
	HeaterButton3Num  = 38

	HeaterButton4GPIO = 21
	HeaterButton4Num  = 40
)

type pin struct {
	name      string
	j8Index   int
	gpioIndex int
}

var _ Pin = pin{}

func (p pin) Name() string {
	return p.name
}

func (p pin) J8Index() int {
	return p.j8Index
}

func (p pin) GPIOIndex() int {
	return p.gpioIndex
}

func (p pin) String() string {
	return fmt.Sprintf("%s(#%v-GPIO%v)", p.name, p.j8Index, p.gpioIndex)
}

var (
	InnerFanPwmPin = pin{
		j8Index:   InnerFanPwmNum,
		gpioIndex: InnerFanPwmGPIO,
		name:      "internal_fan_pwm_pin",
	}

	InnerFanTachoPin = pin{
		j8Index:   InnerFanTachoNum,
		gpioIndex: InnerFanTachoGPIO,
		name:      "internal_fan_tacho_pin",
	}

	OuterFanPwmPin = pin{
		j8Index:   OuterFanPwmNum,
		gpioIndex: OuterFanPwmGPIO,
		name:      "external_fan_pwm_pin",
	}

	OuterFanTachoPin = pin{
		j8Index:   OuterFanTachoNum,
		gpioIndex: OuterFanTachoGPIO,
		name:      "external_fan_tacho_pin",
	}

	RPiFanPwmPin = pin{
		j8Index:   RPiFanPwmNum,
		gpioIndex: RPiFanPwmGPIO,
		name:      "rpi_fan_pwm_pin",
	}

	RPiFanTachoPin = pin{
		j8Index:   RPiFanTachoNum,
		gpioIndex: RPiFanTachoGPIO,
		name:      "rpi_fan_tacho_pin",
	}

	OuterTempPin = pin{
		j8Index:   OuterTempNum,
		gpioIndex: OuterTempGPIO,
		name:      "outer_temp_pin",
	}

	InnerTempPin = pin{
		j8Index:   InnerTempNum,
		gpioIndex: InnerTempGPIO,
		name:      "inner_temp_pin",
	}

	DryerTempPin = pin{
		j8Index:   DryerTempNum,
		gpioIndex: DryerTempGPIO,
		name:      "dryer_temp_pin",
	}

	DryerServoPin = pin{
		j8Index:   DryerServoNum,
		gpioIndex: DryerServoGPIO,
		name:      "dryer_servo_pin",
	}

	HeaterButton1Pin = pin{
		j8Index:   HeaterButton1Num,
		gpioIndex: HeaterButton1GPIO,
		name:      "heater_button_1_pin",
	}

	HeaterButton2Pin = pin{
		j8Index:   HeaterButton2Num,
		gpioIndex: HeaterButton2GPIO,
		name:      "heater_button_2_pin",
	}

	HeaterButton3Pin = pin{
		j8Index:   HeaterButton3Num,
		gpioIndex: HeaterButton3GPIO,
		name:      "heater_button_3_pin",
	}

	HeaterButton4Pin = pin{
		j8Index:   HeaterButton4Num,
		gpioIndex: HeaterButton4GPIO,
		name:      "heater_button_4_pin",
	}
)
