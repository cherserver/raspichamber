package pinout

const (
	InternalFanPwmPin        = 12
	InternalFanTachometerPin = 17

	ExternalFanPwmPin        = 13
	ExternalFanTachometerPin = 4

	PiFanPwmPin        = 0
	PiFanTachometerPin = 1

	OuterTempPin    = 23
	InnerTempPin    = 22
	FilamentTempPin = 2

	ServoPin = "5" // physical pin on header, not a GPIO pin

	HeaterButton1 = 26
	HeaterButton2 = 16
	HeaterButton3 = 20
	HeaterButton4 = 21
)
