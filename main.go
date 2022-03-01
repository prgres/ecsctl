package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/context"
)

var (
	viewClusterId       = "clusterView"
	viewServicesId      = "servicesListView"
	viewServiceDetailId = "servicesDetailView"

	_ctx *context.Context
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	_ctx, err = context.New()
	if err != nil {
		log.Panicln(err)
	}

	g.SetManagerFunc(routes)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func routes(g *gocui.Gui) error {
	ctx := _ctx.Context()

	if g.CurrentView() == nil {
		return layoutClusters(ctx, g)
	}

	switch g.CurrentView().Name() {

	case viewClusterId:
		return layoutClusters(ctx, g)

	case viewServicesId:
		return layoutServices(ctx, g)

	default:
		return nil
	}
}
