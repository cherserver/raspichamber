package web

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

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
	http.HandleFunc("/devices/inner-fan/set-speed-percent", s.innerFanSetSpeedHandler)
	http.HandleFunc("/devices/outer-fan/set-speed-percent", s.outerFanSetSpeedHandler)
	http.HandleFunc("/devices/rpi-fan/set-speed-percent", s.rpiFanSetSpeedHandler)
	http.HandleFunc("/devices/dryer-control/set-state", s.dryerControlSetStateHandler)
	http.HandleFunc("/devices/dryer-hatch/set-angle", s.dryerHatchSetAngleHandler)

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
	_ = w
	_ = r
}

func (s *server) statusHandler(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func (s *server) innerFanSetSpeedHandler(w http.ResponseWriter, r *http.Request) {
	s.fanSetSpeedHandler(w, r, s.innerFan)
}

func (s *server) outerFanSetSpeedHandler(w http.ResponseWriter, r *http.Request) {
	s.fanSetSpeedHandler(w, r, s.outerFan)
}

func (s *server) rpiFanSetSpeedHandler(w http.ResponseWriter, r *http.Request) {
	s.fanSetSpeedHandler(w, r, s.rpiFan)
}

func (s *server) fanSetSpeedHandler(w http.ResponseWriter, r *http.Request, fan software.Fan) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}

	percent, err := s.parsePercent(r.PostFormValue("value"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid 'value' field: %v", err), http.StatusBadRequest)
	}

	log.Printf("set fan percent '%v'", percent)
	err = fan.SetSpeedPercent(percent)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to set speed percent: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) dryerControlSetStateHandler(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func (s *server) dryerHatchSetAngleHandler(w http.ResponseWriter, r *http.Request) {
	_ = w
	_ = r
}

func (s *server) parsePercent(value string) (uint8, error) {
	if len(value) == 0 {
		return 0, errors.New("value is missed")
	}

	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value: %w", err)
	}

	if val < 0 || val > 100 {
		return 0, errors.New("value must be percent value - [0-100]")
	}

	return uint8(val), nil
}
