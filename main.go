package main

import (
	"log"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
	"github.com/prgres/ecsctl/view"
	"github.com/prgres/ecsctl/widget"
)

const (
	viewClusterListId = "viewClusterList"
	viewServiceListId = "viewServiceList"

	widgetClusterListId   = "widgetClusterList"
	widgetServiceListId   = "widgetServiceList"
	widgetServiceDetailId = "widgetServiceDetail"
)

var (
	_ctx  *context.Context
	views []*view.View

	viewClusterList   *view.View
	widgetClusterList *widget.Widget

	viewServiceList     *view.View
	widgetServiceList   *widget.Widget
	widgetServiceDetail *widget.Widget
)

func initWidgets(g *gocui.Gui) {
	maxX, maxY := g.Size()

	//
	widgetClusterList = &widget.Widget{
		Id: widgetClusterListId, X1: maxX / 4, X2: maxY / 4, Y1: 3 * maxX / 4, Y2: 3 * maxY / 4,
	}
	viewClusterList = view.New(viewClusterListId, widgetClusterList)
	views = append(views, viewClusterList)

	///
	widgetServiceList = &widget.Widget{
		Id: widgetServiceListId, X1: 1, X2: 1, Y1: maxX/3 - 1, Y2: maxY - 1,
	}

	widgetServiceDetail = &widget.Widget{
		Id: widgetServiceDetailId, X1: 1 * maxX / 3, X2: 1, Y1: maxX - 1, Y2: maxY - 1,
	}

	viewServiceList = view.New(viewServiceListId, widgetServiceDetail, widgetServiceList)
	views = append(views, viewServiceList)
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

	g.SetManagerFunc(mainLoop)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	initWidgets(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func mainLoop(g *gocui.Gui) error {
	ctx := _ctx.Context()

	if err := routes(ctx, g); err != nil {
		return err
	}

	if err := render(ctx, g); err != nil {
		return err
	}

	return nil
}

func routes(ctx *context.Context, g *gocui.Gui) error {
	if ctx.CurrentView == nil {
		return viewClusterListShow(ctx, g)
	}

	// switch g.CurrentView().Name() {
	switch ctx.CurrentView.Id {
	case viewClusterListId:
		return viewClusterListShow(ctx, g)

	case viewServiceListId:
		return viewServiceListShow(ctx, g)

	default:
		return nil
	}
}

func render(ctx *context.Context, g *gocui.Gui) error {
	return ctx.CurrentView.Render(g)
}
