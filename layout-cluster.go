package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
	"github.com/prgres/ecsctl/widget"
)

func widgetClusterListShow(ctx *context.Context, g *gocui.Gui, widget *widget.Widget) error {
	clustersName := func() []string {
		result := make([]string, len(ctx.ClustersData))
		for i := range ctx.ClustersData {
			result[i] = ctx.ClustersData[i].Name
		}

		return result
	}()

	widget.UpdateData(clustersName)
	v, err := widget.Get(g)
	if err != nil {
		return nil
	}

	_, _ = g.SetCurrentView(v.Name())
	return nil
}

/* --- keybinding func --- */
var prevLineMouseClick = ""

func widgetClusterListClickMouse(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()
	_, cy := v.Cursor()

	clusterId, err := v.Line(cy)
	if err != nil {
		return nil
	}

	if prevLineMouseClick == clusterId {
		if _, err := ctx.SetActiveClusterId(clusterId); err != nil {
			return err
		}

		return viewServiceListShow(ctx, g)
	}

	if prevLineMouseClick == "" {
		prevLineMouseClick = clusterId
	}

	return nil
}

func widgetClusterListClick(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()
	_, cy := v.Cursor()

	clusterId, err := v.Line(cy)
	if err != nil {
		return nil
	}

	if _, err := ctx.SetActiveClusterId(clusterId); err != nil {
		return err
	}

	if err := viewServiceListShow(ctx, g); err != nil {
		return err
	}

	return g.DeleteView(v.Name())
}
