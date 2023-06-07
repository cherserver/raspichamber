package application

import (
	"fmt"

	"github.com/cherserver/raspichamber/service/hardware/uart"
	"github.com/cherserver/raspichamber/service/hardware/uart/devices"
)

const (
	UARTPortName = "/dev/ttyS0"
	UARTBaudRate = 115200
)

type Application struct {
	uart *uart.UART

	dryerControl *devices.DryerControl
	dryerHatch   *devices.DryerHatch

	innerFan   *devices.Fan
	outerFan   *devices.Fan
	rpiFan     *devices.Fan
	printerFan *devices.Fan

	innerThermometer *devices.Thermometer
	outerThermometer *devices.Thermometer
	dryerThermometer *devices.Thermometer
}

func New() *Application {
	uartPort := uart.New(UARTPortName, UARTBaudRate)
	return &Application{
		uart: uartPort,

		dryerControl: devices.NewDryerControl("heater", uartPort),
		dryerHatch:   devices.NewDryerHatch("hatch", uartPort),

		innerFan:   devices.NewFan("dryer_fan", uartPort),
		outerFan:   devices.NewFan("outer_fan", uartPort),
		rpiFan:     devices.NewFan("pi_fan", uartPort),
		printerFan: devices.NewFan("printer_fan", uartPort),

		innerThermometer: devices.NewThermometer("inner_th", uartPort),
		outerThermometer: devices.NewThermometer("outer_th", uartPort),
		dryerThermometer: devices.NewThermometer("dryer_th", uartPort),
	}
}

func (a *Application) Init() error {
	err := a.dryerControl.Register()
	if err != nil {
		return fmt.Errorf("failed to register dryer control: %w", err)
	}

	err = a.dryerHatch.Register()
	if err != nil {
		return fmt.Errorf("failed to register dryer servo: %w", err)
	}

	err = a.innerFan.Register()
	if err != nil {
		return fmt.Errorf("failed to register inner fan: %w", err)
	}

	err = a.outerFan.Register()
	if err != nil {
		return fmt.Errorf("failed to register outer fan: %w", err)
	}

	err = a.rpiFan.Register()
	if err != nil {
		return fmt.Errorf("failed to register RPi fan: %w", err)
	}

	err = a.printerFan.Register()
	if err != nil {
		return fmt.Errorf("failed to register printer fan: %w", err)
	}

	err = a.innerThermometer.Register()
	if err != nil {
		return fmt.Errorf("failed to register inner thermometer: %w", err)
	}

	err = a.outerThermometer.Register()
	if err != nil {
		return fmt.Errorf("failed to register outer thermometer: %w", err)
	}

	err = a.dryerThermometer.Register()
	if err != nil {
		return fmt.Errorf("failed to register dryer thermometer: %w", err)
	}

	err = a.uart.Init()
	if err != nil {
		return fmt.Errorf("failed to initialize UART: %w", err)
	}

	return nil
}

func (a *Application) Stop() {
	a.uart.Stop()
}

func (a *Application) InnerFan() *devices.Fan {
	return a.innerFan
}

func (a *Application) OuterFan() *devices.Fan {
	return a.outerFan
}

func (a *Application) RpiFan() *devices.Fan {
	return a.rpiFan
}

func (a *Application) PrinterFan() *devices.Fan {
	return a.printerFan
}

func (a *Application) InnerThermometer() *devices.Thermometer {
	return a.innerThermometer
}

func (a *Application) OuterThermometer() *devices.Thermometer {
	return a.outerThermometer
}

func (a *Application) DryerThermometer() *devices.Thermometer {
	return a.dryerThermometer
}

func (a *Application) DryerControl() *devices.DryerControl {
	return a.dryerControl
}

func (a *Application) DryerHatch() *devices.DryerHatch {
	return a.dryerHatch
}
