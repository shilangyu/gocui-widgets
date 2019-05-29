package main

import (
	"math/rand"
	"strings"

	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	var inputWi *widgets.Input

	w, h := g.Size()
	textWi := widgets.NewText("text", strings.Repeat(" ", w/4), false, true, 3*w/4, h/2)
	inputWi = widgets.NewInput("input", true, true, w/4, h/2, w/4, 3, func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) bool {
		if key == gocui.KeyEnter {
			return false
		}

		s := ""
		for _, char := range v.Buffer() {
			if rand.Float64() < .5 {
				s += strings.ToUpper(string(char))
			} else {
				s += strings.ToLower(string(char))
			}
		}
		g.Update(textWi.ChangeText(s))

		if len(s) > 10 {
			g.Update(inputWi.ChangeText(""))
		}

		return true
	})

	g.SetManager(textWi, inputWi)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
