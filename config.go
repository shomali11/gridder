package gridder

import "image/color"

const (
	defaultGridPadding     = 32
	defaultGridWidth       = 1024 * 2
	defaultGridHeight      = 1024
	defaultGridLineWidth   = 2
	defaultGridBorderWidth = 4

	defaultLineWidth = 1

	defaultCircleRadius = 10

	defaultRectangleWidth  = 20
	defaultRectangleHeight = 20
)

var (
	defaultGridBackgroundColor = color.White
	defaultGridBorderColor     = color.Black
	defaultGridLineColor       = color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 4}

	defaultLineColor      = color.Gray{}
	defaultCircleColor    = color.Gray{}
	defaultRectangleColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 2}
)

// ImageConfig Grid Configuration
type ImageConfig struct {
	Width   int
	Height  int
	Padding float64
	Name    string
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

// GetPadding gets image padding
func (g *ImageConfig) GetPadding() float64 {
	if g.Padding <= 0 {
		return defaultGridPadding
	}
	return g.Padding
}

// GetName gets image name
func (g *ImageConfig) GetName() string {
	return g.Name
}

// GridConfig Grid Configuration
type GridConfig struct {
	Rows            int
	Columns         int
	LineWidth       float64
	BorderWidth     float64
	LineColor       color.Color
	BorderColor     color.Color
	BackgroundColor color.Color
}

// GetLineWidth gets line width
func (g *GridConfig) GetLineWidth() float64 {
	if g.LineWidth <= 0 {
		return defaultGridLineWidth
	}
	return g.LineWidth
}

// GetBorderWidth gets border width
func (g *GridConfig) GetBorderWidth() float64 {
	if g.BorderWidth <= 0 {
		return defaultGridBorderWidth
	}
	return g.BorderWidth
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
	return g.Rows
}

// GetColumns gets columns
func (g *GridConfig) GetColumns() int {
	return g.Columns
}

// LineConfig Line Configuration
type LineConfig struct {
	Width  float64
	Dashes float64
	Color  color.Color
}

// GetWidth gets width
func (g *LineConfig) GetWidth() float64 {
	if g.Width <= 0 {
		return defaultLineWidth
	}
	return g.Width
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
	Radius float64
	Color  color.Color
	Stroke bool
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

// RectangleConfig Rectangle Configuration
type RectangleConfig struct {
	Width  float64
	Height float64
	Color  color.Color
	Stroke bool
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
