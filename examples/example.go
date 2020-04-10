package main

import (
	"image/color"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{Name: "example1.png"}
	gridConfig := gridder.GridConfig{Rows: 4, Columns: 8}
	grid := gridder.New(imageConfig, gridConfig)

	grid.PaintCell(1, 2, color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 2})
	grid.DrawRectangle(3, 4, gridder.RectangleConfig{Width: 20, Height: 20})

	grid.DrawCircle(0, 0, gridder.CircleConfig{Color: color.NRGBA{R: 255 / 2, G: 0, B: 0, A: 255 / 2}, Radius: 60})
	grid.DrawLine(0, 0, 1, 1, gridder.LineConfig{Dashes: 10})
	grid.DrawCircle(1, 1, gridder.CircleConfig{Color: color.Gray{}, Radius: 10})
	grid.DrawLine(1, 1, 2, 2)
	grid.DrawCircle(2, 2, gridder.CircleConfig{Color: color.Gray{}, Radius: 10})
	grid.DrawLine(2, 2, 2, 3)
	grid.DrawCircle(2, 3)
	grid.DrawLine(2, 3, 2, 4)
	grid.DrawCircle(2, 4)
	grid.DrawLine(2, 4, 2, 5)
	grid.DrawCircle(2, 5, gridder.CircleConfig{Color: color.Gray{}, Radius: 10})
	grid.DrawLine(2, 5, 2, 6)
	grid.DrawCircle(2, 6, gridder.CircleConfig{Color: color.Gray{}, Radius: 10})
	grid.DrawLine(2, 6, 3, 7)
	grid.DrawCircle(3, 7, gridder.CircleConfig{Color: color.NRGBA{R: 0, G: 255 / 2, B: 0, A: 255 / 2}, Radius: 60})

	grid.SavePNG()
}
