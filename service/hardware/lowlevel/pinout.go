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
	RPiFanPwmGPIO = 7
	RPiFanPwmNum  = 26

	RPiFanTachoGPIO = 1
	RPiFanTachoNum  = 28

	// Temp
	OuterTempGPIO = 22
	OuterTempNum  = 15

	InnerTempGPIO = 23
	InnerTempNum  = 16

	DryerTempGPIO = 2
	DryerTempNum  = 3

	// Servo
	DryerServoGPIO = 3
	DryerServoNum  = 5

	// Heater buttons
	HeaterSwitchButtonGPIO = 21
	HeaterSwitchButtonNum  = 40

	HeaterModeButtonGPIO = 26
	HeaterModeButtonNum  = 37

	HeaterPlusButtonGPIO = 16
	HeaterPlusButtonNum  = 36

	HeaterMinusButtonGPIO = 20
	HeaterMinusButtonNum  = 38
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

	HeaterSwitchButtonPin = pin{
		j8Index:   HeaterSwitchButtonNum,
		gpioIndex: HeaterSwitchButtonGPIO,
		name:      "heater_switch_button_pin",
	}

	HeaterModeButtonPin = pin{
		j8Index:   HeaterModeButtonNum,
		gpioIndex: HeaterModeButtonGPIO,
		name:      "heater_mode_button_pin",
	}

	HeaterPlusButtonPin = pin{
		j8Index:   HeaterPlusButtonNum,
		gpioIndex: HeaterPlusButtonGPIO,
		name:      "heater_plus_button_pin",
	}

	HeaterMinusButtonPin = pin{
		j8Index:   HeaterMinusButtonNum,
		gpioIndex: HeaterMinusButtonGPIO,
		name:      "heater_minus_button_pin",
	}
)
