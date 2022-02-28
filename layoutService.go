package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/jroimartin/gocui"
)

type ServiceData struct {
	Name string
	Arn  string

	*types.Service
}

func layoutServicesList(g *gocui.Gui, cluster *ClusterData) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(viewServicesId, 1, 1, maxX/3-1, maxY-1)
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

	for _, s := range cluster.Services {
		fmt.Fprintln(v, s.Name)
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}

func layoutServiceDetail(g *gocui.Gui, service *ServiceData) error {
	maxX, maxY := g.Size()

	oldV, err := g.View(viewServiceDetailId)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	if oldV != nil {
		if err := g.DeleteView(viewServiceDetailId); err != nil {
			return err
		}
	}

	v, err := g.SetView(viewServiceDetailId, 1*maxX/3, 1, maxX-1, maxY-1)
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

	if service.Service != nil {
		fmt.Fprintln(v, "Status: ", *service.Status)
		fmt.Fprintln(v, "Arn: ", *service.ServiceArn)
	}

	return nil
}

var clusterP *ClusterData

func layoutServicesInit(g *gocui.Gui, cluster *ClusterData, service *ServiceData) error {
	clusterP = cluster
	if err := layoutServicesList(g, cluster); err != nil {
		return err
	}

	return nil
}

func layoutServices(g *gocui.Gui, cluster *ClusterData) error {
	if err := cluster.FetchServices(); err != nil {
		return err
	}

	v, err := g.View(viewServicesId)
	if err != nil {
		if err == gocui.ErrUnknownView {
			return layoutServicesInit(g, cluster, cluster.Services[0])
		}

		return err
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}
