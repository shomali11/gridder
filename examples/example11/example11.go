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
		Name:   "example11.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              3,
		Columns:           3,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 4,
		LineColor:         color.Gray{},
		BorderColor:       color.Gray{},
		BackgroundColor:   color.NRGBA{R: 220, G: 220, B: 220, A: 255},
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	blue := color.RGBA{B: 128, A: 255}
	green := color.RGBA{G: 128, A: 255}

	grid.DrawCircle(0, 0, gridder.CircleConfig{Radius: 60, Color: blue, StrokeWidth: 4, Stroke: true})
	grid.DrawCircle(1, 1, gridder.CircleConfig{Radius: 60, Color: blue, StrokeWidth: 4, Stroke: true})
	grid.DrawCircle(2, 2, gridder.CircleConfig{Radius: 60, Color: blue, StrokeWidth: 4, Stroke: true})
	grid.DrawLine(2, 0, gridder.LineConfig{Length: 120, Color: green, StrokeWidth: 4, Rotate: 45})
	grid.DrawLine(2, 0, gridder.LineConfig{Length: 120, Color: green, StrokeWidth: 4, Rotate: 135})
	grid.DrawLine(2, 1, gridder.LineConfig{Length: 120, Color: green, StrokeWidth: 4, Rotate: 45})
	grid.DrawLine(2, 1, gridder.LineConfig{Length: 120, Color: green, StrokeWidth: 4, Rotate: 135})
	grid.DrawPath(0, 0, 2, 2, gridder.PathConfig{StrokeWidth: 4, Color: color.RGBA{R: 128, A: 255}})
	grid.SavePNG()
}
