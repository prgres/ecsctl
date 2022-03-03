package main

import (
	"errors"

	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/gui"
)

/* --- keybinding func --- */
func widgetServiceListClick(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()
	_, cy := v.Cursor()

	serviceId, err := v.Line(cy)
	if err != nil {
		return errors.New("service: " + serviceId) //TODO
	}

	service, err := ctx.ActiveCluster.Service(serviceId)
	if err != nil {
		return err
	}

	widget, err := ctx.CurrentView.Widget(gui.WidgetServiceDetailId)
	if err != nil {
		return err
	}

	widget.UpdateData(service.Render())

	return nil
}
