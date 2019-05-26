package main

import (
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/gocui-widgets"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	w, h := g.Size()
	collectionWi := widgets.NewCollection("collection", "grouper", true, w/2, h/2, w/5, h/5)

	g.SetManager(collectionWi)

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
