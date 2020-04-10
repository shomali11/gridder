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

	grid.DrawString(1, 1, "Hello!")
	grid.DrawString(2, 2, "Hello!", gridder.StringConfig{FontSize: 35, Color: color.RGBA{B: 255 / 2, A: 255 / 2}})
	grid.SavePNG()
}
