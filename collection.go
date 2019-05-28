package widgets

import (
	"github.com/jroimartin/gocui"
)

// Collection is a widget that groups other widgets
type Collection struct {
	// View is a pointer to the gocui view of this widget
	View *gocui.View

	name   string
	title  string
	center bool
	x, y   int
	w, h   int
}

// NewCollection initializes the Collection widget
// if center is true x and y becomes the center not start
func NewCollection(name, title string, center bool, x, y int, w, h int) *Collection {
	w--
	h--

	if center {
		x = x - w/2
		y = y - h/2
	}

	return &Collection{nil, name, title, center, x, y, w, h}
}

// Layout renders the Collection widget
func (w *Collection) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	w.View = v
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = w.title
	}

	return nil
}

// ChangeTitle changes the title of the widget
func (w *Collection) ChangeTitle(t string) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		w.View.Title = t
		return nil
	}
}
