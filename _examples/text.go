package main

import (
	"github.com/jroimartin/gocui"
	"github.com/shilangyu/gocui-widgets"
	"time"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	wiTexts := []string{
		"Hello world!",
		"!dlrow olleH",
	}
	currText := 0

	w, h := g.Size()
	textWi := widgets.NewText("text", wiTexts[currText], true, true, w/2, h/2)

	g.SetManager(textWi)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			currText++
			g.Update(textWi.ChangeText(wiTexts[currText%2]))
		}
	}()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
