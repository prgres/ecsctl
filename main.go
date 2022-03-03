package main

import (
	"log"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/gui"
)

var (
	_ctx *gui.Context
)

func initWidgets(ctx *gui.Context, g *gocui.Gui) {
	//
	viewClusterList := viewClusterList(g)
	ctx.Views = append(ctx.Views, viewClusterList)

	//
	viewServiceList := viewServiceList(g)
	ctx.Views = append(ctx.Views, viewServiceList)
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	_ctx, err = gui.NewContext()
	if err != nil {
		log.Panicln(err)
	}

	g.SetManagerFunc(mainLoop)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}

	initWidgets(_ctx, g)

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

func routes(ctx *gui.Context, g *gocui.Gui) error {
	if ctx.CurrentView == nil {
		_, err := ctx.SetCurrentView(gui.ViewClusterListId) // fallback
		return err
	}

	return ctx.CurrentView.Show(ctx, g)
}

func render(ctx *gui.Context, g *gocui.Gui) error {
	return ctx.CurrentView.Render(g)
}
