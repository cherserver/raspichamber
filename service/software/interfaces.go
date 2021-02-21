package software

type Fan interface {
	RPM() uint32
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

type DryerState string

const (
	DryerStateOff         DryerState = "off"
	DryerStateOn35Degrees DryerState = "on35degrees"
	DryerStateOn40Degrees DryerState = "on40degrees"
	DryerStateOn45Degrees DryerState = "on45degrees"
	DryerStateOn50Degrees DryerState = "on50degrees"
	DryerStateOn55Degrees DryerState = "on55degrees"
	DryerStateOn60Degrees DryerState = "on60degrees"
	DryerStateOn65Degrees DryerState = "on65degrees"
	DryerStateOn70Degrees DryerState = "on70degrees"
)

type DryerControl interface {
	State() DryerState

	SetState(state DryerState)
}

type DryerHatch interface {
	Angle() uint8

	SetAngle(angle uint8) error
}
