package view

import (
	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/widget"
)

type View struct {
	Id      string
	Widgets []*widget.Widget
}

func New(id string, widgets ...*widget.Widget) *View {
	return &View{
		Id:      id,
		Widgets: widgets,
	}
}

func (v *View) Render(g *gocui.Gui) error {
	for i := range v.Widgets {
		if err := v.Widgets[i].Render(g); err != nil {
			return err
		}
	}

	return nil
}
