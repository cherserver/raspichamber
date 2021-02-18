package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cherserver/raspichamber/service/hardware/application"
	"github.com/cherserver/raspichamber/service/software/web"
)

func main() {
	var err error

	hardwareApp := application.New()
	err = hardwareApp.Init()
	if err != nil {
		log.Fatalf("Failed to initialize hardware application: %v", err)
	}

	webServer := web.NewServer(hardwareApp)
	err = webServer.Init()
	if err != nil {
		log.Fatalf("Failed to initialize web server: %v", err)
	}

	stopSignalCh := make(chan os.Signal, 1)
	signal.Notify(stopSignalCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	stopSignal := <-stopSignalCh
	log.Printf("Signal '%+v' caught, exit", stopSignal)

	webServer.Stop()
	hardwareApp.Stop()
}
