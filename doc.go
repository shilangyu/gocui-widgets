/*
Package widgets gives you widgets that allow you to quickly start writing TUI apps with gocui.
Creating a widget is easy:
	package main

	import (
		"github.com/jroimartin/gocui"
		widgets "github.com/shilangyu/gocui-widgets"
	)

	func main() {
		g, err := gocui.NewGui(gocui.OutputNormal)
		if err != nil {
			panic(err)
		}
		defer g.Close()

		w, h := g.Size()
		text := widgets.NewText("example-widget", "hello world", true, true, w/2, h/2)

		g.SetManager(text)

		if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
			panic(err)
		}
	}
*/
package widgets
