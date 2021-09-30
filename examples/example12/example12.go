package main

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  500,
		Height: 500,
		Name:   "example12.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              20,
		Columns:           20,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 4,
		LineColor:         color.Gray{},
		BorderColor:       color.Gray{},
		BackgroundColor:   color.NRGBA{R: 220, G: 220, B: 220, A: 255},
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	blue := color.RGBA{B: 128, A: 255}

	// create a random chart
	rand.Seed(time.Now().UnixNano())
	for col := 0; col < gridConfig.Columns; col++ {
		height := rand.Intn(gridConfig.Rows - 1)
		for topRow := 0; topRow < height; topRow++ {
			grid.DrawCircle(gridConfig.Rows-topRow, col, gridder.CircleConfig{Radius: 5, Color: blue, StrokeWidth: 4, Stroke: true})
		}
	}

	// encode image as byte string
	bImage := new(bytes.Buffer)
	grid.EncodePNG(bImage)

	// convert to base64 string to support storing into database
	imageString := base64.StdEncoding.EncodeToString(bImage.Bytes())

	// convert back from string and save as binary image
	bDecodedImage, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(imageConfig.Name, bDecodedImage, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
