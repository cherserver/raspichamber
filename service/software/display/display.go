package display

import (
	"fmt"
	"image/color"
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
	intHeight  = 240
	intWidth   = 240
	height     = float64(intHeight)
	width      = float64(intWidth)
	halfHeight = height / 2
	halfWidth  = width / 2

	mainBorderSize  = 3
	childBorderSize = 2

	defaultMargin  = 10
	leftmostMargin = 5

	firstLineOffsetY  = 20
	firstLineBottom   = 40
	secondLineOffsetY = 60
	secondLineBottom  = 80
	thirdLineOffsetY  = 100

	fontSize            = 30
	statusImageFilePath = "../raspichamber_display/status.png"
	temperatureTxtFmt   = "%+05.1f°"
	humidityTxtFmt      = "%5.1f%%"
	fanTxtFmt           = "%-1s %3d%%"
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
	statusDraw.DrawString(caption, offsetX, offsetY+firstLineOffsetY)
	statusDraw.DrawString(fmt.Sprintf(temperatureTxtFmt, thermometer.Temperature()), offsetX, offsetY+secondLineOffsetY)
	statusDraw.DrawString(fmt.Sprintf(humidityTxtFmt, thermometer.Humidity()), offsetX, offsetY+thirdLineOffsetY)
}

func (d *display) saveStatusImage() error {
	// deep blue
	backgroundColor := color.RGBA{
		R: 0,
		G: 0x26,
		B: 0x4D,
		A: 0xff,
	}

	// light yellow
	textColor := color.RGBA{
		R: 0xFF,
		G: 0xFF,
		B: 0xAD,
		A: 0xff,
	}

	statusDraw := gg.NewContext(intWidth, intHeight)
	statusDraw.SetColor(backgroundColor)
	statusDraw.Clear()

	statusDraw.RotateAbout(gg.Radians(180), halfWidth, halfHeight)

	statusDraw.SetColor(textColor)

	// border cross
	statusDraw.SetLineWidth(mainBorderSize)
	statusDraw.DrawLine(halfWidth, 0, halfWidth, height)
	statusDraw.Stroke()
	statusDraw.DrawLine(0, halfHeight, width, halfHeight)
	statusDraw.Stroke()

	// fan borders
	statusDraw.SetLineWidth(childBorderSize)
	statusDraw.DrawLine(halfWidth, halfHeight+firstLineBottom, width, halfHeight+firstLineBottom)
	statusDraw.Stroke()
	statusDraw.DrawLine(halfWidth, halfHeight+secondLineBottom, width, halfHeight+secondLineBottom)
	statusDraw.Stroke()

	font, err := truetype.Parse(gomonobold.TTF)
	fontFace := truetype.NewFace(font, &truetype.Options{Size: fontSize})
	statusDraw.SetFontFace(fontFace)

	secondHalfX := halfWidth + defaultMargin
	secondHalfY := halfHeight + defaultMargin

	d.printTemp(statusDraw, "Inner", d.hardware.InnerThermometer(), leftmostMargin, defaultMargin)
	d.printTemp(statusDraw, "Outer", d.hardware.OuterThermometer(), secondHalfX, defaultMargin)
	d.printTemp(statusDraw, "Dryer", d.hardware.DryerThermometer(), leftmostMargin, secondHalfY)

	statusDraw.DrawString(fmt.Sprintf(fanTxtFmt, "I", d.hardware.InnerFan().SpeedPercent()), secondHalfX, secondHalfY+firstLineOffsetY)
	statusDraw.DrawString(fmt.Sprintf(fanTxtFmt, "O", d.hardware.OuterFan().SpeedPercent()), secondHalfX, secondHalfY+secondLineOffsetY)
	statusDraw.DrawString(fmt.Sprintf(fanTxtFmt, "R", d.hardware.RpiFan().SpeedPercent()), secondHalfX, secondHalfY+thirdLineOffsetY)

	tmpPath := statusImageFilePath + "_tmp"
	err = statusDraw.SavePNG(tmpPath)
	if err != nil {
		return fmt.Errorf("failed to save status image file '%v': %w", tmpPath, err)
	}

	// "atomic write"
	if err = os.Rename(tmpPath, statusImageFilePath); err != nil {
		return fmt.Errorf("failed to rename status image file from '%v' to '%v': %w", tmpPath, statusImageFilePath, err)
	}

	return nil
}
