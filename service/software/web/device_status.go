package web

import (
	"github.com/google/uuid"

	"github.com/cherserver/raspichamber/service/software"
)

type Fan struct {
	SpeedPercent uint8  `json:"speed_percent"`
	RPM          uint32 `json:"rpm"`
}

type Thermometer struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

type DryerControl struct {
	State software.DryerState `json:"state"`
}

type DryerHatch struct {
	Angle uint8 `json:"angle"`
}

type Devices struct {
	InnerFan         Fan          `json:"inner_fan"`
	OuterFan         Fan          `json:"outer_fan"`
	RPiFan           Fan          `json:"rpi_fan"`
	PrinterFan       Fan          `json:"printer_fan"`
	InnerThermometer Thermometer  `json:"inner_thermometer"`
	OuterThermometer Thermometer  `json:"outer_thermometer"`
	DryerThermometer Thermometer  `json:"dryer_thermometer"`
	DryerControl     DryerControl `json:"dryer_control"`
	DryerHatch       DryerHatch   `json:"dryer_hatch"`
}

type DevicesStatus struct {
	CurrentSession uuid.UUID `json:"current_session"`

	Devices Devices `json:"devices"`
}
