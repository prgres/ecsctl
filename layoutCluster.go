package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/cluster"
	"github.com/prgres/ecsctl/context"
)

func layoutClustersInit(g *gocui.Gui, clustersData []*cluster.ClusterData) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(viewClusterId, maxX/4, maxY/4, 3*maxX/4, 3*maxY/4)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	v.Highlight = true
	v.Autoscroll = true
	v.Frame = true
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack

	for _, c := range clustersData {
		fmt.Fprintln(v, c.Name)
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}

func layoutClusters(ctx *context.Context, g *gocui.Gui) error {
	v, err := g.View(viewClusterId)
	if err != nil {
		if err == gocui.ErrUnknownView {
			return layoutClustersInit(g, ctx.ClustersData)
		}

		return err
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}

var prevLineMouseClick = ""

func showServicesLayoutMouse(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	clusterId, err := v.Line(cy)
	if err != nil {
		clusterId = ""
	}

	if prevLineMouseClick == clusterId {
		return showServicesLayout(g, v)
	}

	if prevLineMouseClick == "" {
		prevLineMouseClick = clusterId
	}

	return nil
}

func showServicesLayout(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()

	_, cy := v.Cursor()
	clusterId, err := v.Line(cy)
	if err != nil {
		return nil // most likely 'clusterId == "" '
	}

	ctx.SetActiveClusterId(clusterId)

	if err := layoutServices(ctx, g); err != nil {
		return err
	}

	return g.DeleteView(v.Name())
}
