package gridder

import (
	"image/color"
	"testing"

	"github.com/golang/freetype/truetype"
	"github.com/stretchr/testify/assert"
	"golang.org/x/image/font/gofont/goregular"
)

func TestNew(t *testing.T) {
	_, err := New(ImageConfig{}, GridConfig{Rows: 0})
	assert.NotNil(t, err)

	_, err = New(ImageConfig{}, GridConfig{Rows: 10, Columns: 0})
	assert.NotNil(t, err)

	_, err = New(ImageConfig{}, GridConfig{Rows: 10, Columns: 10})
	assert.Nil(t, err)
}

func TestPaintCell(t *testing.T) {
	gridder, err := New(ImageConfig{}, GridConfig{Rows: 1, Columns: 1})
	assert.Nil(t, err)

	err = gridder.PaintCell(-1, -1, color.Black)
	assert.NotNil(t, err)

	err = gridder.PaintCell(0, 0, color.Black)
	assert.Nil(t, err)
}

func TestDrawCircle(t *testing.T) {
	gridder, err := New(ImageConfig{}, GridConfig{Rows: 1, Columns: 1})
	assert.Nil(t, err)

	err = gridder.DrawCircle(-1, -1)
	assert.NotNil(t, err)

	err = gridder.DrawCircle(0, 0)
	assert.Nil(t, err)
}

func TestDrawRectangle(t *testing.T) {
	gridder, err := New(ImageConfig{}, GridConfig{Rows: 1, Columns: 1})
	assert.Nil(t, err)

	err = gridder.DrawRectangle(-1, -1)
	assert.NotNil(t, err)

	err = gridder.DrawRectangle(0, 0)
	assert.Nil(t, err)
}

func TestDrawLine(t *testing.T) {
	gridder, err := New(ImageConfig{}, GridConfig{Rows: 1, Columns: 1})
	assert.Nil(t, err)

	err = gridder.DrawLine(-1, -1)
	assert.NotNil(t, err)

	err = gridder.DrawLine(0, 0)
	assert.Nil(t, err)
}

func TestDrawPath(t *testing.T) {
	gridder, err := New(ImageConfig{}, GridConfig{Rows: 1, Columns: 1})
	assert.Nil(t, err)

	err = gridder.DrawPath(-1, -1, -1, -1)
	assert.NotNil(t, err)

	err = gridder.DrawPath(0, 0, 0, 0)
	assert.Nil(t, err)
}

func TestDrawString(t *testing.T) {
	gridder, err := New(ImageConfig{}, GridConfig{Rows: 1, Columns: 1})
	assert.Nil(t, err)

	font, _ := truetype.Parse(goregular.TTF)
	fontFace := truetype.NewFace(font, &truetype.Options{Size: 24})

	err = gridder.DrawString(-1, -1, "Test", fontFace)
	assert.NotNil(t, err)

	err = gridder.DrawString(0, 0, "Test", fontFace)
	assert.Nil(t, err)
}

func TestSave(t *testing.T) {
	gridder, err := New(ImageConfig{}, GridConfig{Rows: 1, Columns: 1})
	assert.Nil(t, err)

	err = gridder.SavePNG()
	assert.NotNil(t, err)
}
