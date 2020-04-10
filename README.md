# gridder [![Build Status](https://travis-ci.com/shomali11/gridder.svg?branch=master)](https://travis-ci.com/shomali11/gridder) [![Go Report Card](https://goreportcard.com/badge/github.com/shomali11/gridder)](https://goreportcard.com/report/github.com/shomali11/gridder) [![GoDoc](https://godoc.org/github.com/shomali11/gridder?status.svg)](https://godoc.org/github.com/shomali11/gridder) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Built on top of Go Graphics [github.com/fogleman/gg](https://github.com/fogleman/gg) with the idea to simplify visualizing Grids using 2D Graphics.

## Dependencies

- `gg` [github.com/fogleman/gg](https://github.com/fogleman/gg)

# Install

```
go get github.com/shomali11/gridder
```

# Examples

## Example 1

Exploring the `GridConfig` configuration object. The `Rows` and `Columns` represent the rows and columns of the grid you wish to visualize. You can customize the grid by playing various options from style, color and width.

```go
package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imageConfig := gridder.ImageConfig{Name: "example1.png"}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           8,
		MarginWidth:       0,
		LineDashes:        0,
		LineStrokeWidth:   2,
		BorderDashes:      0,
		BorderStrokeWidth: 10,
		LineColor:         color.Gray{},
		BorderColor:       color.RGBA{B: 255, A: 255},
		BackgroundColor:   color.White,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}
```

![Result](examples/example1/example1.png)


## Example 2

Exploring the `ImageConfig` configuration object. This defines the image that will be generated. You can customize the `width` and `height` to generate larger images.

```go
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
		Name:   "example2.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              4,
		Columns:           4,
		MarginWidth:       20,
		LineDashes:        10,
		LineStrokeWidth:   2,
		BorderDashes:      20,
		BorderStrokeWidth: 4,
		LineColor:         color.RGBA{R: 255 / 2, A: 255},
		BorderColor:       color.RGBA{B: 255 / 2, A: 255},
		BackgroundColor:   color.RGBA{G: 255 / 2, A: 255},
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	grid.SavePNG()
}
```

![Result](examples/example2/example2.png)
