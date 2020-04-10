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
		LineDashes:        0,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 10,
		LineColor:         color.Gray{},
		BorderColor:       color.Gray{},
		BackgroundColor:   color.White,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}
