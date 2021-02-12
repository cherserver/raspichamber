package devices

import (
	"fmt"
	"sync"
	"time"

	"github.com/cherserver/raspichamber/service/hardware/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/lowlevel/pins"
)

const (
	// TODO: move to config
	pressDuration  = 200 * time.Millisecond
	releaseTimeout = 100 * time.Millisecond
)

type DryerControl struct {
	powerButton *button
	minusButton *button
	plusButton  *button
	modeButton  *button

	isSwitchedOn bool
	stateLock    sync.Mutex
}

func NewDryerControl(subsystems *PinSubsystems) *DryerControl {
	return &DryerControl{
		powerButton: newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterButton1Pin)),
		minusButton: newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterButton2Pin)),
		plusButton:  newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterButton3Pin)),
		modeButton:  newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterButton4Pin)),
	}
}

func (c *DryerControl) Init() error {
	err := c.powerButton.Init()
	if err != nil {
		return fmt.Errorf("failed to init 'power' button of dryer control: %w", err)
	}

	err = c.minusButton.Init()
	if err != nil {
		return fmt.Errorf("failed to init 'minus' button of dryer control: %w", err)
	}

	err = c.plusButton.Init()
	if err != nil {
		return fmt.Errorf("failed to init 'plus' button of dryer control: %w", err)
	}

	err = c.modeButton.Init()
	if err != nil {
		return fmt.Errorf("failed to init 'mode' button of dryer control: %w", err)
	}

	return nil
}

func (c *DryerControl) Stop() {
	c.modeButton.Stop()
	c.plusButton.Stop()
	c.minusButton.Stop()
	c.powerButton.Stop()
}

func (c *DryerControl) unsafeSwitchOff() {
	if !c.isSwitchedOn {
		return
	}

	c.powerButton.Press()
	c.isSwitchedOn = false
}

func (c *DryerControl) unsafeReset() {
	c.unsafeSwitchOff()

	c.powerButton.Press()
	c.isSwitchedOn = true
}

func (c *DryerControl) SwitchOff() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeSwitchOff()
}

func (c *DryerControl) Switch55Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()
}

func (c *DryerControl) Switch60Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()
	c.plusButton.PressTimes(1)
}

func (c *DryerControl) Switch65Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()
	c.plusButton.PressTimes(2)
}

func (c *DryerControl) Switch70Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()
	c.plusButton.PressTimes(3)
}

type button struct {
	pin lowlevel.SwitchPin
}

func newButton(pin lowlevel.SwitchPin) *button {
	return &button{
		pin: pin,
	}
}

func (b *button) Init() error {
	return b.pin.Init()
}

func (b *button) Stop() {
	b.pin.Stop()
}

func (b *button) Press() {
	b.pin.TurnOn()
	time.Sleep(pressDuration)

	b.pin.TurnOff()
	time.Sleep(releaseTimeout)
}

func (b *button) PressTimes(times uint8) {
	for cnt := uint8(0); cnt < times; cnt++ {
		b.Press()
	}
}
