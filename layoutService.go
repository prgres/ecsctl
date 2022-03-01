package main

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
	"github.com/prgres/ecsctl/service"
)

func layoutServicesList(g *gocui.Gui, services []*service.ServiceData) error {
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
	// v.Title = cluster.Name

	for _, s := range services {
		fmt.Fprintln(v, s.Name)
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}

func layoutServiceDetail(g *gocui.Gui, service *service.ServiceData) error {
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

func layoutServices(ctx *context.Context, g *gocui.Gui) error {
	cluster := ctx.ActiveCluster

	if !ctx.IsServiceFetched {
		if err := cluster.FetchServices(); err != nil {
			return err
		}

		ctx.IsServiceFetched = true
	}

	v, err := g.View(viewServicesId)
	if err != nil {
		if err == gocui.ErrUnknownView {
			return layoutServicesList(g, cluster.Services)
		}

		return err
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}

func showServiceDetails(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()

	_, cy := v.Cursor()
	l, err := v.Line(cy)
	if err != nil {
		l = ""
	}

	if l == "" {
		return errors.New("l: " + l)
	}

	for _, s := range ctx.ActiveCluster.Services {
		if s.Name == l {
			return layoutServiceDetail(g, s)
		}
	}

	return errors.New("cluster: " + l + " not found")
}
