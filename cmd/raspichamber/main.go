package main

import (
	"log"

	"github.com/cherserver/raspichamber/service/fan"
	"github.com/cherserver/raspichamber/service/servo"
)

func main() {
	capServo := servo.New()
	err := capServo.Init()
	if err != nil {
		log.Fatalf("Failed to initialize capServo: %v", err)
	}

	firstFan := fan.New()
	err = firstFan.Init()
	if err != nil {
		log.Fatalf("Failed to initialize fan: %v", err)
	}

	// TODO: wait for sigterm, sighup
	// signal.Notify()

	firstFan.Stop()
	capServo.Stop()
}
