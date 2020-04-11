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
		Name:   "example4.png",
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

	grid.DrawRectangle(0, 0, gridder.RectangleConfig{Width: 60, Height: 60, Color: color.Black, Stroke: true, Rotate: 45})
	grid.DrawRectangle(0, 3, gridder.RectangleConfig{Width: 60, Height: 60, Color: color.Black, Stroke: true, StrokeWidth: 25})
	grid.DrawRectangle(2, 1, gridder.RectangleConfig{Width: 180, Height: 180, Color: color.RGBA{R: 255 / 2, A: 255 / 2}})
	grid.DrawRectangle(3, 3, gridder.RectangleConfig{Width: 60, Height: 60, Color: color.Black, Stroke: false})
	grid.SavePNG()
}
