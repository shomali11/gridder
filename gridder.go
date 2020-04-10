package gridder

import (
	"errors"
	"image/color"

	"github.com/fogleman/gg"
	"golang.org/x/image/font"
)

var (
	errNoRows      = errors.New("no rows provided")
	errNoColumns   = errors.New("no columns provided")
	errOutOfBounds = errors.New("out of bounds")
)

// New creates a new gridder and sets it up with its configuration
func New(imageConfig ImageConfig, gridConfig GridConfig) (*Gridder, error) {
	rows := gridConfig.GetRows()
	if rows == 0 {
		return nil, errNoRows
	}

	columns := gridConfig.GetColumns()
	if columns == 0 {
		return nil, errNoColumns
	}

	gridder := Gridder{
		imageConfig: imageConfig,
		gridConfig:  gridConfig,
		ctx:         gg.NewContext(imageConfig.GetWidth(), imageConfig.GetHeight()),
	}
	gridder.paintBackground()
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
	g.paintGrid()
	g.paintBorder()
	return g.ctx.SavePNG(g.imageConfig.GetName())
}

// PaintCell paints Cell
func (g *Gridder) PaintCell(row int, column int, color color.Color) error {
	err := g.verifyInBounds(row, column)
	if err != nil {
		return err
	}

	cellWidth, cellHeight := g.getCellDimensions()
	paintWidth := cellWidth - g.gridConfig.GetLineStrokeWidth()
	paintHeight := cellHeight - g.gridConfig.GetLineStrokeWidth()
	return g.DrawRectangle(row, column, RectangleConfig{Width: paintWidth, Height: paintHeight, Color: color})
}

// DrawRectangle draws a rectangle in a cell
func (g *Gridder) DrawRectangle(row int, column int, rectangleConfigs ...RectangleConfig) error {
	err := g.verifyInBounds(row, column)
	if err != nil {
		return err
	}

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
	err := g.verifyInBounds(row, column)
	if err != nil {
		return err
	}

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
	err := g.verifyInBounds(row1, column1)
	if err != nil {
		return err
	}

	err = g.verifyInBounds(row2, column2)
	if err != nil {
		return err
	}

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
func (g *Gridder) DrawString(row int, column int, text string, fontFace font.Face, stringConfigs ...StringConfig) error {
	err := g.verifyInBounds(row, column)
	if err != nil {
		return err
	}

	center, err := g.getCellCenter(row, column)
	if err != nil {
		return err
	}

	stringConfig := getFirstStringConfig(stringConfigs...)
	g.ctx.SetFontFace(fontFace)
	g.ctx.SetColor(stringConfig.GetColor())
	g.ctx.DrawStringAnchored(text, center.X, center.Y, 0.5, 0.35)
	return nil
}

func (g *Gridder) paintBackground() {
	margin := float64(g.gridConfig.GetMarginWidth())
	g.ctx.Translate(margin, margin)
	g.ctx.SetColor(g.gridConfig.GetBackgroundColor())
	g.ctx.Clear()
}

func (g *Gridder) paintGrid() {
	canvasWidth, canvasHeight := g.getGridDimensions()
	cellWidth, cellHeight := g.getCellDimensions()

	columns := float64(g.gridConfig.GetColumns())
	for i := 1.0; i < columns; i++ {
		x := i * cellWidth
		g.ctx.MoveTo(x, 0)
		g.ctx.LineTo(x, canvasHeight)
	}

	rows := float64(g.gridConfig.GetRows())
	for i := 1.0; i < rows; i++ {
		y := i * cellHeight
		g.ctx.MoveTo(0, y)
		g.ctx.LineTo(canvasWidth, y)
	}

	dashes := g.gridConfig.GetLineDashes()
	if dashes > 0 {
		g.ctx.SetDash(dashes)
	} else {
		g.ctx.SetDash()
	}
	g.ctx.SetColor(g.gridConfig.GetLineColor())
	g.ctx.SetLineWidth(g.gridConfig.GetLineStrokeWidth())
	g.ctx.Stroke()
}

func (g *Gridder) paintBorder() {
	canvasWidth, canvasHeight := g.getGridDimensions()
	cellWidth, cellHeight := g.getCellDimensions()

	columns := float64(g.gridConfig.GetColumns())
	g.ctx.MoveTo(0, 0)
	g.ctx.LineTo(0, canvasHeight)
	g.ctx.MoveTo(cellWidth*columns, 0)
	g.ctx.LineTo(cellWidth*columns, canvasHeight)

	rows := float64(g.gridConfig.GetRows())
	g.ctx.MoveTo(0, 0)
	g.ctx.LineTo(canvasWidth, 0)
	g.ctx.MoveTo(0, cellHeight*rows)
	g.ctx.LineTo(canvasWidth, cellHeight*rows)

	dashes := g.gridConfig.GetBorderDashes()
	if dashes > 0 {
		g.ctx.SetDash(dashes)
	} else {
		g.ctx.SetDash()
	}
	g.ctx.SetLineWidth(g.gridConfig.GetBorderStrokeWidth())
	g.ctx.SetColor(g.gridConfig.GetBorderColor())
	g.ctx.Stroke()
}

func (g *Gridder) getCellDimensions() (float64, float64) {
	gridWidth, gridHeight := g.getGridDimensions()
	cellWidth := gridWidth / float64(g.gridConfig.GetColumns())
	cellHeight := gridHeight / float64(g.gridConfig.GetRows())
	return cellWidth, cellHeight
}

func (g *Gridder) getGridDimensions() (float64, float64) {
	imageWidth := g.imageConfig.GetWidth()
	imageHeight := g.imageConfig.GetHeight()

	gridWidth := float64(g.gridConfig.GetWidth(imageWidth))
	gridHeight := float64(g.gridConfig.GetHeight(imageHeight))
	return gridWidth, gridHeight
}

func (g *Gridder) getCellCenter(row, column int) (*gg.Point, error) {
	columns := float64(g.gridConfig.GetColumns())
	rows := float64(g.gridConfig.GetRows())

	cellWidth, cellHeight := g.getCellDimensions()
	gridWidth, gridHeight := g.getGridDimensions()

	x := float64(column)*(gridWidth/columns) + cellWidth/2
	y := float64(row)*(gridHeight/rows) + cellHeight/2
	return &gg.Point{X: x, Y: y}, nil
}

func (g *Gridder) verifyInBounds(row, column int) error {
	columns := g.gridConfig.GetColumns()
	rows := g.gridConfig.GetRows()
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return errOutOfBounds
	}
	return nil
}
