package software

type Fan interface {
	RPM() uint64
	SpeedPercent() uint8

	SetSpeedPercent(value uint8) error
}

type Thermometer interface {
	Temperature() float32
	Humidity() float32
}

type InnerFan interface {
	Fan
}

type OuterFan interface {
	Fan
}

type RpiFan interface {
	Fan
}

type InnerThermometer interface {
	Thermometer
}

type OuterThermometer interface {
	Thermometer
}

type DryerThermometer interface {
	Thermometer
}

type DryerState uint8

const (
	DryerStateOff         DryerState = 0
	DryerStateOn55Degrees DryerState = 1
	DryerStateOn60Degrees DryerState = 2
	DryerStateOn65Degrees DryerState = 3
	DryerStateOn70Degrees DryerState = 4
)

type DryerControl interface {
	State() DryerState

	SetState(state DryerState)
}

type DryerHatch interface {
	Angle() uint8

	SetAngle(angle uint8) error
}
