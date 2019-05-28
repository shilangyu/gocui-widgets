package widgets

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/gocui-widgets/utils"
)

// Text is a widget that displays text
type Text struct {
	// View is a pointer to the gocui view of this widget
	View *gocui.View
	// Type is an enum WidgetType
	Type WidgetType

	name   string
	text   string
	frame  bool
	center bool
	x, y   int
	w, h   int
}

// NewText initializes the Text widget
// if frame is true a border is rendered
// if center is true x and y becomes the center not start
func NewText(name, text string, frame, center bool, x, y int) *Text {
	w, h := utils.StringDimensions(text)
	w++
	h++

	if center {
		x = x - w/2
		y = y - h/2
	}

	return &Text{nil, TypeText, name, text, frame, center, x, y, w, h}
}

// Layout renders the Text widget
func (w *Text) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	w.View = v

	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = w.frame
		fmt.Fprint(v, w.text)
	}

	return nil
}

// ChangeText changes the text within
func (w *Text) ChangeText(s string) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		w.View.Clear()
		fmt.Fprint(w.View, s)

		return nil
	}
}
