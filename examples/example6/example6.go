package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  500,
		Height: 500,
		Name:   "example6.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           4,
		MarginWidth:       30,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 4,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawLine(0, 0, gridder.LineConfig{Length: 60, Color: color.Black, Dashes: 0})
	grid.DrawLine(0, 0, gridder.LineConfig{Length: 60, Color: color.Black, Dashes: 0, Rotate: 90})
	grid.DrawLine(0, 3, gridder.LineConfig{Length: 60, Color: color.Black, Dashes: 0, StrokeWidth: 25})
	grid.DrawLine(2, 1, gridder.LineConfig{Length: 90, Color: color.RGBA{R: 255 / 2, A: 255 / 2}, Rotate: 45})
	grid.DrawLine(2, 1, gridder.LineConfig{Length: 90, Color: color.RGBA{R: 255 / 2, A: 255 / 2}, Rotate: 135})
	grid.DrawLine(3, 3, gridder.LineConfig{Length: 60, Color: color.Black, Dashes: 5})
	grid.SavePNG()
}
