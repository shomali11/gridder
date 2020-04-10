package main

import (
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{
		Width:   1024,
		Height:  1024,
		Padding: 200,
		Name:    "example2.png",
	}
	gridConfig := gridder.GridConfig{Rows: 4, Columns: 4}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}
