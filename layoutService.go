package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type ServiceData struct {
	Name string
	Arn  string
}

func layoutServicesInit(g *gocui.Gui, cluster *ClusterData) error {
	services, err := cluster.GetServices()
	if err != nil {
		return nil
	}

	maxX, maxY := g.Size()
	v, err := g.SetView(viewServicesId, 1, 1, maxX-1, maxY/2)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	v.Highlight = true
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack
	v.Autoscroll = true
	v.Frame = true
	v.Title = cluster.Name

	for _, s := range services {
		fmt.Fprintln(v, s.Name)
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}

func layoutServices(g *gocui.Gui, cluster *ClusterData) error {
	v, err := g.View(viewServicesId)
	if err != nil {
		if err == gocui.ErrUnknownView {
			return layoutServicesInit(g, cluster)
		}

		return err
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}
