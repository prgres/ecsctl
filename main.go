package main

import (
	"log"

	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/context"
	"github.com/prgres/ecsctl/layout"
)

const (
	viewClusterListId   = "clusterListView"
	viewServiceListId   = "serviceListView"
	viewServiceDetailId = "serviceDetailView"
)

var (
	layoutClusterList   *layout.Layout
	layoutServiceList   *layout.Layout
	layoutServiceDetail *layout.Layout

	_ctx *context.Context
)

func initLayouts(g *gocui.Gui) {
	maxX, maxY := g.Size()

	layoutClusterList = &layout.Layout{
		ViewId: viewClusterListId, X1: maxX / 4, X2: maxY / 4, Y1: 3 * maxX / 4, Y2: 3 * maxY / 4,
	}

	layoutServiceList = &layout.Layout{
		ViewId: viewServiceListId, X1: 1, X2: 1, Y1: maxX/3 - 1, Y2: maxY - 1,
	}

	layoutServiceDetail = &layout.Layout{
		ViewId: viewServiceDetailId, X1: 1 * maxX / 3, X2: 1, Y1: maxX - 1, Y2: maxY - 1,
	}
}

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

	initLayouts(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func routes(g *gocui.Gui) error {
	ctx := _ctx.Context()

	if g.CurrentView() == nil {
		return layoutClusterListShow(ctx, g)
	}

	switch g.CurrentView().Name() {
	case viewClusterListId:
		return layoutClusterListShow(ctx, g)

	case viewServiceListId:
		return layoutServiceListShow(ctx, g)

	default:
		return nil
	}
}
