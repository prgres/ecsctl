package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

var (
	viewClusterId  = "clusterView"
	viewServicesId = "servicesView"

	_clustersData = []*ClusterData{}
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(managerFunc)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func managerFunc(g *gocui.Gui) error {
	if len(g.Views()) == 0 {
		return layoutClusters(g)
	}

	return nil
}
