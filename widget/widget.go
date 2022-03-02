package widget

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Widget struct {
	Id string

	X1, X2 int
	Y1, Y2 int

	Data []string
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

func (w *Widget) init(g *gocui.Gui) (*gocui.View, error) {
	v, err := g.SetView(w.Id, w.X1, w.X2, w.Y1, w.Y2)
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

func (w *Widget) Render(g *gocui.Gui) error {
	// func (w *Widget) Render(g *gocui.Gui) (*gocui.View, error) {
	v, err := w.Get(g)
	if err != nil {
		return err
		// return nil, err
	}

	v.Clear()

	for i := range w.Data {
		fmt.Fprintln(v, w.Data[i])
	}

	// return v, nil
	return nil
}

func (w *Widget) UpdateData(data []string) {
	w.Data = data
}
