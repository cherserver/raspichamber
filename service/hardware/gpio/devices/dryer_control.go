package devices

import (
	"fmt"
	"sync"
	"time"

	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel"
	"github.com/cherserver/raspichamber/service/hardware/gpio/lowlevel/pins"
	"github.com/cherserver/raspichamber/service/software"
)

const (
	// TODO: move to config
	pressDuration  = 200 * time.Millisecond
	releaseTimeout = 100 * time.Millisecond
)

var _ software.DryerControl = &DryerControl{}

type DryerControl struct {
	powerButton *button
	minusButton *button
	plusButton  *button
	modeButton  *button

	isSwitchedOn bool
	state        software.DryerState
	stateLock    sync.RWMutex
}

func (c *DryerControl) State() software.DryerState {
	c.stateLock.RLock()
	defer c.stateLock.RUnlock()

	return c.state
}

func (c *DryerControl) SetState(state software.DryerState) error {
	switch state {
	case software.DryerStateOff:
		c.SwitchOff()
	case software.DryerStateOn35Degrees:
		c.Switch35Degrees()
	case software.DryerStateOn40Degrees:
		c.Switch40Degrees()
	case software.DryerStateOn45Degrees:
		c.Switch45Degrees()
	case software.DryerStateOn50Degrees:
		c.Switch50Degrees()
	case software.DryerStateOn55Degrees:
		c.Switch55Degrees()
	case software.DryerStateOn60Degrees:
		c.Switch60Degrees()
	case software.DryerStateOn65Degrees:
		c.Switch65Degrees()
	case software.DryerStateOn70Degrees:
		c.Switch70Degrees()
	default:
		panic(fmt.Sprintf("Try to set unknown dryer state '%v'", state))
	}

	return nil
}

func NewDryerControl(subsystems *PinSubsystems) *DryerControl {
	return &DryerControl{
		powerButton: newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterSwitchButtonPin)),
		minusButton: newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterMinusButtonPin)),
		plusButton:  newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterPlusButtonPin)),
		modeButton:  newButton(pins.NewSwitchPin(subsystems.NativePinSubsystem, lowlevel.HeaterModeButtonPin)),

		state: software.DryerStateOff,
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
	if c.state == software.DryerStateOff {
		return
	}

	c.powerButton.Press()
	c.state = software.DryerStateOff
}

func (c *DryerControl) unsafeReset() {
	c.unsafeSwitchOff()

	c.powerButton.Press()
	c.modeButton.Press()
	c.state = software.DryerStateOn35Degrees
}

func (c *DryerControl) SwitchOff() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeSwitchOff()
}

func (c *DryerControl) Switch35Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()
}

func (c *DryerControl) Switch40Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()

	c.plusButton.PressTimes(1)
	c.state = software.DryerStateOn40Degrees
}

func (c *DryerControl) Switch45Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()

	c.plusButton.PressTimes(2)
	c.state = software.DryerStateOn45Degrees
}

func (c *DryerControl) Switch50Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()

	c.plusButton.PressTimes(3)
	c.state = software.DryerStateOn50Degrees
}

func (c *DryerControl) Switch55Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()

	c.plusButton.PressTimes(4)
	c.state = software.DryerStateOn55Degrees
}

func (c *DryerControl) Switch60Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()

	c.plusButton.PressTimes(5)
	c.state = software.DryerStateOn60Degrees
}

func (c *DryerControl) Switch65Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()
	c.plusButton.PressTimes(6)
	c.state = software.DryerStateOn65Degrees
}

func (c *DryerControl) Switch70Degrees() {
	c.stateLock.Lock()
	defer c.stateLock.Unlock()

	c.unsafeReset()
	c.plusButton.PressTimes(7)
	c.state = software.DryerStateOn70Degrees
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
