package gui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type WidgetUpdateFuncType func(ctx *Context, g *gocui.Gui, widget *Widget) error

type Widget struct {
	Id string

	X1, X2 int
	Y1, Y2 int

	View       *gocui.View //TODO:
	Data       []string
	updateFunc WidgetUpdateFuncType
}

func NewWidget(id string, updateFunc WidgetUpdateFuncType, X1, X2, Y1, Y2 int) *Widget {
	return &Widget{
		Id: id,
		X1: X1, X2: X2,
		Y1: Y1, Y2: Y2,
		updateFunc: updateFunc,
	}
}

func (w *Widget) Update(ctx *Context, g *gocui.Gui) error {
	return w.updateFunc(ctx, g, w)
}

func (w *Widget) Get(g *gocui.Gui) (*gocui.View, error) {
	v, err := g.View(w.Id)
	if err != nil {
		if err == gocui.ErrUnknownView {
			return w.init(g)
		}

		return nil, err
	}

	return v, nil
}

func (w *Widget) Render(g *gocui.Gui) error {
	v, err := w.Get(g)
	if err != nil {
		return err
	}

	v.Clear()

	for i := range w.Data {
		fmt.Fprintln(v, w.Data[i])
	}

	return nil
}

func (w *Widget) UpdateData(data []string) {
	w.Data = data
}

func (w *Widget) init(g *gocui.Gui) (*gocui.View, error) {
	v, err := g.SetView(w.Id, w.X1, w.Y1, w.X2, w.Y2)
	if err != nil && err != gocui.ErrUnknownView {
		return nil, err
	}

	v.Highlight = true
	v.Autoscroll = true
	v.Frame = true
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack

	return v, nil
}
