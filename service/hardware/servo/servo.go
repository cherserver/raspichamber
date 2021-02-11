package servo

import (
	"github.com/cherserver/raspichamber/service/hardware"
)

type servo struct {
	pin           int
	requiredAngle uint32

	pwmSys hardware.ExternalPwmSubsystem
}

func New(pin int, pwmSys hardware.ExternalPwmSubsystem) *servo {
	return &servo{
		pin:           pin,
		requiredAngle: 0,

		pwmSys: pwmSys,
	}
}

func (s *servo) Init() error {
	return nil
}

func (s *servo) Stop() {
}

func (s *servo) SetAngle(angle uint8) error {
	return s.pwmSys.SetPinPwmPercent(s.pin, angle)
}
