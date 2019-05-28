/*
Package widgets gives you widgets that allow you to quickly start writing TUI apps with gocui.
Creating a widget is easy:
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
		textWi := widgets.NewText("example-widget", "hello world", true, true, w/2, h/2)

		g.SetManager(text)

		if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
			panic(err)
		}
	}
To see more please visit '_examples/'
*/
package widgets

// WidgetType is a enum describing the widget
type WidgetType int

const (
	// TypeText for Text widget
	TypeText WidgetType = iota
	// TypeMenu for Menu widget
	TypeMenu
	// TypeCollection for Collection widget
	TypeCollection
	// TypeInput for Input widget
	TypeInput
)
