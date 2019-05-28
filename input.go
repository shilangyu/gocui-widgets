package widgets

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

// Input is a widget allows for user input
type Input struct {
	// View is a pointer to the gocui view of this widget
	View *gocui.View

	name     string
	text     string
	frame    bool
	center   bool
	x, y     int
	w, h     int
	onChange gocui.EditorFunc
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

	return &Input{nil, name, "", frame, center, x, y, w, h, onChange}
}

// Layout renders the Input widget
func (w *Input) Layout(g *gocui.Gui) error {
	g.Cursor = true

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	w.View = v

	if err == gocui.ErrUnknownView {
		v.Editor = gocui.EditorFunc(func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
			gocui.DefaultEditor.Edit(v, key, ch, mod)
			w.onChange(v, key, ch, mod)
		})
		g.SetCurrentView(w.name)
	} else if err != nil {
		return nil
	}

	v.Editable = true
	v.Wrap = true

	v.Frame = w.frame

	return nil
}

// ChangeText changes the inner text
func (w *Input) ChangeText(s string) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		w.View.Clear()
		fmt.Fprint(w.View, s)
		w.View.SetCursor(len(s), len(strings.Split(s, "\n"))-1)

		return nil
	}
}
