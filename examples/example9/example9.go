package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  1000,
		Height: 1200,
		Name:   "example9.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:            6,
		Columns:         5,
		LineStrokeWidth: 2,
		BackgroundColor: color.RGBA{R: 135, G: 211, B: 124, A: 255},
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	headerStringConfig := gridder.StringConfig{FontSize: 100}
	valueStringConfig := gridder.StringConfig{FontSize: 50}

	headers := []string{"B", "I", "N", "G", "O"}
	values := [][]string{
		{"10", "22", "41", "53", "71"},
		{"66", "20", "40", "50", "2"},
		{"14", "26", "FREE", "52", "69"},
		{"15", "29", "37", "51", "65"},
		{"17", "6", "35", "55", "64"},
	}

	circleConfig := gridder.CircleConfig{Radius: 60, Color: color.White}
	for i, header := range headers {
		grid.DrawCircle(0, i, circleConfig)
		grid.DrawString(0, i, header, headerStringConfig)
	}

	for row := range values {
		for column := range values[0] {
			grid.PaintCell(row+1, column, color.White)
			grid.DrawString(row+1, column, values[row][column], valueStringConfig)
		}
	}
	grid.SavePNG()
}
