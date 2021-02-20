package web

import (
	"github.com/google/uuid"

	"github.com/cherserver/raspichamber/service/software"
)

type fan struct {
	speedPercent uint8
	rpm          uint32
}

type thermometer struct {
	temperature float32
	humidity    float32
}

type dryerControl struct {
	state software.DryerState
}

type dryerHatch struct {
	angle uint8
}

type devices struct {
	innerFan         fan
	outerFan         fan
	rpiFan           fan
	innerThermometer thermometer
	outerThermometer thermometer
	dryerThermometer thermometer
	dryerControl     dryerControl
	dryerHatch       dryerHatch
}

type DevicesStatus struct {
	CurrentSession uuid.UUID

	devices devices
}
