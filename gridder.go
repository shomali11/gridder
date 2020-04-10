package gridder

import (
	"errors"
	"image/color"

	"github.com/fogleman/gg"
)

var (
	errNoRows    = errors.New("no rows provided")
	errNoColumns = errors.New("no columns provided")
)

// New creates a new gridder and sets it up with its configuration
func New(imageConfig ImageConfig, gridConfig GridConfig) *Gridder {
	gridder := Gridder{imageConfig: imageConfig, gridConfig: gridConfig, ctx: gg.NewContext(imageConfig.GetWidth(), imageConfig.GetHeight())}
	gridder.setupCanvas()
	gridder.paintBackground()
	gridder.paintGrid()
	gridder.paintBorder()
	return &gridder
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
func (g *Gridder) PaintCell(x int, y int, color color.Color) error {
	cellWidth, err := g.getCellWidth()
	if err != nil {
		return err
	}

	cellHeight, err := g.getCellHeight()
	if err != nil {
		return err
	}

	w := cellWidth * float64(g.imageConfig.GetWidth())
	h := cellHeight * float64(g.imageConfig.GetHeight())

	return g.DrawRectangle(x, y, RectangleConfig{Width: w, Height: h, Color: color})
}

// DrawRectangle draws a rectangle in a cell
func (g *Gridder) DrawRectangle(x int, y int, rectangleConfigs ...RectangleConfig) error {
	center, err := g.getCellCenter(x, y)
	if err != nil {
		return err
	}

	rectangleConfig := getFirstRectangleConfig(rectangleConfigs...)
	w := rectangleConfig.GetWidth() / float64(g.imageConfig.GetWidth())
	h := rectangleConfig.GetHeight() / float64(g.imageConfig.GetHeight())

	i := center.X - w/2
	j := center.Y - h/2

	g.ctx.DrawRectangle(i, j, w, h)
	g.ctx.SetColor(rectangleConfig.GetColor())

	if rectangleConfig.IsStroke() {
		g.ctx.Stroke()
	} else {
		g.ctx.Fill()
	}
	return nil
}

// DrawCircle draws a circle in a cell
func (g *Gridder) DrawCircle(x int, y int, circleConfigs ...CircleConfig) error {
	center, err := g.getCellCenter(x, y)
	if err != nil {
		return err
	}

	circleConfig := getFirstCircleConfig(circleConfigs...)
	g.ctx.DrawPoint(center.X, center.Y, circleConfig.GetRadius())
	g.ctx.SetColor(circleConfig.GetColor())

	if circleConfig.IsStroke() {
		g.ctx.Stroke()
	} else {
		g.ctx.Fill()
	}
	return nil
}

// DrawLine draws a line between two cells
func (g *Gridder) DrawLine(x1 int, y1 int, x2 int, y2 int, lineConfigs ...LineConfig) error {
	center1, err := g.getCellCenter(x1, y1)
	if err != nil {
		return err
	}

	center2, err := g.getCellCenter(x2, y2)
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
	g.ctx.SetLineWidth(lineConfig.GetWidth())
	g.ctx.DrawLine(center1.X, center1.Y, center2.X, center2.Y)
	g.ctx.Stroke()
	return nil
}

func (g *Gridder) setupCanvas() {
	padding := g.imageConfig.GetPadding()
	width := float64(g.imageConfig.GetWidth())
	height := float64(g.imageConfig.GetHeight())
	g.ctx.Translate(padding, padding)
	g.ctx.Scale(width-padding*2, height-padding*2)
}

func (g *Gridder) paintBackground() {
	g.ctx.SetColor(g.gridConfig.GetBackgroundColor())
	g.ctx.Clear()
}

func (g *Gridder) paintGrid() {
	columns := float64(g.gridConfig.GetColumns())
	for i := 1.0; i <= columns; i++ {
		x := i / columns
		g.ctx.MoveTo(x, 0)
		g.ctx.LineTo(x, 1)
	}

	rows := float64(g.gridConfig.GetRows())
	for i := 1.0; i <= rows; i++ {
		y := i / rows
		g.ctx.MoveTo(0, y)
		g.ctx.LineTo(1, y)
	}

	g.ctx.SetColor(g.gridConfig.GetLineColor())
	g.ctx.SetLineWidth(g.gridConfig.GetLineWidth())
	g.ctx.Stroke()
}

func (g *Gridder) paintBorder() {
	g.ctx.MoveTo(0, 0)
	g.ctx.LineTo(1, 0)

	g.ctx.MoveTo(0, 0)
	g.ctx.LineTo(0, 1)

	g.ctx.MoveTo(1, 1)
	g.ctx.LineTo(1, 0)

	g.ctx.MoveTo(1, 1)
	g.ctx.LineTo(0, 1)

	g.ctx.SetColor(g.gridConfig.GetBorderColor())
	g.ctx.SetLineWidth(g.gridConfig.GetBorderWidth())
	g.ctx.Stroke()
}

func (g *Gridder) getCellWidth() (float64, error) {
	columns := g.gridConfig.GetColumns()
	if columns == 0 {
		return 0, errNoColumns
	}
	return 1 / float64(g.gridConfig.GetColumns()), nil
}

func (g *Gridder) getCellHeight() (float64, error) {
	rows := g.gridConfig.GetRows()
	if rows == 0 {
		return 0, errNoRows
	}
	return 1 / float64(g.gridConfig.GetRows()), nil
}

func (g *Gridder) getCellCenter(x, y int) (*gg.Point, error) {
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

	i := float64(y)/columns + cellWidth/2
	j := float64(x)/rows + cellHeight/2
	return &gg.Point{X: i, Y: j}, nil
}
