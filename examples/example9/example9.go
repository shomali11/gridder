package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{Name: "example9.png"}
	gridConfig := gridder.GridConfig{Rows: 4, Columns: 8}
	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	lineConfig := gridder.LineConfig{Dashes: 10}
	stringConfig := gridder.StringConfig{FontSize: 48}
	circleConfig := gridder.CircleConfig{Color: color.Gray{}, Radius: 10}

	grid.PaintCell(1, 2, color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 2})
	grid.DrawString(1, 2, "Block", stringConfig)

	grid.DrawCircle(0, 0, gridder.CircleConfig{Color: color.NRGBA{R: 255 / 2, G: 0, B: 0, A: 255 / 2}, Radius: 60})
	grid.DrawLine(0, 0, 1, 1, lineConfig)
	grid.DrawCircle(1, 1, circleConfig)
	grid.DrawLine(1, 1, 2, 2, lineConfig)
	grid.DrawCircle(2, 2, circleConfig)
	grid.DrawLine(2, 2, 2, 3, lineConfig)
	grid.DrawCircle(2, 3, circleConfig)
	grid.DrawLine(2, 3, 2, 4, lineConfig)
	grid.DrawCircle(2, 4, circleConfig)
	grid.DrawLine(2, 4, 2, 5, lineConfig)
	grid.DrawCircle(2, 5, circleConfig)
	grid.DrawLine(2, 5, 2, 6, lineConfig)
	grid.DrawCircle(2, 6, circleConfig)
	grid.DrawLine(2, 6, 3, 7, lineConfig)
	grid.DrawCircle(3, 7, gridder.CircleConfig{Color: color.NRGBA{R: 0, G: 255 / 2, B: 0, A: 255 / 2}, Radius: 60})

	grid.SavePNG()
}
