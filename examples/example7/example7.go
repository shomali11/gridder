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
		Name:   "example7.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           4,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 4,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawPath(0, 0, 1, 1, gridder.PathConfig{Dashes: 0, StrokeWidth: 10, Color: color.Black})
	grid.DrawPath(1, 1, 2, 1, gridder.PathConfig{Dashes: 5, StrokeWidth: 1, Color: color.Black})
	grid.DrawPath(2, 1, 3, 1, gridder.PathConfig{Dashes: 8, StrokeWidth: 1, Color: color.Black})
	grid.DrawPath(3, 1, 3, 2, gridder.PathConfig{Dashes: 0, StrokeWidth: 1, Color: color.Black})
	grid.SavePNG()
}
