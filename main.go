package main

import (
	"log"
	"strings"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/cluster"
	"github.com/prgres/ecsctl/gui"
)

var (
	_ctx *gui.Context
)

func initWidgets(ctx *gui.Context, g *gocui.Gui) {
	/* --- */
	viewClusterList := viewClusterList(g)
	ctx.Views = append(ctx.Views, viewClusterList)

	/* --- */
	viewServiceList := viewServiceList(g)
	ctx.Views = append(ctx.Views, viewServiceList)
}

func initClusterData(ctx *gui.Context) error {
	clustersData, err := cluster.GetClusters(_ctx.Ctx, _ctx.AwsCfg)
	if err != nil {
		return err
	}
	_ctx.ClustersData = clustersData

	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		handleErr(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	_ctx, err = gui.NewContext()
	if err != nil {
		handleErr(err)
	}

	if err := initClusterData(_ctx); err != nil {
		handleErr(err)
	}

	g.SetManagerFunc(mainLoop)

	if err := keybindings(g); err != nil {
		handleErr(err)
	}

	initWidgets(_ctx, g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		handleErr(err)
	}
}

func handleErr(err error) {
	switch {
	case strings.Contains(err.Error(), "ExpiredTokenException"):
		log.Fatal("AWS credentials expired. Please re-login.")

	default:
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

func routes(ctx *gui.Context, g *gocui.Gui) error {
	if ctx.CurrentView == nil {
		_, err := ctx.SetCurrentView(gui.ViewClusterListId) // fallback
		return err
	}

	return ctx.CurrentView.Show(ctx, g)
}

func render(ctx *gui.Context, g *gocui.Gui) error {
	return ctx.CurrentView.Render(ctx, g)
}
