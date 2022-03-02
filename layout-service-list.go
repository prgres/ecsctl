package main

import (
	"errors"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
)

func widgetServiceListShow(ctx *context.Context, g *gocui.Gui) error {
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

	widgetServiceList.UpdateData(servicesName)
	v, err := widgetServiceList.Get(g)
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

	for _, s := range ctx.ActiveCluster.Services {
		if s.Name == serviceId {
			widgetServiceDetail.UpdateData(s.Render())
			return nil
		}
	}

	return errors.New("service: " + serviceId + " not found")
}
