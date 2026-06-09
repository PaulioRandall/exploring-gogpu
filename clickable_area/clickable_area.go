package clickable_area

import (
	_e "github.com/gogpu/ui/event"
	_geo "github.com/gogpu/ui/geometry"
	_w "github.com/gogpu/ui/widget"
)

var maxClickedPoints int = 3

type ClickableArea struct {
	_w.WidgetBase

	clickedPoints []ClickedPoint
}

func NewClickableArea() *ClickableArea {
	ca := &ClickableArea{
		clickedPoints: []ClickedPoint{},
	}

	ca.SetVisible(true)
	ca.SetEnabled(true)

	return ca
}

func (ca *ClickableArea) AsWidget() _w.Widget {
	return _w.Widget(ca)
}

func (ca *ClickableArea) Layout(ctx _w.Context, constraints _geo.Constraints) _geo.Size {
	// The app takes the returned size and sets this widgets
	// bounds.
	// https://github.com/gogpu/ui/blob/543b0ccb53da91664434cf1d96d05b4eb0433d69/app/window.go#L767.

	// Layout all children ignoring size.
	for _, cp := range ca.clickedPoints {
		cp.Layout(ctx, constraints)
	}
	return constraints.Biggest()
}

func (ca *ClickableArea) Draw(ctx _w.Context, canvas _w.Canvas) {
	// TODO: Is this correct? There are gg errors.
	ca.drawBG(canvas)
	ca.drawPoints(ctx, canvas)
}

func (ca *ClickableArea) drawBG(canvas _w.Canvas) {
	veryDarkGrey := _w.RGB8(30, 30, 30)
	canvas.DrawRect(ca.Bounds(), veryDarkGrey)
}

func (ca *ClickableArea) drawPoints(ctx _w.Context, canvas _w.Canvas) {
	minOffset := _geo.Pt(0, 0)
	maxOffset := ca.Bounds().BottomRight().Sub(pointSize.ToPoint())

	for _, cp := range ca.clickedPoints {
		offset := cp.Center().Clamp(minOffset, maxOffset)

		canvas.PushTransform(offset)
		cp.Draw(ctx, canvas)
		canvas.PopTransform()
	}
}

func (ca *ClickableArea) Event(ctx _w.Context, ev _e.Event) bool {
	if me, ok := ev.(*_e.MouseEvent); ok {
		if me.MouseType == _e.MousePress {
			// If a mouse button is pressed.
			return ca.handleClick(ctx, me)
		}
	}

	return false
}

func (ca *ClickableArea) handleClick(ctx _w.Context, ev *_e.MouseEvent) bool {
	cp := MakeClickedPoint(ev.Position)

	if len(ca.clickedPoints) < maxClickedPoints {
		ca.clickedPoints = append(ca.clickedPoints, cp)
	} else {
		ca.clickedPoints = append(ca.clickedPoints[1:], cp)
	}

	ctx.Invalidate() // Redraw entire window.
	return true
}

// WidgetBase already provides a suitable implementation
// of the Children function.
// func (wb *WidgetBase) Children() []widget.Widget { ... }

// Ensure satisfies Widget & Lifecycle interface.
var _ _w.Widget = (*ClickableArea)(nil)
