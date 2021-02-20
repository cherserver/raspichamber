package web

import (
	"github.com/google/uuid"

	"github.com/cherserver/raspichamber/service/software"
)

type Fan struct {
	SpeedPercent uint8
	RPM          uint32
}

type Thermometer struct {
	Temperature float32
	Humidity    float32
}

type DryerControl struct {
	State software.DryerState
}

type DryerHatch struct {
	Angle uint8
}

type Devices struct {
	InnerFan         Fan
	OuterFan         Fan
	RPiFan           Fan
	InnerThermometer Thermometer
	OuterThermometer Thermometer
	DryerThermometer Thermometer
	DryerControl     DryerControl
	DryerHatch       DryerHatch
}

type DevicesStatus struct {
	CurrentSession uuid.UUID

	Devices Devices
}
