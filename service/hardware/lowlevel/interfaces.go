package lowlevel

import "fmt"

type PinSubsystem interface {
	Init() error
	Stop()

	IsInitialized() bool
	Type() interface{} // not for type assertion, just for equation
}

type Pin interface {
	fmt.Stringer

	Name() string
	J8Index() int
	GPIOIndex() int
}

type PwmPin interface {
	Init() error
	Stop()

	SetPwmPercent(value uint8) error
}

type RpmSensorPin interface {
	Init() error
	Stop()

	RPM() (uint64, error)
}

type SwitchPin interface {
	Init() error
	Stop()

	TurnOn()
	TurnOff()
}
