package gridder

import (
	"image/color"
)

const (
	defaultGridWidth             = 500 * 2
	defaultGridHeight            = 500
	defaultGridMarginWidth       = 0
	defaultGridLineStrokeWidth   = 2
	defaultGridBorderStrokeWidth = 4

	defaultLineStrokeWidth = 1

	defaultCircleRadius      = 10
	defaultCircleStrokeWidth = 1

	defaultRectangleWidth       = 20
	defaultRectangleHeight      = 20
	defaultRectangleStrokeWidth = 1
)

var (
	defaultGridBackgroundColor = color.White
	defaultGridBorderColor     = color.Black
	defaultGridLineColor       = color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 4}

	defaultStringColor    = color.Gray{}
	defaultLineColor      = color.Gray{}
	defaultCircleColor    = color.Gray{}
	defaultRectangleColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 2}
)

// ImageConfig Grid Configuration
type ImageConfig struct {
	Width  int
	Height int
	Name   string
}

// GetWidth gets image width
func (g *ImageConfig) GetWidth() int {
	if g.Width <= 0 {
		return defaultGridWidth
	}
	return g.Width
}

// GetHeight gets image height
func (g *ImageConfig) GetHeight() int {
	if g.Height <= 0 {
		return defaultGridHeight
	}
	return g.Height
}

// GetName gets image name
func (g *ImageConfig) GetName() string {
	return g.Name
}

// GridConfig Grid Configuration
type GridConfig struct {
	Rows              int
	Columns           int
	MarginWidth       int
	LineDashes        float64
	LineStrokeWidth   float64
	BorderDashes      float64
	BorderStrokeWidth float64
	LineColor         color.Color
	BorderColor       color.Color
	BackgroundColor   color.Color
}

// GetWidth gets grid width
func (g *GridConfig) GetWidth(imageWidth int) int {
	return imageWidth - g.GetMarginWidth()*2
}

// GetHeight gets grid height
func (g *GridConfig) GetHeight(imageHeight int) int {
	return imageHeight - g.GetMarginWidth()*2
}

// GetMarginWidth gets margin width
func (g *GridConfig) GetMarginWidth() int {
	if g.MarginWidth < 0 {
		return defaultGridMarginWidth
	}
	return g.MarginWidth
}

// GetLineDashes gets line dashes
func (g *GridConfig) GetLineDashes() float64 {
	return g.LineDashes
}

// GetLineStrokeWidth gets line stroke width
func (g *GridConfig) GetLineStrokeWidth() float64 {
	if g.LineStrokeWidth < 0 {
		return defaultGridLineStrokeWidth
	}
	return g.LineStrokeWidth
}

// GetBorderDashes gets border dashes
func (g *GridConfig) GetBorderDashes() float64 {
	return g.BorderDashes
}

// GetBorderStrokeWidth gets border stroke width
func (g *GridConfig) GetBorderStrokeWidth() float64 {
	if g.BorderStrokeWidth < 0 {
		return defaultGridBorderStrokeWidth
	}
	return g.BorderStrokeWidth
}

// GetLineColor gets line color
func (g *GridConfig) GetLineColor() color.Color {
	if g.LineColor == nil {
		return defaultGridLineColor
	}
	return g.LineColor
}

// GetBorderColor gets border color
func (g *GridConfig) GetBorderColor() color.Color {
	if g.BorderColor == nil {
		return defaultGridBorderColor
	}
	return g.BorderColor
}

// GetBackgroundColor gets background color
func (g *GridConfig) GetBackgroundColor() color.Color {
	if g.BackgroundColor == nil {
		return defaultGridBackgroundColor
	}
	return g.BackgroundColor
}

// GetRows gets rows
func (g *GridConfig) GetRows() int {
	if g.Rows < 0 {
		return 0
	}
	return g.Rows
}

// GetColumns gets columns
func (g *GridConfig) GetColumns() int {
	if g.Columns < 0 {
		return 0
	}
	return g.Columns
}

// LineConfig Line Configuration
type LineConfig struct {
	StrokeWidth float64
	Dashes      float64
	Color       color.Color
}

// GetStrokeWidth gets stroke width
func (g *LineConfig) GetStrokeWidth() float64 {
	if g.StrokeWidth <= 0 {
		return defaultLineStrokeWidth
	}
	return g.StrokeWidth
}

// GetColor gets color
func (g *LineConfig) GetColor() color.Color {
	if g.Color == nil {
		return defaultLineColor
	}
	return g.Color
}

// GetDashes gets dashes
func (g *LineConfig) GetDashes() float64 {
	return g.Dashes
}

// CircleConfig Grid Circle Configuration
type CircleConfig struct {
	Radius      float64
	Color       color.Color
	Stroke      bool
	StrokeWidth float64
}

// GetRadius gets radius
func (g *CircleConfig) GetRadius() float64 {
	if g.Radius <= 0 {
		return defaultCircleRadius
	}
	return g.Radius
}

// GetColor gets color
func (g *CircleConfig) GetColor() color.Color {
	if g.Color == nil {
		return defaultCircleColor
	}
	return g.Color
}

// IsStroke determines if Stroke or Fill
func (g *CircleConfig) IsStroke() bool {
	return g.Stroke
}

// GetStrokeWidth gets stroke width
func (g *CircleConfig) GetStrokeWidth() float64 {
	if g.StrokeWidth <= 0 {
		return defaultCircleStrokeWidth
	}
	return g.StrokeWidth
}

// RectangleConfig Rectangle Configuration
type RectangleConfig struct {
	Width       float64
	Height      float64
	Color       color.Color
	Stroke      bool
	StrokeWidth float64
}

// GetWidth gets width
func (g *RectangleConfig) GetWidth() float64 {
	if g.Width <= 0 {
		return defaultRectangleWidth
	}
	return g.Width
}

// GetHeight gets height
func (g *RectangleConfig) GetHeight() float64 {
	if g.Height <= 0 {
		return defaultRectangleHeight
	}
	return g.Height
}

// GetColor gets color
func (g *RectangleConfig) GetColor() color.Color {
	if g.Color == nil {
		return defaultRectangleColor
	}
	return g.Color
}

// IsStroke determines if Stroke or Fill
func (g *RectangleConfig) IsStroke() bool {
	return g.Stroke
}

// GetStrokeWidth gets stroke width
func (g *RectangleConfig) GetStrokeWidth() float64 {
	if g.StrokeWidth <= 0 {
		return defaultRectangleStrokeWidth
	}
	return g.StrokeWidth
}

// StringConfig Grid String Configuration
type StringConfig struct {
	Color color.Color
}

// GetColor gets color
func (g *StringConfig) GetColor() color.Color {
	if g.Color == nil {
		return defaultStringColor
	}
	return g.Color
}

func getFirstRectangleConfig(configs ...RectangleConfig) RectangleConfig {
	if len(configs) == 0 {
		return RectangleConfig{}
	}
	return configs[0]
}

func getFirstCircleConfig(configs ...CircleConfig) CircleConfig {
	if len(configs) == 0 {
		return CircleConfig{}
	}
	return configs[0]
}

func getFirstLineConfig(configs ...LineConfig) LineConfig {
	if len(configs) == 0 {
		return LineConfig{}
	}
	return configs[0]
}

func getFirstStringConfig(configs ...StringConfig) StringConfig {
	if len(configs) == 0 {
		return StringConfig{}
	}
	return configs[0]
}
