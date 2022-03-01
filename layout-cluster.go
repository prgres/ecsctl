package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
)

func layoutClusterListShow(ctx *context.Context, g *gocui.Gui) error {
	clustersName := func() []string {
		result := make([]string, len(ctx.ClustersData))
		for i := range ctx.ClustersData {
			result[i] = ctx.ClustersData[i].Name
		}

		return result
	}()

	v, err := layoutClusterList.Render(g, clustersName)
	if err != nil {
		return nil
	}

	_, _ = g.SetCurrentView(v.Name())

	return nil
}

/* --- keybinding func --- */
var prevLineMouseClick = ""

func layoutClustersClickMouse(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()
	_, cy := v.Cursor()

	clusterId, err := v.Line(cy)
	if err != nil {
		clusterId = ""
	}

	if prevLineMouseClick == clusterId {
		return layoutServiceListShow(ctx, g)
	}

	if prevLineMouseClick == "" {
		prevLineMouseClick = clusterId
	}

	return nil
}

func layoutClustersClick(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()
	_, cy := v.Cursor()

	clusterId, err := v.Line(cy)
	if err != nil {
		return nil
	}

	if _, err := ctx.SetActiveClusterId(clusterId); err != nil {
		return err
	}

	if err := layoutServiceListShow(ctx, g); err != nil {
		return err
	}

	return g.DeleteView(v.Name())
}
