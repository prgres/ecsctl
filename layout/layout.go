package layout

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Layout struct {
	ViewId string

	X1, X2 int
	Y1, Y2 int
}

func (l *Layout) Get(g *gocui.Gui) (*gocui.View, error) {
	v, err := g.View(l.ViewId)
	if err != nil {
		if err == gocui.ErrUnknownView {
			return l.init(g)
		}

		return nil, err
	}

	return v, nil
}

func (l *Layout) init(g *gocui.Gui) (*gocui.View, error) {
	v, err := g.SetView(l.ViewId, l.X1, l.X2, l.Y1, l.Y2)
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

func (l *Layout) Render(g *gocui.Gui, data []string) (*gocui.View, error) {
	v, err := l.Get(g)
	if err != nil {
		return nil, err
	}

	v.Clear()

	for i := range data {
		fmt.Fprintln(v, data[i])
	}

	return v, nil
}
