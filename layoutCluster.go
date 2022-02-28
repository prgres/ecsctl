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

func (c *ClusterData) FetchServices() error {
	if err := c.FetchServicesNameArn(); err != nil {
		return err
	}

	if err := c.FetchServicesDetail(); err != nil {
		return err
	}

	return nil
}

func (c *ClusterData) FetchServicesNameArn() error {
	services, err := cmd.GetServices(c.Name)
	if err != nil {
		return err
	}

	for _, s := range services {
		c.Services = append(c.Services, &ServiceData{
			Name: cmd.GetNameFromResourceId(s),
			Arn:  s,
		})
	}

	return nil
}

func (c *ClusterData) FetchServicesDetail() error {
	services, err := cmd.DescribeServices(c.Name, c.ServicesArn())
	if err != nil {
		return err
	}

	// file, err := os.Create("test.log")
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// json1, err := json.MarshalIndent(&services, "", "  ")
	// if err != nil {
	// 	return err
	// }

	// json2, _ := json.MarshalIndent(&c.Services, "", "  ")
	// if err != nil {
	// 	return err
	// }

	// bw := bufio.NewWriter(file)
	// if _, err := bw.WriteString(string(json1)); err != nil {
	// 	return err
	// }
	// if _, err := bw.WriteString(string(json2)); err != nil {
	// 	return err
	// }

	for _, s := range services {
		for j, is := range c.Services {
			// if *s.ServiceName == "eden-kafka-schema-registry-dev-ServiceDefinitionInternalELB-1QY0GMMC0OX3Q" {
			// 	fmt.Println("")
			// }
			if *s.ServiceArn == is.Arn {
				c.Services[j].Service = s
			}
		}
	}

	return nil
}

func (c *ClusterData) ServicesArn() []string {
	result := make([]string, len(c.Services))
	for i, s := range c.Services {
		result[i] = s.Arn
	}

	return result
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

func showServiceDetails(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	l, err := v.Line(cy)
	if err != nil {
		l = ""
	}

	if l == "" {
		return errors.New("l: " + l)
	}

	for _, s := range clusterP.Services {
		if s.Name == l {
			return layoutServiceDetail(g, s)
		}
	}

	return errors.New("cluster: " + l + " not found")
}
