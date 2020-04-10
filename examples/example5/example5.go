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
		Name:   "example5.png",
	}
	gridConfig := gridder.GridConfig{Rows: 4, Columns: 4}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawRectangle(0, 0, gridder.RectangleConfig{Width: 120, Height: 120, Color: color.Black, Stroke: true})
	grid.DrawRectangle(0, 3, gridder.RectangleConfig{Width: 120, Height: 120, Color: color.Black, Stroke: true, StrokeWidth: 50})
	grid.DrawRectangle(3, 3, gridder.RectangleConfig{Width: 120, Height: 120, Color: color.Black, Stroke: false})
	grid.SavePNG()
}
