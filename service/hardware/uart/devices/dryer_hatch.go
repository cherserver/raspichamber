package devices

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/uart"
	"github.com/cherserver/raspichamber/service/software"
)

var _ software.DryerHatch = &DryerHatch{}
var _ uart.StatusDataProcessor = &DryerHatch{}

type DryerHatch struct {
	uart *uart.UART

	name  string
	angle uint8
}

func NewDryerHatch(name string, uart *uart.UART) *DryerHatch {
	return &DryerHatch{
		uart: uart,

		name:  name,
		angle: 0,
	}
}

func (f *DryerHatch) Name() string {
	return f.name
}

func (f *DryerHatch) Register() error {
	return f.uart.RegisterStatusDataProcessor(f)
}

func (f *DryerHatch) Angle() uint8 {
	return f.angle
}

func (f *DryerHatch) SetAngle(value uint8) error {
	err := f.uart.Send(fmt.Sprintf("set-hatch-angle -angle %d", value))
	if err != nil {
		return err
	}

	f.angle = value

	return nil
}

func (f *DryerHatch) ProcessStatusData(data map[string]interface{}) {
	for key, val := range data {
		switch key {
		case "angle":
			floatVal, ok := val.(float64)
			if ok {
				f.angle = uint8(floatVal)
			}
			return
		}
	}
}
