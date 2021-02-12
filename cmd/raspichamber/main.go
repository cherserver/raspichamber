package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cherserver/raspichamber/service/hardware/application"
)

func main() {
	var err error

	hardwareApp := application.New()
	err = hardwareApp.Init()
	if err != nil {
		log.Fatalf("Failed to initialize hardware application: %v", err)
	}

	stopSignalCh := make(chan os.Signal, 1)
	signal.Notify(stopSignalCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	stopSignal := <-stopSignalCh
	log.Printf("Signal '%+v' caught, exit", stopSignal)

	hardwareApp.Stop()
}
