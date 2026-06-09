package clickable_area

import (
	"fmt"

	_geo "github.com/gogpu/ui/geometry"
	_p "github.com/gogpu/ui/primitives"
	_w "github.com/gogpu/ui/widget"
)

// TODO: How to determine the size from the primitive.Text?
var (
	pointSize         _geo.Size  = _geo.Sz(80, 32)
	pointHalfSize     _geo.Size  = pointSize.Scale(0.5)
	pointCenterOffset _geo.Point = pointHalfSize.ToPoint()
)

type ClickedPoint struct {
	point _geo.Point
	box   *_p.BoxWidget
}

func MakeClickedPoint(point _geo.Point) ClickedPoint {
	text := createPointText(point)
	box := createPointBox(text)

	return ClickedPoint{
		point: point,
		box:   box,
	}
}

func (cp *ClickedPoint) Center() _geo.Point {
	return cp.point.Sub(pointCenterOffset)
}

// No need to fully implement widget.Widget.

func (cp *ClickedPoint) Draw(ctx _w.Context, canvas _w.Canvas) {
	_w.DrawChild(cp.box, ctx, canvas)
	canvas.StrokeRect(cp.box.Bounds(), _w.RGB(1, 0, 0), 2)
}

func (cp *ClickedPoint) Layout(ctx _w.Context, constraints _geo.Constraints) _geo.Size {
	return cp.box.Layout(ctx, constraints)
}

func createPointText(point _geo.Point) *_p.TextWidget {
	label := fmt.Sprintf("%.0f:%.0f", point.X, point.Y)
	color := _w.RGB8(150, 200, 250)
	text := _p.Text(label)

	fontSize := pointHalfSize.Height
	text.FontSize(fontSize)
	text.Color(color)
	text.Align(_p.TextAlignCenter)

	return text
}

func createPointBox(text *_p.TextWidget) *_p.BoxWidget {
	box := _p.Box(text)

	box.Height(pointSize.Height)
	box.Width(pointSize.Width)
	box.CrossAlign(_p.CrossAxisCenter)

	pad := pointSize.Scale(0.25).Height
	box.Padding(pad)

	return box
}
