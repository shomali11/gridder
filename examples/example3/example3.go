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
		Name:   "example3.png",
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

	grid.PaintCell(1, 2, color.Black)
	grid.SavePNG()
}
