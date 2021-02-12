package lowlevel

import "fmt"

type PinSubsystem interface {
	Init() error
	Stop()

	IsInitialized() bool
	Type() interface{}
}

type Pin interface {
	fmt.Stringer

	GPIOIndex() int
	Name() string
}

type PwmPin interface {
	Init() error
	Stop()

	SetPwmPercent(value uint8) error
}

type RpmSensorPin interface {
	Init() error
	Stop()

	RPM() (uint32, error)
}

type SwitchPin interface {
	Init() error
	Stop()

	TurnOn()
	TurnOff()
}
