package main

import (
	"strconv"

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

	menuItems := []string{
		"Click me!",
		"No! Click me!",
		"Please click me!",
	}

	w, h := g.Size()
	textWi1 := widgets.NewText("text1", "i see you selected #1", true, true, 3*w/4, h/2)
	textWi2 := widgets.NewText("text2", "                    ", true, true, w/4, h/2)
	menuWi := widgets.NewMenu("menu", menuItems, w/2, h/2, true, true, func(i int) {
		g.Update(textWi1.ChangeText("i see you selected #" + strconv.Itoa(i+1)))
	}, func(i int) {
		g.Update(textWi2.ChangeText("i see you clicked #" + strconv.Itoa(i+1)))
	})

	g.SetManager(textWi1, textWi2, menuWi)

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
