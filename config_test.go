package gridder

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageConfig(t *testing.T) {
	config1 := &ImageConfig{Name: "Hello"}
	assert.Equal(t, config1.GetWidth(), defaultGridWidth)
	assert.Equal(t, config1.GetHeight(), defaultGridHeight)
	assert.Equal(t, config1.GetName(), "Hello")

	config2 := &ImageConfig{Name: "Bye", Width: 10, Height: 100}
	assert.Equal(t, config2.GetWidth(), 10)
	assert.Equal(t, config2.GetHeight(), 100)
	assert.Equal(t, config2.GetName(), "Bye")
}

func TestGridConfig(t *testing.T) {
	config1 := &GridConfig{}
	assert.Equal(t, config1.GetRows(), 0)
	assert.Equal(t, config1.GetColumns(), 0)
	assert.Equal(t, config1.GetMarginWidth(), defaultGridMarginWidth)
	assert.Equal(t, config1.GetWidth(100), 100)
	assert.Equal(t, config1.GetHeight(100), 100)
	assert.Equal(t, config1.GetLineDashes(), 0.0)
	assert.Equal(t, config1.GetLineStrokeWidth(), 0.0)
	assert.Equal(t, config1.GetBorderDashes(), 0.0)
	assert.Equal(t, config1.GetBorderStrokeWidth(), 0.0)
	assert.Equal(t, config1.GetLineColor(), defaultGridLineColor)
	assert.Equal(t, config1.GetBorderColor(), defaultGridBorderColor)
	assert.Equal(t, config1.GetBackgroundColor(), defaultGridBackgroundColor)

	config2 := &GridConfig{
		Rows: 100, Columns: 200, MarginWidth: 1, LineDashes: 1, BorderDashes: 2,
		LineStrokeWidth: 4, BorderStrokeWidth: 8,
		LineColor: color.White, BorderColor: color.White, BackgroundColor: color.White,
	}
	assert.Equal(t, config2.GetRows(), 100)
	assert.Equal(t, config2.GetColumns(), 200)
	assert.Equal(t, config2.GetMarginWidth(), 1)
	assert.Equal(t, config2.GetWidth(100), 98)
	assert.Equal(t, config2.GetHeight(100), 98)
	assert.Equal(t, config2.GetLineDashes(), 1.0)
	assert.Equal(t, config2.GetLineStrokeWidth(), 4.0)
	assert.Equal(t, config2.GetBorderDashes(), 2.0)
	assert.Equal(t, config2.GetBorderStrokeWidth(), 8.0)
	assert.Equal(t, config2.GetLineColor(), color.White)
	assert.Equal(t, config2.GetBorderColor(), color.White)
	assert.Equal(t, config2.GetBackgroundColor(), color.White)
}

func TestPathConfig(t *testing.T) {
	config1 := &PathConfig{}
	assert.Equal(t, config1.GetDashes(), 0.0)
	assert.Equal(t, config1.GetStrokeWidth(), defaultLineStrokeWidth)
	assert.Equal(t, config1.GetColor(), defaultLineColor)

	config2 := &PathConfig{Dashes: 1, StrokeWidth: 10, Color: color.White}
	assert.Equal(t, config2.GetDashes(), 1.0)
	assert.Equal(t, config2.GetStrokeWidth(), 10.0)
	assert.Equal(t, config2.GetColor(), color.White)
}

func TestLineConfig(t *testing.T) {
	config1 := &LineConfig{}
	assert.Equal(t, config1.GetLength(), defaultLineLength)
	assert.Equal(t, config1.GetRotate(), 0.0)
	assert.Equal(t, config1.GetDashes(), 0.0)
	assert.Equal(t, config1.GetStrokeWidth(), defaultLineStrokeWidth)
	assert.Equal(t, config1.GetColor(), defaultLineColor)

	config2 := &LineConfig{Length: 5, Rotate: 90, Dashes: 1, StrokeWidth: 10, Color: color.White}
	assert.Equal(t, config2.GetLength(), 5.0)
	assert.Equal(t, config2.GetRotate(), 90.0)
	assert.Equal(t, config2.GetDashes(), 1.0)
	assert.Equal(t, config2.GetStrokeWidth(), 10.0)
	assert.Equal(t, config2.GetColor(), color.White)
}

func TestCircleConfig(t *testing.T) {
	config1 := &CircleConfig{}
	assert.Equal(t, config1.GetRadius(), defaultCircleRadius)
	assert.Equal(t, config1.GetDashes(), 0.0)
	assert.Equal(t, config1.IsStroke(), false)
	assert.Equal(t, config1.GetStrokeWidth(), defaultCircleStrokeWidth)
	assert.Equal(t, config1.GetColor(), defaultCircleColor)

	config2 := &CircleConfig{Radius: 1, Dashes: 1, Stroke: true, StrokeWidth: 10, Color: color.White}
	assert.Equal(t, config2.GetRadius(), 1.0)
	assert.Equal(t, config2.GetDashes(), 1.0)
	assert.Equal(t, config2.IsStroke(), true)
	assert.Equal(t, config2.GetStrokeWidth(), 10.0)
	assert.Equal(t, config2.GetColor(), color.White)
}

func TestRectangleConfig(t *testing.T) {
	config1 := &RectangleConfig{}
	assert.Equal(t, config1.GetWidth(), defaultRectangleWidth)
	assert.Equal(t, config1.GetHeight(), defaultRectangleHeight)
	assert.Equal(t, config1.GetRotate(), 0.0)
	assert.Equal(t, config1.GetDashes(), 0.0)
	assert.Equal(t, config1.IsStroke(), false)
	assert.Equal(t, config1.GetStrokeWidth(), defaultRectangleStrokeWidth)
	assert.Equal(t, config1.GetColor(), defaultRectangleColor)

	config2 := &RectangleConfig{Width: 1, Height: 2, Rotate: 90, Dashes: 1, Stroke: true, StrokeWidth: 10, Color: color.White}
	assert.Equal(t, config2.GetWidth(), 1.0)
	assert.Equal(t, config2.GetHeight(), 2.0)
	assert.Equal(t, config2.GetRotate(), 90.0)
	assert.Equal(t, config2.GetDashes(), 1.0)
	assert.Equal(t, config2.IsStroke(), true)
	assert.Equal(t, config2.GetStrokeWidth(), 10.0)
	assert.Equal(t, config2.GetColor(), color.White)
}

func TestStringConfig(t *testing.T) {
	config1 := &StringConfig{}
	assert.Equal(t, config1.GetRotate(), 0.0)
	assert.Equal(t, config1.GetColor(), defaultStringColor)

	config2 := &StringConfig{Rotate: 1, Color: color.White}
	assert.Equal(t, config2.GetRotate(), 1.0)
	assert.Equal(t, config2.GetColor(), color.White)
}

func TestFirstRectangleConfig(t *testing.T) {
	config1 := getFirstRectangleConfig()
	assert.Equal(t, config1, RectangleConfig{})

	config2 := getFirstRectangleConfig(config1)
	assert.Equal(t, config2, config1)
}

func TestFirstCircleConfig(t *testing.T) {
	config1 := getFirstCircleConfig()
	assert.Equal(t, config1, CircleConfig{})

	config2 := getFirstCircleConfig(config1)
	assert.Equal(t, config2, config1)
}

func TestFirstLineConfig(t *testing.T) {
	config1 := getFirstLineConfig()
	assert.Equal(t, config1, LineConfig{})

	config2 := getFirstLineConfig(config1)
	assert.Equal(t, config2, config1)
}

func TestFirstPathConfig(t *testing.T) {
	config1 := getFirstPathConfig()
	assert.Equal(t, config1, PathConfig{})

	config2 := getFirstPathConfig(config1)
	assert.Equal(t, config2, config1)
}

func TestFirstStringConfig(t *testing.T) {
	config1 := getFirstStringConfig()
	assert.Equal(t, config1, StringConfig{})

	config2 := getFirstStringConfig(config1)
	assert.Equal(t, config2, config1)
}
