package application

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/devices"
)

type application struct {
	pinSubsystems *devices.PinSubsystems

	dryerControl *devices.DryerControl
	dryerServo   *devices.DryerServo

	innerFan *devices.Fan
	outerFan *devices.Fan
	rpiFan   *devices.Fan

	innerThermometer *devices.Thermometer
	outerThermometer *devices.Thermometer
	dryerThermometer *devices.Thermometer
}

func New() *application {
	pinSubsystems := devices.NewPinSubsystems()
	return &application{
		pinSubsystems: pinSubsystems,

		dryerControl: devices.NewDryerControl(pinSubsystems),
		dryerServo:   devices.NewDryerServo(pinSubsystems),

		innerFan: devices.NewInnerFan(pinSubsystems),
		outerFan: devices.NewOuterFan(pinSubsystems),
		rpiFan:   devices.NewRPiFan(pinSubsystems),

		innerThermometer: devices.NewInnerThermometer(pinSubsystems),
		outerThermometer: devices.NewOuterThermometer(pinSubsystems),
		dryerThermometer: devices.NewDryerThermometer(pinSubsystems),
	}
}

func (a *application) Init() error {
	err := a.pinSubsystems.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize pin subsystems: %w", err)
	}

	err = a.dryerControl.Init()
	if err != nil {
		return fmt.Errorf("failed to init dryer control: %w", err)
	}

	err = a.dryerServo.Init()
	if err != nil {
		return fmt.Errorf("failed to init dryer servo: %w", err)
	}

	err = a.innerFan.Init()
	if err != nil {
		return fmt.Errorf("failed to init inner fan: %w", err)
	}

	err = a.outerFan.Init()
	if err != nil {
		return fmt.Errorf("failed to init outer fan: %w", err)
	}

	err = a.rpiFan.Init()
	if err != nil {
		return fmt.Errorf("failed to init RPi fan: %w", err)
	}

	err = a.innerThermometer.Init()
	if err != nil {
		return fmt.Errorf("failed to init inner thermometer: %w", err)
	}

	err = a.outerThermometer.Init()
	if err != nil {
		return fmt.Errorf("failed to init outer thermometer: %w", err)
	}

	err = a.dryerThermometer.Init()
	if err != nil {
		return fmt.Errorf("failed to init dryer thermometer: %w", err)
	}

	return nil
}

func (a *application) Stop() {
	a.dryerThermometer.Stop()
	a.outerThermometer.Stop()
	a.innerThermometer.Stop()

	a.rpiFan.Stop()
	a.outerFan.Stop()
	a.innerFan.Stop()

	a.dryerServo.Stop()
	a.dryerControl.Stop()

	a.pinSubsystems.Stop()
}
