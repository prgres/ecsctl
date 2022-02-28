package main

import (
	"errors"
	"fmt"

	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/cmd"
)

type ClusterData struct {
	Name     string
	Services []*ServiceData
}

func (c *ClusterData) GetServices() ([]*ServiceData, error) {
	if len(c.Services) == 0 {
		services, err := cmd.GetServices(c.Name)
		if err != nil {
			return nil, err
		}

		for _, s := range services {
			c.Services = append(c.Services, &ServiceData{
				Name: cmd.GetNameFromResourceId(s),
				Arn:  s,
			})
		}
	}

	return c.Services, nil
}

func GetClusters(clustersData []*ClusterData) ([]*ClusterData, error) {
	//TODO: background refresh
	if len(clustersData) == 0 {
		clusters, err := cmd.GetClustersNames()
		if err != nil {
			return nil, err
		}

		for _, c := range clusters {
			clustersData = append(clustersData, &ClusterData{
				Name: c,
			})
		}
	}

	return clustersData, nil
}

func layoutClustersInit(g *gocui.Gui) error {
	clusters, err := GetClusters(_clustersData)
	if err != nil {
		return err
	}

	_clustersData = clusters

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

	for _, c := range clusters {
		fmt.Fprintln(v, c.Name)
	}

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	return nil
}

func layoutClusters(g *gocui.Gui) error {
	v, err := g.View(viewClusterId)
	if err != nil {
		if err == gocui.ErrUnknownView {
			return layoutClustersInit(g)
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
	l, err := v.Line(cy)
	if err != nil {
		l = ""
	}

	if prevLineMouseClick == l {
		return showServicesLayout(g, v)
	}

	if prevLineMouseClick == "" {
		prevLineMouseClick = l
	}

	return nil
}

func showServicesLayout(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	l, err := v.Line(cy)
	if err != nil {
		l = ""
	}

	if l == "" {
		return nil
	}

	for _, c := range _clustersData {
		if l == c.Name {
			g.DeleteView(v.Name())
			return layoutServices(g, c)
		}
	}

	return errors.New("cluster: " + l + " not found")
}
