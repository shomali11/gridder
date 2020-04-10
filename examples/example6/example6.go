package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  1024,
		Height: 1024,
		Name:   "example6.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:            4,
		Columns:         4,
		LineStrokeWidth: 2,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawLine(0, 0, 1, 1, gridder.LineConfig{Dashes: 0, StrokeWidth: 10, Color: color.Black})
	grid.DrawLine(1, 1, 2, 1, gridder.LineConfig{Dashes: 5, StrokeWidth: 1, Color: color.Black})
	grid.DrawLine(2, 1, 3, 1, gridder.LineConfig{Dashes: 8, StrokeWidth: 1, Color: color.Black})
	grid.DrawLine(3, 1, 3, 2, gridder.LineConfig{Dashes: 0, StrokeWidth: 1, Color: color.Black})
	grid.SavePNG()
}
