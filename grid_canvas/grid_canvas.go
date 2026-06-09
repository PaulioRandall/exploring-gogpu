package grid_canvas

import (
	_e "github.com/gogpu/ui/event"
	_geo "github.com/gogpu/ui/geometry"
	_w "github.com/gogpu/ui/widget"
)

type GridCanvas struct {
	_w.WidgetBase
}

func NewGridCanvas() *GridCanvas {
	gc := &GridCanvas{}

	gc.SetVisible(true)
	gc.SetEnabled(true)

	return gc
}

func (gc *GridCanvas) AsWidget() _w.Widget {
	return _w.Widget(gc)
}

func (gc *GridCanvas) Layout(ctx _w.Context, constraints _geo.Constraints) _geo.Size {
	for _, _ = range gc.Children() {
		// TODO: What to do here?
	}
	return constraints.Biggest()
}

func (gc *GridCanvas) Draw(ctx _w.Context, canvas _w.Canvas) {
	for _, c := range gc.Children() {
		_w.DrawChild(c, ctx, canvas)
	}
}

func (gc *GridCanvas) Event(ctx _w.Context, ev _e.Event) bool {
	if me, ok := ev.(*_e.MouseEvent); ok {
		if me.MouseType == _e.MousePress {
			return gc.handleClick(ctx, me)
		}
	}

	return false
}

func (gc *GridCanvas) handleClick(ctx _w.Context, ev *_e.MouseEvent) bool {
	println(ev.Position.String())

	p := NewClickedPoint(ev.Position)
	gc.AddChild(p.AsWidget())
	ctx.Invalidate() // Redraw entire window.

	return true
}

func (gc *GridCanvas) Children() []_w.Widget {
	children := gc.WidgetBase.Children()
	count := gc.WidgetBase.ChildCount()

	if count == 0 {
		return nil
	}

	result := make([]_w.Widget, count)
	copy(result, children)
	return result
}

func (gc *GridCanvas) Close() {
	// TODO
}

// CHeck implements interfaces.
var _ _w.Widget = (*GridCanvas)(nil)
