package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"

	hardware "github.com/cherserver/raspichamber/service/hardware/application"
	"github.com/cherserver/raspichamber/service/software"
)

type server struct {
	currentSessionId uuid.UUID

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
		currentSessionId: uuid.New(),
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

	http.HandleFunc("/reset", s.resetHandler)

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

func (s *server) resetHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("reset caught, exit")
	os.Exit(0)
}

func (s *server) statusHandler(w http.ResponseWriter, r *http.Request) {
	_ = r

	statusData, err := json.Marshal(*s.devicesStatus())
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to encode status: %v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("status: %s", statusData)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(statusData)
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

	strVal, err := s.getFormValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	percent, err := s.parsePercent(strVal)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid 'value' field: %v", err), http.StatusBadRequest)
	}

	log.Printf("set fan speed percent '%v'", percent)
	err = fan.SetSpeedPercent(percent)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to set fan speed percent: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) dryerControlSetStateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	strVal, err := s.getFormValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("set hatch control state '%v'", strVal)

	switch strVal {
	case "off":
		s.dryerControl.SetState(software.DryerStateOff)
	case "on35degrees":
		s.dryerControl.SetState(software.DryerStateOn35Degrees)
	case "on40degrees":
		s.dryerControl.SetState(software.DryerStateOn40Degrees)
	case "on45degrees":
		s.dryerControl.SetState(software.DryerStateOn45Degrees)
	case "on50degrees":
		s.dryerControl.SetState(software.DryerStateOn50Degrees)
	case "on55degrees":
		s.dryerControl.SetState(software.DryerStateOn55Degrees)
	case "on60degrees":
		s.dryerControl.SetState(software.DryerStateOn60Degrees)
	case "on65degrees":
		s.dryerControl.SetState(software.DryerStateOn65Degrees)
	case "on70degrees":
		s.dryerControl.SetState(software.DryerStateOn70Degrees)
	default:
		http.Error(
			w,
			"value must be state - [off, on35degrees, on40degrees, on45degrees, on50degrees, on55degrees, on60degrees, on65degrees, on70degrees]",
			http.StatusBadRequest,
		)
		return
	}
}

func (s *server) dryerHatchSetAngleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	strVal, err := s.getFormValue(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	angle, err := s.parseAngle(strVal)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid 'value' field: %v", err), http.StatusBadRequest)
	}

	log.Printf("set dryer hatch angle '%v'", angle)
	err = s.dryerHatch.SetAngle(angle)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to set dryer hatch angle: %v", err), http.StatusInternalServerError)
	}
}

func (s *server) getFormValue(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", errors.New("invalid form")
	}

	strVal := r.PostFormValue("value")
	if len(strVal) == 0 {
		return "", errors.New("'value' field is missed")
	}

	return strVal, nil
}

func (s *server) parseInt(value string) (int64, error) {
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value: %w", err)
	}

	return val, nil
}

func (s *server) parsePercent(value string) (uint8, error) {
	val, err := s.parseInt(value)
	if err != nil {
		return 0, fmt.Errorf("invalid value: %w", err)
	}

	if val < 0 || val > 100 {
		return 0, errors.New("value must be percent value - [0-100]")
	}

	return uint8(val), nil
}

func (s *server) parseAngle(value string) (uint8, error) {
	val, err := s.parseInt(value)
	if err != nil {
		return 0, fmt.Errorf("invalid value: %w", err)
	}

	if val < 0 || val > 90 {
		return 0, errors.New("value must be angle value - [0-90]")
	}

	return uint8(val), nil
}

func (s *server) devicesStatus() *DevicesStatus {
	return &DevicesStatus{
		CurrentSession: s.currentSessionId,
		Devices: Devices{
			InnerFan: Fan{
				SpeedPercent: s.innerFan.SpeedPercent(),
				RPM:          s.innerFan.RPM(),
			},
			OuterFan: Fan{
				SpeedPercent: s.outerFan.SpeedPercent(),
				RPM:          s.outerFan.RPM(),
			},
			RPiFan: Fan{
				SpeedPercent: s.rpiFan.SpeedPercent(),
				RPM:          s.rpiFan.RPM(),
			},
			InnerThermometer: Thermometer{
				Temperature: s.innerThermometer.Temperature(),
				Humidity:    s.innerThermometer.Humidity(),
			},
			OuterThermometer: Thermometer{
				Temperature: s.outerThermometer.Temperature(),
				Humidity:    s.outerThermometer.Humidity(),
			},
			DryerThermometer: Thermometer{
				Temperature: s.dryerThermometer.Temperature(),
				Humidity:    s.dryerThermometer.Humidity(),
			},
			DryerControl: DryerControl{
				State: s.dryerControl.State(),
			},
			DryerHatch: DryerHatch{
				Angle: s.dryerHatch.Angle(),
			},
		},
	}
}
