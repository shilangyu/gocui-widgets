package main

import (
	"os"

	"github.com/jroimartin/gocui"
	widgets "github.com/shilangyu/gocui-widgets"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()
	g.Mouse = true

	choices := []string{
		"OK",
		"CLOSE",
	}

	w, h := g.Size()
	var modalWi *widgets.Modal
	modalWi = widgets.NewModal("modal", "Hey im a pop-up.\nOk bye", choices, true, w/2, h/2, nil, func(i int) {
		switch i {
		case 0:
			modalWi.Hide()
		case 1:
			g.Close()
			os.Exit(0)
		}
	})

	g.SetManager(modalWi)

	g.Update(func(g *gocui.Gui) error {
		g.SetCurrentView("modal")
		return nil
	})

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
