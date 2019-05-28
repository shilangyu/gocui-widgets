# gocui-widgets

[![](https://img.shields.io/badge/godoc-reference-5272B4.svg)](http://godoc.org/github.com/shilangyu/gocui-widgets)
[![](https://goreportcard.com/badge/github.com/shilangyu/gocui-widgets)](https://goreportcard.com/report/github.com/shilangyu/gocui-widgets)

Set of gocui widgets to help you build TUI apps. It is meant to use **with** [gocui](https://github.com/jroimartin/gocui) not instead.

```sh
go get github.com/shilangyu/gocui-widgets
```

- [usage](#usage)
- [widgets](#widgets)
  - [Text](#Text)
  - [Collection](#Collection)
  - [Menu](#Menu)
  - [Input](#Input)
- [examples](#examples)

## Usage

Each widgets implements:

```go
type Widget interface {
	gocui.Manager
	Name() string
	Coord() (int, int)
	Size() (int, int)
}
```

Some widgets accept callback functions called `OnSomething` that are called when `Something` happens

Some widgets have modifier methods called `ChangeSomething` that you call with `gocui.Update` whenever you need to change something. [Text](#Text) example:

```go
text := NewText("text", "bad text", false, false, 0, 0)

g.Update(text.ChangeText("good text"))

```

Check the [\_examples](https://github.com/shilangyu/gocui-widgets/tree/master/_examples) to find out more

## Widgets

Common parameters:

- `name`: name needed for gocui (ID)
- `frame`: if true renders a frame around the widget
- `center`: if true `x` and `y` is the center of the widget
- `x`, `y`: x and y position of the widget
- `w`, `h`: width and height of the widget

### Text

Renders text in a given position

```go
NewText(name, text string, frame, center bool, x, y int)
```

- `ChangeText`: changes the inner text of the widget

### Collection

Renders a grouping frame with title

```go
NewCollection(name, title string, center bool, x, y int, w, h int)
```

### Input

Renders an input

```go
NewInput(name string, frame, center bool, x, y int, w, h int, onChange gocui.EditorFunc)
```

- `onChange` is being called with the `gocui.EditorFunc` parameters when text is inputted (you dont have to call `gocui.DefaultEditor.Edit`, it calls it for you)
- `ChangeText`: changes the inner text of the widget

### Menu

Renders a menu

```go
NewMenu(name string, items []string, x, y int, center, arrows bool, onChange, onSubmit func(i int))
```

Menu is controlled by mouse clicks (therefore make sure to set `g.Mouse` to true), if an item is already highlighted and clicked it is counted as a submit action. If arrows is true, the menu will also be controlled by keyboard arrows where <kbd>Enter</kbd> is submit action.

- `onChange` is being called with the index when user changes his highlighted item
- `onSubmit` is being called with the index when user submits a choice

## Examples

To run a widget example clone this repo and go to the `_examples` directory. Then:

```sh
go run <widget-name>.go
```
