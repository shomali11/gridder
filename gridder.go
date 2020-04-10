package gridder

import (
	"errors"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

var (
	errNoRows    = errors.New("no rows provided")
	errNoColumns = errors.New("no columns provided")
)

// New creates a new gridder and sets it up with its configuration
func New(imageConfig ImageConfig, gridConfig GridConfig) (*Gridder, error) {
	gridder := Gridder{
		imageConfig: imageConfig,
		gridConfig:  gridConfig,
		ctx:         gg.NewContext(imageConfig.GetWidth(), imageConfig.GetHeight()),
	}
	gridder.paintBackground()
	err := gridder.paintGrid()
	if err != nil {
		return nil, err
	}
	gridder.paintBorder()
	return &gridder, nil
}

// Gridder gridder structure
type Gridder struct {
	imageConfig ImageConfig
	gridConfig  GridConfig
	ctx         *gg.Context
}

// SavePNG saves to PNG
func (g *Gridder) SavePNG() error {
	return g.ctx.SavePNG(g.imageConfig.GetName())
}

// PaintCell paints Cell
func (g *Gridder) PaintCell(row int, column int, color color.Color) error {
	cellWidth, err := g.getCellWidth()
	if err != nil {
		return err
	}

	cellHeight, err := g.getCellHeight()
	if err != nil {
		return err
	}

	paintWidth := cellWidth - g.gridConfig.GetLineStrokeWidth()*2
	paintHeight := cellHeight - g.gridConfig.GetLineStrokeWidth()*2
	return g.DrawRectangle(row, column, RectangleConfig{Width: paintWidth, Height: paintHeight, Color: color})
}

// DrawRectangle draws a rectangle in a cell
func (g *Gridder) DrawRectangle(row int, column int, rectangleConfigs ...RectangleConfig) error {
	center, err := g.getCellCenter(row, column)
	if err != nil {
		return err
	}

	rectangleConfig := getFirstRectangleConfig(rectangleConfigs...)
	rectangleWidth := rectangleConfig.GetWidth()
	rectangleHeight := rectangleConfig.GetHeight()

	x := center.X - rectangleWidth/2
	y := center.Y - rectangleHeight/2

	g.ctx.DrawRectangle(x, y, rectangleWidth, rectangleHeight)
	g.ctx.SetLineWidth(rectangleConfig.GetStrokeWidth())
	g.ctx.SetColor(rectangleConfig.GetColor())

	if rectangleConfig.IsStroke() {
		g.ctx.Stroke()
	} else {
		g.ctx.Fill()
	}
	return nil
}

// DrawCircle draws a circle in a cell
func (g *Gridder) DrawCircle(row int, column int, circleConfigs ...CircleConfig) error {
	center, err := g.getCellCenter(row, column)
	if err != nil {
		return err
	}

	circleConfig := getFirstCircleConfig(circleConfigs...)
	g.ctx.DrawPoint(center.X, center.Y, circleConfig.GetRadius())
	g.ctx.SetLineWidth(circleConfig.GetStrokeWidth())
	g.ctx.SetColor(circleConfig.GetColor())

	if circleConfig.IsStroke() {
		g.ctx.Stroke()
	} else {
		g.ctx.Fill()
	}
	return nil
}

// DrawLine draws a line between two cells
func (g *Gridder) DrawLine(row1 int, column1 int, row2 int, column2 int, lineConfigs ...LineConfig) error {
	center1, err := g.getCellCenter(row1, column1)
	if err != nil {
		return err
	}

	center2, err := g.getCellCenter(row2, column2)
	if err != nil {
		return err
	}

	lineConfig := getFirstLineConfig(lineConfigs...)
	dashes := lineConfig.GetDashes()
	if dashes > 0 {
		g.ctx.SetDash(dashes)
	} else {
		g.ctx.SetDash()
	}

	g.ctx.SetColor(lineConfig.GetColor())
	g.ctx.SetLineWidth(lineConfig.GetStrokeWidth())
	g.ctx.DrawLine(center1.X, center1.Y, center2.X, center2.Y)
	g.ctx.Stroke()
	return nil
}

// DrawString draws a string in a cell
func (g *Gridder) DrawString(row int, column int, text string, stringConfigs ...StringConfig) error {
	center, err := g.getCellCenter(row, column)
	if err != nil {
		return err
	}

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}

	stringConfig := getFirstStringConfig(stringConfigs...)
	g.ctx.SetFontFace(truetype.NewFace(font, &truetype.Options{Size: stringConfig.GetFontSize()}))
	g.ctx.SetColor(stringConfig.GetColor())
	g.ctx.DrawStringAnchored(text, center.X, center.Y, 0.5, 0.35)
	return nil
}

func (g *Gridder) paintBackground() {
	g.ctx.SetColor(g.gridConfig.GetBackgroundColor())
	g.ctx.Clear()
}

func (g *Gridder) paintGrid() error {
	imageWidth := float64(g.imageConfig.GetWidth())
	imageHeight := float64(g.imageConfig.GetHeight())

	cellWidth, err := g.getCellWidth()
	if err != nil {
		return err
	}

	cellHeight, err := g.getCellHeight()
	if err != nil {
		return err
	}

	columns := float64(g.gridConfig.GetColumns())
	for i := 1.0; i <= columns; i++ {
		x := i * cellWidth
		g.ctx.MoveTo(x, 0)
		g.ctx.LineTo(x, imageHeight)
	}

	rows := float64(g.gridConfig.GetRows())
	for i := 1.0; i <= rows; i++ {
		y := i * cellHeight
		g.ctx.MoveTo(0, y)
		g.ctx.LineTo(imageWidth, y)
	}

	g.ctx.SetColor(g.gridConfig.GetLineColor())
	g.ctx.SetLineWidth(g.gridConfig.GetLineStrokeWidth())
	g.ctx.Stroke()
	return nil
}

func (g *Gridder) paintBorder() {
	width := float64(g.imageConfig.GetWidth())
	height := float64(g.imageConfig.GetHeight())

	g.ctx.DrawRectangle(0, 0, width, height)
	g.ctx.SetLineWidth(g.gridConfig.GetBorderStrokeWidth())
	g.ctx.SetColor(g.gridConfig.GetBorderColor())
	g.ctx.Stroke()
}

func (g *Gridder) getCellWidth() (float64, error) {
	columns := g.gridConfig.GetColumns()
	if columns == 0 {
		return 0, errNoColumns
	}
	return float64(g.imageConfig.GetWidth()) / float64(g.gridConfig.GetColumns()), nil
}

func (g *Gridder) getCellHeight() (float64, error) {
	rows := g.gridConfig.GetRows()
	if rows == 0 {
		return 0, errNoRows
	}
	return float64(g.imageConfig.GetHeight()) / float64(g.gridConfig.GetRows()), nil
}

func (g *Gridder) getCellCenter(row, column int) (*gg.Point, error) {
	columns := float64(g.gridConfig.GetColumns())
	if columns == 0 {
		return nil, errNoColumns
	}

	rows := float64(g.gridConfig.GetRows())
	if rows == 0 {
		return nil, errNoRows
	}

	cellWidth, err := g.getCellWidth()
	if err != nil {
		return nil, err
	}

	cellHeight, err := g.getCellHeight()
	if err != nil {
		return nil, err
	}

	imageWidth := float64(g.imageConfig.GetWidth())
	imageHeight := float64(g.imageConfig.GetHeight())

	x := float64(column)*(imageWidth/columns) + cellWidth/2
	y := float64(row)*(imageHeight/rows) + cellHeight/2
	return &gg.Point{X: x, Y: y}, nil
}
