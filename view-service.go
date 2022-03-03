package main

import (
	"errors"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/gui"
)

func viewServiceList(g *gocui.Gui) *gui.View {
	maxX, maxY := g.Size()

	widgetServiceList := gui.NewWidget(gui.WidgetServiceListId, func(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
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
	},
		1, maxX/3-1, 1, maxY-1,
	)

	widgetServiceDetail := gui.NewWidget(gui.WidgetServiceDetailId,
		func(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
			service := ctx.ActiveService
			if service == nil {
				return nil
			}

			widget.UpdateData(service.Render())

			return nil
		}, 1*maxX/3, maxX-1, 1, maxY-1,
	)

	viewServiceList := gui.NewView(gui.ViewServiceListId, func(ctx *gui.Context, g *gocui.Gui) error {
		view, err := ctx.SetCurrentView(gui.ViewServiceListId)
		if err != nil {
			return err
		}

		w, err := view.Widget(gui.WidgetServiceListId)
		if err != nil {
			return err
		}

		if err := w.Update(ctx, g); err != nil {
			return err
		}

		w, err = view.Widget(gui.WidgetServiceDetailId)
		if err != nil {
			return err
		}

		if err := w.Update(ctx, g); err != nil {
			return err
		}

		return nil
	}, widgetServiceDetail, widgetServiceList)

	return viewServiceList
}

func viewServiceListClear(ctx *gui.Context, g *gocui.Gui) error {
	// omit error because it can only return ErrUnknownView which does not bother us at this moment
	_ = g.DeleteView(gui.WidgetServiceListId)
	_ = g.DeleteView(gui.WidgetServiceDetailId)

	return nil
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

	widget, err := ctx.CurrentView.Widget(gui.WidgetServiceDetailId)
	if err != nil {
		return err
	}

	widget.UpdateData(service.Render())

	return nil
}
