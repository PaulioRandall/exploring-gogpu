package app

import (
	_e "github.com/gogpu/ui/event"
	_geo "github.com/gogpu/ui/geometry"
	_w "github.com/gogpu/ui/widget"
)

type ClickableArea struct {
	// Embedding WidgetBase which provides common widget
	// functionality.
	_w.WidgetBase
}

func NewClickableArea() *ClickableArea {
	ca := &ClickableArea{}

	// Both these functions are provided by WidgetBase.
	ca.SetVisible(true)
	ca.SetEnabled(true)

	return ca
}

// AsWidget is convinence for type conversion to a
// widget.Widget.
func (ca *ClickableArea) AsWidget() _w.Widget {
	return _w.Widget(ca)
}

// Layout satisfies the widget.Widget interface.
func (ca *ClickableArea) Layout(ctx _w.Context, constraints _geo.Constraints) _geo.Size {
	for _, child := range ca.Children() {
		// Ignore their size because the clickable area will
		// be fullscreen.
		_ = child.Layout(ctx, constraints)
	}

	return constraints.Biggest()
}

// Draw satisfies the widget.Widget interface.
func (ca *ClickableArea) Draw(ctx _w.Context, canvas _w.Canvas) {
	// TODO: Is this correct? There are gg errors.
	ca.drawBackground(canvas)
	ca.drawPoints(ctx, canvas)
}

func (ca *ClickableArea) drawBackground(canvas _w.Canvas) {
	veryDarkGrey := _w.RGB8(30, 30, 30)
	canvas.DrawRect(ca.Bounds(), veryDarkGrey)
}

func (ca *ClickableArea) drawPoints(ctx _w.Context, canvas _w.Canvas) {
	// To avoid the box and its text from rendering half off
	// the screen (clickable area), I'm shifting (clamping)
	// any ClickedPoint so it renders fully visible.

	// Minimum is the zero point, which is top left of the
	// clickable area.
	minPos := _geo.Pt(0, 0)

	// Maximum is inset from the bottom right by the boxes
	// height and width respectively. Imagine a box in the
	// bottom right corner of the screen (clickable area),
	// the top left point of that box is the maximum
	// allowable point.
	maxPos := ca.Bounds().
		BottomRight().
		Sub(pointSize.ToPoint())

	for _, child := range ca.Children() {
		// Should never fail because all children are
		// ClickedPoints.
		cp := WidgetToClickedPoint(child)

		// Clamping the point. If the ClickedPoint's top left
		// coords are not within the allowable range, then it
		// will shifted so it is.
		pos := cp.topLeft.Clamp(minPos, maxPos)

		// I assume this is the standard way to draw absolutely
		// positioned widgets. We could move this logic into
		// ClickedPoint.Draw but passing minPos and maxPos
		// would be a pain, so CBA.

		// 1. We move the canvas to the clamped position.
		canvas.PushTransform(pos)

		// 2. We draw the widget at that position.
		cp.Draw(ctx, canvas)

		// 3. We move the canvas back to its original position.
		canvas.PopTransform()
	}
}

// Event satisfies the widget.Widget interface.
//
// Note that UI app events propagate from the root widget
// down the tree. This is opposite to JavaScript's default
// event dispatching mode which visits the DOM tree's leaf
// nodes first then bubbles back up to the root.
func (ca *ClickableArea) Event(ctx _w.Context, ev _e.Event) bool {
	me, ok := ev.(*_e.MouseEvent)

	if !ok {
		// Not a mouse event so return false to allow further
		// event propagation.
		return false
	}

	if me.MouseType != _e.MousePress {
		// Not a mouse button click event so return false to
		// allow further event propagation.
		return false
	}

	ca.handleMouseClick(ctx, me)

	// We handled the event so return true because we don't
	// want things further down the tree handling the event
	// (although, in this example, there is nothing further
	// down the widget tree that handles mouse click events,
	// it's good practice since it can prevent difficult
	// to track issues if you decide to add new event
	// capturing widgets later on).
	return true
}

// handleMouseClick creates a ClickedPoint from the click
// position then adds it as a child. It removes the oldest
// ClickedPoint first if the number of children will exceed
// an arbitrary maximum.
func (ca *ClickableArea) handleMouseClick(ctx _w.Context, ev *_e.MouseEvent) {
	// Arbitrary, feel free to change the number.
	// Must be positive.
	var maxClickedPoints int = 3

	// Create a new ClickedPoint and append to our list of
	// points.
	cp := MakeClickedPoint(ev.Position)

	if ca.ChildCount() >= maxClickedPoints {
		ca.RemoveChildAt(0)
	}

	// Function provided by WidgetBase.
	ca.AddChild(cp.AsWidget())

	// Force redraw of entire window.
	ctx.Invalidate()
}

// WidgetBase already provides a suitable implementation
// of the Children function to satisfy widget.Widget.
