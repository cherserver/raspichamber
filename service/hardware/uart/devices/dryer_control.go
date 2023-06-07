package devices

import (
	"fmt"
	"log"

	"github.com/cherserver/raspichamber/service/hardware/uart"
	"github.com/cherserver/raspichamber/service/software"
)

var _ software.DryerControl = &DryerControl{}
var _ uart.StatusDataProcessor = &DryerControl{}

type DryerControl struct {
	uart *uart.UART

	name  string
	state software.DryerState
}

func NewDryerControl(name string, uart *uart.UART) *DryerControl {
	return &DryerControl{
		uart: uart,

		name:  name,
		state: software.DryerStateOff,
	}
}

func (f *DryerControl) Name() string {
	return f.name
}

func (f *DryerControl) Register() error {
	return f.uart.RegisterStatusDataProcessor(f)
}

func (f *DryerControl) State() software.DryerState {
	return f.state
}

func (f *DryerControl) SetState(value software.DryerState) error {
	mode := ""
	switch value {
	case software.DryerStateOff:
		mode = "off"
	case software.DryerStateOn35Degrees:
		mode = "35"
	case software.DryerStateOn40Degrees:
		mode = "40"
	case software.DryerStateOn45Degrees:
		mode = "45"
	case software.DryerStateOn50Degrees:
		mode = "50"
	case software.DryerStateOn55Degrees:
		mode = "55"
	case software.DryerStateOn60Degrees:
		mode = "60"
	case software.DryerStateOn65Degrees:
		mode = "65"
	case software.DryerStateOn70Degrees:
		mode = "70"
	default:
		return fmt.Errorf("unknown dryer control state: %v", value)
	}

	err := f.uart.Send(fmt.Sprintf("set-heater-mode -mode %s", mode))
	if err != nil {
		return err
	}

	f.state = value

	return nil
}

func (f *DryerControl) ProcessStatusData(data map[string]interface{}) {
	for key, val := range data {
		switch key {
		case "mode":
			stringVal, ok := val.(string)
			if !ok {
				return
			}

			switch stringVal {
			case "off":
				f.state = software.DryerStateOff
			case "35":
				f.state = software.DryerStateOn35Degrees
			case "40":
				f.state = software.DryerStateOn40Degrees
			case "45":
				f.state = software.DryerStateOn45Degrees
			case "50":
				f.state = software.DryerStateOn50Degrees
			case "55":
				f.state = software.DryerStateOn55Degrees
			case "60":
				f.state = software.DryerStateOn60Degrees
			case "65":
				f.state = software.DryerStateOn65Degrees
			case "70":
				f.state = software.DryerStateOn70Degrees
			default:
				log.Printf("Unknown dryer mode '%s' came in status data", stringVal)
			}

			return
		}
	}
}
