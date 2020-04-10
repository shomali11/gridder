package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{Name: "example3.png"}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           8,
		LineStrokeWidth:   10,
		LineColor:         color.RGBA{R: 255 / 2, A: 255},
		BorderStrokeWidth: 20,
		BorderColor:       color.RGBA{B: 255 / 2, A: 255},
		BackgroundColor:   color.RGBA{G: 255 / 2, A: 255},
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}
