package widgets

import (
	"github.com/jroimartin/gocui"
)

// Input is a widget allows for user input
type Input struct {
	name     string
	Text     string
	Frame    bool
	Center   bool
	x, y     int
	w, h     int
	OnChange gocui.EditorFunc
}

// NewInput initializes the Input widget
// if frame is true a border is rendered
// if center is true x and y becomes the center not start
func NewInput(name string, frame, center bool, x, y int, w, h int, onChange gocui.EditorFunc) *Input {
	w--
	h--

	if center {
		x = x - w/2
		y = y - h/2
	}

	if onChange == nil {
		onChange = func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {}
	}

	return &Input{name, "", frame, center, x, y, w, h, onChange}
}

// Name returns the widget name
func (w *Input) Name() string {
	return w.name
}

// Coord returns the x and y of the widget
func (w *Input) Coord() (int, int) {
	return w.x, w.y
}

// Size returns the width and height of the widget
func (w *Input) Size() (int, int) {
	return w.w, w.h
}

// Layout renders the Input widget
func (w *Input) Layout(g *gocui.Gui) error {
	g.Cursor = true

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	if err == gocui.ErrUnknownView {
		v.Editor = gocui.EditorFunc(func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
			gocui.DefaultEditor.Edit(v, key, ch, mod)
			w.OnChange(v, key, ch, mod)
		})
		g.SetCurrentView(w.name)
	} else if err != nil {
		return nil
	}

	v.Editable = true
	v.Wrap = true

	v.Frame = w.Frame

	return nil
}
