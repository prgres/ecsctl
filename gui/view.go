package gui

import (
	"errors"

	"github.com/jroimartin/gocui"
)

type ShowFuncType func(ctx *Context, g *gocui.Gui) error

type View struct {
	Id      string
	Widgets []*Widget
	Show    ShowFuncType
}

func NewView(id string, showFunc ShowFuncType, widgets ...*Widget) *View {
	return &View{
		Id:      id,
		Widgets: widgets,
		Show:    showFunc,
	}
}

func (v *View) Render(ctx *Context, g *gocui.Gui) error {
	for i := range v.Widgets {
		if err := v.Widgets[i].Update(ctx, g); err != nil {
			return err
		}

		if err := v.Widgets[i].Render(g); err != nil {
			return err
		}
	}

	return nil
}

func (v *View) Widget(id string) (*Widget, error) {
	for i := range v.Widgets {
		if id == v.Widgets[i].Id {
			return v.Widgets[i], nil
		}
	}

	return nil, errors.New("widget: " + id + " not found")
}
