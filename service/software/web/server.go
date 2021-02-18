package web

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/google/uuid"

	hardware "github.com/cherserver/raspichamber/service/hardware/application"
	"github.com/cherserver/raspichamber/service/software"
)

type server struct {
	currentSession uuid.UUID

	innerFan         software.InnerFan
	outerFan         software.OuterFan
	rpiFan           software.RpiFan
	innerThermometer software.InnerThermometer
	outerThermometer software.OuterThermometer
	dryerThermometer software.DryerThermometer
	dryerControl     software.DryerControl
	dryerHatch       software.DryerHatch
}

func NewServer(hardware *hardware.Application) *server {
	return &server{
		currentSession:   uuid.New(),
		innerFan:         hardware.InnerFan(),
		outerFan:         hardware.OuterFan(),
		rpiFan:           hardware.RpiFan(),
		innerThermometer: hardware.InnerThermometer(),
		outerThermometer: hardware.OuterThermometer(),
		dryerThermometer: hardware.DryerThermometer(),
		dryerControl:     hardware.DryerControl(),
		dryerHatch:       hardware.DryerHatch(),
	}
}

func (s *server) Init() error {
	fileServer := http.FileServer(http.Dir("./http"))
	http.Handle("/", fileServer)

	http.HandleFunc("/devices", s.statusHandler)
	http.HandleFunc("/devices/inner-fan", s.fanHandler)
	http.HandleFunc("/devices/outer-fan", s.fanHandler)
	http.HandleFunc("/devices/rpi-fan", s.fanHandler)
	http.HandleFunc("/devices/dryer-control", s.dryerControlHandler)
	http.HandleFunc("/devices/dryer-hatch/set-angle", s.dryerHatchHandler)

	server := &http.Server{Addr: ":8080", Handler: nil}
	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return fmt.Errorf("failed to start web server: %w", err)
	}

	go func() {
		srvError := server.Serve(ln)
		log.Printf("HTTP server stopped: %v", srvError)
	}()

	return nil
}

func (s *server) Stop() {

}

func (s *server) indexHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *server) statusHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *server) fanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	log.Printf("path '%v', form '%v'", r.URL.Path, r.Form)
}

func (s *server) dryerControlHandler(w http.ResponseWriter, r *http.Request) {

}
func (s *server) dryerHatchHandler(w http.ResponseWriter, r *http.Request) {

}
