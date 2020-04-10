package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{Name: "example1.png"}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           8,
		MarginWidth:       0,
		LineDashes:        0,
		LineStrokeWidth:   2,
		BorderDashes:      0,
		BorderStrokeWidth: 10,
		LineColor:         color.Gray{},
		BorderColor:       color.RGBA{B: 255, A: 255},
		BackgroundColor:   color.White,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}
