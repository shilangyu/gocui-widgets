package widgets

import (
	"fmt"
	"math"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/shilangyu/gocui-widgets/utils"
)

// Modal is a widget creates a pop-up with buttons
type Modal struct {
	// View is a pointer to the gocui view of this widget
	View *gocui.View
	// Type is an enum WidgetType
	Type WidgetType

	name     string
	text     string
	choices  []string
	x, y     int
	w, h     int
	center   bool
	onChange func(i int)
	onSubmit func(i int)
	currItem int
	killed   bool
}

// NewModal initializes the Menu widget
func NewModal(name string, text string, choices []string, center bool, x, y int, onChange, onSubmit func(i int)) *Modal {
	max1, h := utils.StringDimensions(text)
	max2 := len(strings.Join(choices, " "))
	w := int(math.Max(float64(max1), float64(max2)))
	h += 2
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

	return &Modal{nil, TypeModal, name, text, choices, x, y, w, h, center, onChange, onSubmit, 0, false}
}

// handles keystroke events
func (w *Modal) onArrow(change int) func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		w.currItem += change
		if w.currItem == -1 {
			w.currItem++
		} else if w.currItem == len(w.choices) {
			w.currItem--
		} else {
			w.View.Clear()
			w.onChange(w.currItem)
		}

		return nil
	}
}

// handles mouse event
func (w *Modal) onMouse(g *gocui.Gui, v *gocui.View) error {
	xCursor, yCursor := v.Cursor()
	xText, yText := utils.StringDimensions(w.text)
	if yText+1 != yCursor {
		return nil
	}

	cum := xCursor - int(math.Max(float64(xText-len(strings.Join(w.choices, " "))), 0))
	for i, choice := range w.choices {
		currLength := len(choice)
		if cum < 0 {
			return nil
		} else if cum < currLength {
			if i != w.currItem {
				w.currItem = i
				w.View.Clear()
				w.onChange(w.currItem)
			} else {
				w.onSubmit(i)
			}
			break
		} else {
			cum -= (currLength + 1)
		}
	}
	return nil
}

// Layout renders the Modal widget
func (w *Modal) Layout(g *gocui.Gui) error {
	if w.killed {
		g.DeleteKeybindings("modal")
		g.DeleteView("modal")
		w.View.Frame = false
		w.View.Clear()
		return nil
	}

	v, err := g.SetView(w.name, w.x, w.y, w.x+w.w, w.y+w.h)
	w.View = v

	if err == gocui.ErrUnknownView {
		if err := g.SetKeybinding(w.name, gocui.MouseLeft, gocui.ModNone, w.onMouse); err != nil {
			return err
		}

		if err := g.SetKeybinding(w.name, gocui.KeyArrowLeft, gocui.ModNone, w.onArrow(-1)); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyArrowRight, gocui.ModNone, w.onArrow(1)); err != nil {
			return err
		}
		if err := g.SetKeybinding(w.name, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			w.onSubmit(w.currItem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	v.SelFgColor = gocui.ColorBlack

	fmt.Fprintf(v, "%s\n\n", w.text)
	var joinedChoices string
	for i, choice := range w.choices {
		if i == w.currItem {
			joinedChoices += "\u001b[42m" + choice + "\u001b[0m "
		} else {
			joinedChoices += choice + " "
		}
	}
	fmt.Fprintf(v, "%s%s", strings.Repeat(" ", w.w-len(strings.Join(w.choices, " "))-1), joinedChoices)

	return nil
}

// Hide "kills" the modal
func (w *Modal) Hide() {
	w.killed = true
}
