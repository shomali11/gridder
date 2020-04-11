package main

import (
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/shomali11/gridder"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:  500,
		Height: 500,
		Name:   "example8.png",
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

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	fontFace1 := truetype.NewFace(font, &truetype.Options{Size: 24})
	fontFace2 := truetype.NewFace(font, &truetype.Options{Size: 35})

	grid.DrawString(0, 0, "Hello!", fontFace1)
	grid.DrawString(1, 1, "Hello!", fontFace1, gridder.StringConfig{Rotate: 45})
	grid.DrawString(2, 2, "Hello!", fontFace2, gridder.StringConfig{Color: color.RGBA{B: 255 / 2, A: 255 / 2}})
	grid.SavePNG()
}
