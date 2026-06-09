package grid_canvas

import (
	"fmt"

	_e "github.com/gogpu/ui/event"
	_geo "github.com/gogpu/ui/geometry"
	_p "github.com/gogpu/ui/primitives"
	_w "github.com/gogpu/ui/widget"
)

type ClickedPoint struct {
	point _geo.Point
	box   *_p.BoxWidget
}

func NewClickedPoint(p _geo.Point) *ClickedPoint {
	label := fmt.Sprintf("%v:%v", p.X, p.Y)
	color := _w.RGB(0.5, 0.5, 0.5)
	text := _p.Text(label).FontSize(12).Color(color)

	return &ClickedPoint{
		point: p,
		box:   _p.Box(text),
	}
}

func (cp *ClickedPoint) AsWidget() _w.Widget {
	return _w.Widget(cp)
}

func (cp *ClickedPoint) Layout(ctx _w.Context, constraints _geo.Constraints) _geo.Size {
	return cp.box.Layout(ctx, constraints)
}

func (cp *ClickedPoint) Draw(ctx _w.Context, canvas _w.Canvas) {
	_w.DrawChild(cp.box, ctx, canvas)
}

func (cp *ClickedPoint) Event(ctx _w.Context, ev _e.Event) bool {
	return false
}

func (cp *ClickedPoint) Children() []_w.Widget {
	return nil
}
