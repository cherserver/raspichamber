package software

type Fan interface {
	RPM() uint64
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

type DryerControl interface {
	SwitchOff()
	Switch55Degrees()
	Switch60Degrees()
	Switch65Degrees()
	Switch70Degrees()
}

type DryerHatch interface {
	SetAngle(angle uint8) error
}
