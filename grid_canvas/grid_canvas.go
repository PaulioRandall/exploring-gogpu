package grid_canvas

import (
	"fmt"

	"github.com/gogpu/gg"

	_ev "github.com/gogpu/ui/event"
	_geo "github.com/gogpu/ui/geometry"
	_prim "github.com/gogpu/ui/primitives"
	_w "github.com/gogpu/ui/widget"
)

type GridCanvas struct {
	_w.WidgetBase
	drawCtx *gg.Context
	points  *[]_geo.Point
}

func Make(drawCtx *gg.Context) GridCanvas {
	gc := GridCanvas{
		drawCtx: drawCtx,
		points:  &[]_geo.Point{},
	}

	gc.SetVisible(true)
	gc.SetEnabled(true)

	return gc
}

func (gc GridCanvas) AsWidget() _w.Widget {
	return _w.Widget(gc)
}

func (gc GridCanvas) Layout(ctx _w.Context, constraints _geo.Constraints) _geo.Size {
	for i, c := range gc.WidgetBase.Children() {
		size := c.Layout(ctx, constraints)
		point := (*gc.points)[i]
		rect := _geo.FromCenter(point, size)
		c.SetBounds(rect)
	}

	return _geo.Expand().Biggest()
}

func (gc GridCanvas) Draw(ctx _w.Context, canvas _w.Canvas) {
	for _, c := range gc.WidgetBase.Children() {
		c.Draw(ctx, canvas)
	}
}

func (gc GridCanvas) Event(ctx _w.Context, ev _ev.Event) bool {
	switch evType := ev.(type) {
	case *_ev.MouseEvent:
		return gc.handleClick(evType)
	default:
		return false
	}
}

func (gc GridCanvas) handleClick(ev *_ev.MouseEvent) bool {
	pos := ev.Position
	println(pos.String())

	txt := fmt.Sprintf("%v:%v", pos.X, pos.Y)
	white := _w.Hex(0xFFFFFF)
	child := _prim.Box(
		_prim.Text(txt).FontSize(12).Color(white),
	)

	*gc.points = append(*gc.points, pos)
	gc.WidgetBase.AddChild(child)
	//gc.WidgetBase.SetNeedsRedraw(true)
	return true
}

func (gc GridCanvas) Children() []_w.Widget {
	return gc.WidgetBase.Children()
}

func (gc *GridCanvas) Close() {
	// TODO
}
