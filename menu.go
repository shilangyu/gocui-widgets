package widgets

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/gocui-widgets/utils"
)

// Menu is a widget that creates a vertical clickable menu
type Menu struct {
	// View is a pointer to the gocui view of this widget
	View *gocui.View
	// Type is an enum WidgetType
	Type WidgetType

	name     string
	items    []string
	x, y     int
	w, h     int
	center   bool
	onChange func(i int)
	onSubmit func(i int)
	currItem int
}

// NewMenu initializes the Menu widget
func NewMenu(name string, items []string, center bool, x, y int, onChange, onSubmit func(i int)) *Menu {
	w, h := utils.StringDimensions(strings.Join(items, "\n"))
	w++
	h++

	if center {
		x = x - w/2
		y = y - h/2
	}

	if onChange == nil {
		onChange = func(i int) {}
	}
	if onSubmit == nil {
		onSubmit = func(i int) {}
	}

	return &Menu{nil, TypeMenu, name, items, x, y, w, h, center, onChange, onSubmit, 0}
}

// handles keystroke events
func (w *Menu) onArrow(change int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		w.currItem += change
		if w.currItem == -1 {
			w.currItem++
		} else if w.currItem == len(w.items) {
			w.currItem--
		} else {
			v, err := g.View(w.name)
			if err != nil {
				return err
			}
			v.SetCursor(0, w.currItem)

			w.onChange(w.currItem)
		}

		return nil
	}
}

// handles mouse event
func (w *Menu) onMouse(g *gocui.Gui, v *gocui.View) error {
	_, currItem := v.Cursor()
	if currItem != w.currItem {
		w.currItem = currItem
		w.onChange(w.currItem)
	} else {
		w.onSubmit(currItem)
	}
	return nil
}

// Layout renders the Menu widget
func (w *Menu) Layout(g *gocui.Gui) error {
	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	w.View = v

	if err == gocui.ErrUnknownView {
		if err := g.SetKeybinding(w.name, gocui.MouseLeft, gocui.ModNone, w.onMouse); err != nil {
			return err
		}

		if err := g.SetKeybinding(w.name, gocui.KeyArrowDown, gocui.ModNone, w.onArrow(1)); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyArrowUp, gocui.ModNone, w.onArrow(-1)); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			w.onSubmit(w.currItem)
			return nil
		}); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	v.Highlight = true
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack

	for _, text := range w.items {
		fmt.Fprintln(v, text)
	}

	return nil
}

// ChangeSelected changes the selected menu item
func (w *Menu) ChangeSelected(i int) func(g *gocui.Gui) error {
	return func(g *gocui.Gui) error {
		w.View.SetCursor(0, i)
		w.currItem = i

		return nil
	}
}
