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
	gridConfig := gridder.GridConfig{Rows: 4, Columns: 4}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.DrawCircle(0, 0, gridder.CircleConfig{Radius: 60, Color: color.Black, Stroke: true})
	grid.DrawCircle(0, 3, gridder.CircleConfig{Radius: 60, Color: color.Black, Stroke: true, StrokeWidth: 50})
	grid.DrawCircle(3, 3, gridder.CircleConfig{Radius: 60, Color: color.Black, Stroke: false})
	grid.SavePNG()
}
