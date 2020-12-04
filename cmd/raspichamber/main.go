package main

import (
	"log"

	"github.com/cherserver/raspichamber/service/servo"
)

func main() {
	capServo := servo.New()
	err := capServo.Init()
	if err != nil {
		log.Fatalf("Failed to initialize capServo: %v", err)
	}
}
