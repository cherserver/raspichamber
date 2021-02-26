package display

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gomonobold"

	hardware "github.com/cherserver/raspichamber/service/hardware/application"
	"github.com/cherserver/raspichamber/service/software"
)

const (
	statusImageFilePath = "../raspichamber_display/status.jpg"
	temperatureTxtFmt   = "🌡%+2.2f°C"
	humidityTxtFmt      = "💧%2.2f%%"
	fanTxtFmt           = "💨%3s %03d%%"
)

type display struct {
	hardware *hardware.Application
}

func New(hardware *hardware.Application) *display {
	return &display{
		hardware: hardware,
	}
}

func (d *display) Init() error {
	go d.worker()
	return nil
}

func (d *display) Stop() {

}

func (d *display) worker() {
	for {
		err := d.saveStatusImage()
		if err != nil {
			log.Printf("Failed to save status image: %v", err)
		}

		time.Sleep(1 * time.Second)
	}
}

func (d *display) printTemp(statusDraw *gg.Context, caption string, thermometer software.Thermometer, offsetX float64, offsetY float64) {
	statusDraw.DrawString(caption, offsetX, offsetY+20)
	statusDraw.DrawString(fmt.Sprintf(temperatureTxtFmt, thermometer.Temperature()), offsetX, offsetY+60)
	statusDraw.DrawString(fmt.Sprintf(humidityTxtFmt, thermometer.Humidity()), offsetX, offsetY+100)
}

func (d *display) saveStatusImage() error {
	// 00172D
	// 00264D
	backgroundColor := color.RGBA{
		R: 0,
		G: 0x26,
		B: 0x4D,
		A: 0xff,
	}

	textColor := color.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0xAD,
		A: 0xff,
	}

	height := 240
	width := 240

	statusDraw := gg.NewContext(width, height)
	statusDraw.SetColor(backgroundColor)
	statusDraw.Clear()

	statusDraw.RotateAbout(gg.Radians(180), float64(width/2), float64(height/2))

	statusDraw.SetColor(textColor)
	statusDraw.SetLineWidth(3)
	statusDraw.DrawLine(float64(width/2), 0, float64(width/2), float64(height))
	statusDraw.Stroke()
	statusDraw.DrawLine(0, float64(height/2), float64(width), float64(height/2))
	statusDraw.Stroke()

	font, err := truetype.Parse(gomonobold.TTF)
	fontFace := truetype.NewFace(font, &truetype.Options{
		Size: 20,
		// Hinting: font.HintingFull,
	})
	statusDraw.SetFontFace(fontFace)

	d.printTemp(statusDraw, "Inner", d.hardware.InnerThermometer(), 0, 0)
	d.printTemp(statusDraw, "Outer", d.hardware.OuterThermometer(), float64(width/2), 0)
	d.printTemp(statusDraw, "Dryer", d.hardware.DryerThermometer(), 0, float64(height/2))

	statusDraw.DrawString(fmt.Sprintf(fanTxtFmt, "IN", d.hardware.InnerFan().SpeedPercent()), float64(width/2), float64(height/2)+60)
	statusDraw.DrawString(fmt.Sprintf(fanTxtFmt, "OUT", d.hardware.OuterFan().SpeedPercent()), float64(width/2), float64(height/2)+60)
	statusDraw.DrawString(fmt.Sprintf(fanTxtFmt, "RPi", d.hardware.RpiFan().SpeedPercent()), float64(width/2), float64(height/2)+100)

	tmpPath := statusImageFilePath + "_tmp"
	f, err := os.Create(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to open status image file '%v': %w", tmpPath, err)
	}

	defer func() { _ = f.Close() }()

	if err = jpeg.Encode(f, statusDraw.Image(), &jpeg.Options{
		Quality: jpeg.DefaultQuality,
	}); err != nil {
		return fmt.Errorf("failed to encode status image to file '%v': %w", tmpPath, err)
	}

	if err = os.Rename(tmpPath, statusImageFilePath); err != nil {
		return fmt.Errorf("failed to rename status image file from '%v' to '%v': %w", tmpPath, statusImageFilePath, err)
	}

	return nil
}
