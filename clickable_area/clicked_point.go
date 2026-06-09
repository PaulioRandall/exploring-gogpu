package clickable_area

import (
	"fmt"

	_e "github.com/gogpu/ui/event"
	_geo "github.com/gogpu/ui/geometry"
	_p "github.com/gogpu/ui/primitives"
	_w "github.com/gogpu/ui/widget"
)

// TODO: How to determine the size from the primitive.Text?
var (
	// Change font size as you desire.
	// Other sizes will be calculated accordingly.
	fontSize          float32    = 24
	paddingSize       float32    = fontSize / 2
	pointSize         _geo.Size  = _geo.Sz(fontSize*5, fontSize*2)
	pointHalfSize     _geo.Size  = pointSize.Scale(0.5)
	pointCenterOffset _geo.Point = pointHalfSize.ToPoint()
)

// ClickedPoint represents a clicked point on the clickable
// area. It is a functional container for a box centered
// on the clicked point, with a rounded border, and
// containing the point's coords as centred text.
type ClickedPoint struct {
	// The original clicked point (not actually needed but
	// was useful for debugging).
	point _geo.Point

	// Bordered box that presents the point's coords as text.
	box *_p.BoxWidget

	// Convinence point representing the top left coords
	// for rendering the box.
	topLeft _geo.Point
}

func MakeClickedPoint(point _geo.Point) ClickedPoint {
	text := createPointTextWidget(point)
	box := createPointBoxWidget(text)
	topLeft := point.Sub(pointCenterOffset)

	return ClickedPoint{
		topLeft: topLeft,
		point:   point,
		box:     box,
	}
}

// WidgetToClickedPoint is convinence for type conversion
// to a ClickedPoint from a widget.Widget.
func WidgetToClickedPoint(w _w.Widget) *ClickedPoint {
	return w.(*ClickedPoint)
}

// AsWidget is convinence for type conversion to a
// widget.Widget.
func (cp *ClickedPoint) AsWidget() _w.Widget {
	return _w.Widget(cp)
}

// Draw satisfies the widget.Widget interface.
func (cp *ClickedPoint) Draw(ctx _w.Context, canvas _w.Canvas) {
	// I saw this somewhere and assumed I had to use it
	// instead of calling cp.box.Draw directly.
	_w.DrawChild(cp.box, ctx, canvas)
}

// Layout satisfies the widget.Widget interface.
func (cp *ClickedPoint) Layout(ctx _w.Context, constraints _geo.Constraints) _geo.Size {
	// Return the box size because ClickedPoint is a
	// functional container, not an aesthetic one.
	return cp.box.Layout(ctx, constraints)
}

// Event satisfies the widget.Widget interface.
func (cp *ClickedPoint) Event(ctx _w.Context, ev _e.Event) bool {
	// Return false because we're ignoring all events.
	return false
}

// Children satisfies the widget.Widget interface.
func (cp *ClickedPoint) Children() []_w.Widget {
	// The box is always the only child.
	return []_w.Widget{cp.box}
}

func createPointTextWidget(point _geo.Point) *_p.TextWidget {
	// Print coords without decimal points.
	textLabel := fmt.Sprintf("%.0f:%.0f", point.X, point.Y)
	text := _p.Text(textLabel)

	text.Align(_p.TextAlignCenter)
	text.FontSize(fontSize)

	textColor := _w.RGB8(150, 200, 250) // Bluish.
	text.Color(textColor)

	return text
}

func createPointBoxWidget(text *_p.TextWidget) *_p.BoxWidget {
	box := _p.Box(text)

	box.Height(pointSize.Height)
	box.Width(pointSize.Width)
	box.CrossAlign(_p.CrossAxisCenter)
	box.Padding(paddingSize)

	borderColor := _w.RGB(0.6, 0, 0) // Slightly dark red.
	borderWidth := float32(2)
	box.BorderStyle(borderWidth, borderColor)

	return box
}
