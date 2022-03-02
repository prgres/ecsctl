package main

import (
	"errors"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
	"github.com/prgres/ecsctl/widget"
)

func widgetServiceListShow(ctx *context.Context, g *gocui.Gui, widget *widget.Widget) error {
	cluster := ctx.ActiveCluster

	if err := cluster.FetchServices(); err != nil {
		return err
	}

	servicesName := func() []string {
		result := make([]string, len(cluster.Services))
		for i := range cluster.Services {
			result[i] = cluster.Services[i].Name
		}

		return result
	}()

	widget.UpdateData(servicesName)
	v, err := widget.Get(g)
	if err != nil {
		return err
	}

	_, err = g.SetCurrentView(v.Name())

	return err
}

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

	widget, err := ctx.CurrentView.Widget(widgetServiceDetailId)
	if err != nil {
		return err
	}

	widget.UpdateData(service.Render())

	return nil
}
