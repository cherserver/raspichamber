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
	"golang.org/x/image/font/gofont/gomono"
)

const (
	statusImageFilePath = "../raspichamber_display/status.jpg"
)

type display struct {
}

func New() *display {
	return &display{}
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

func (d *display) saveStatusImage() error {
	// 00172D
	// 00264D
	backgroundColor := color.RGBA{
		R: 0,
		G: 0x17,
		B: 0x2D,
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

	temperatureTxtFmt := "🌡 %+2.2f °C"
	humidityTxtFmt := "💧 %2.2f %%"
	fanTxtFmt := "💨 %03d %%"

	statusDraw := gg.NewContext(width, height)
	statusDraw.SetColor(backgroundColor)
	statusDraw.Fill()

	statusDraw.SetColor(textColor)
	statusDraw.DrawLine(float64(width/2), 0, float64(width/2), float64(height))
	statusDraw.DrawLine(0, float64(height/2), float64(width), float64(height/2))

	font, err := truetype.Parse(gomono.TTF)
	fontFace := truetype.NewFace(font, &truetype.Options{
		Size: 25,
		// Hinting: font.HintingFull,
	})
	statusDraw.SetFontFace(fontFace)

	statusDraw.DrawString(fmt.Sprintf(temperatureTxtFmt, 18.7), 0, 10)
	statusDraw.DrawString(fmt.Sprintf(humidityTxtFmt, 18.7), 0, 50)
	statusDraw.DrawString(fmt.Sprintf(fanTxtFmt, 18), 0, 90)

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
