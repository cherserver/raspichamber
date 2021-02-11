package hardware

type ExternalPwmSubsystem interface {
	Init() error
	Stop()

	SetPinPwmPercent(pin int, value uint8) error
}
