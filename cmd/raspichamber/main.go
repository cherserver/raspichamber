package main

import (
	"log"

	"github.com/cherserver/raspichamber/service/fan"
	"github.com/cherserver/raspichamber/service/pinout"
	"github.com/cherserver/raspichamber/service/servo"
	"github.com/cherserver/raspichamber/service/thermometer"
)

func main() {
	capServo := servo.New(pinout.ServoPin)
	err := capServo.Init()
	if err != nil {
		log.Fatalf("Failed to initialize capServo: %v", err)
	}

	externalFan := fan.New(pinout.ExternalFanPwmPin, pinout.ExternalFanTachometerPin)
	err = externalFan.Init()
	if err != nil {
		log.Fatalf("Failed to initialize external fan: %v", err)
	}

	internalFan := fan.New(pinout.InternalFanPwmPin, pinout.InternalFanTachometerPin)
	err = internalFan.Init()
	if err != nil {
		log.Fatalf("Failed to initialize internal fan: %v", err)
	}

	outerTemp := thermometer.New(pinout.OuterTempPin)
	err = outerTemp.Init()
	if err != nil {
		log.Fatalf("Failed to initialize outer thermometer: %v", err)
	}

	innerTemp := thermometer.New(pinout.InnerTempPin)
	err = innerTemp.Init()
	if err != nil {
		log.Fatalf("Failed to initialize inner thermometer: %v", err)
	}

	filamentTemp := thermometer.New(pinout.FilamentTempPin)
	err = filamentTemp.Init()
	if err != nil {
		log.Fatalf("Failed to initialize filament thermometer: %v", err)
	}

	// TODO: wait for sigterm, sighup
	// signal.Notify()

	filamentTemp.Stop()
	innerTemp.Stop()
	outerTemp.Stop()
	internalFan.Stop()
	externalFan.Stop()
	capServo.Stop()
}
