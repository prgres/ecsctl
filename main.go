package main

import (
	"log"

	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/gui"
)

const (
	viewClusterListId = "viewClusterList"
	viewServiceListId = "viewServiceList"

	widgetClusterListId   = "widgetClusterList"
	widgetServiceListId   = "widgetServiceList"
	widgetServiceDetailId = "widgetServiceDetail"
)

var (
	_ctx *gui.Context
)

func initWidgets(ctx *gui.Context, g *gocui.Gui) {
	maxX, maxY := g.Size()

	//
	widgetClusterList := gui.NewWidget(widgetClusterListId, func(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
		clustersName := func() []string {
			result := make([]string, len(ctx.ClustersData))
			for i := range ctx.ClustersData {
				result[i] = ctx.ClustersData[i].Name
			}

			return result
		}()

		widget.UpdateData(clustersName)
		v, err := widget.Get(g)
		if err != nil {
			return nil
		}

		_, _ = g.SetCurrentView(v.Name())
		return nil
	}, maxX/4, maxY/4, 3*maxX/4, 3*maxY/4)

	viewClusterList := gui.NewView(viewClusterListId, func(ctx *gui.Context, g *gocui.Gui) error {
		view, err := ctx.SetCurrentView(viewClusterListId)
		if err != nil {
			return err
		}

		w, err := view.Widget(widgetClusterListId)
		if err != nil {
			return err
		}

		if err := w.Update(ctx, g); err != nil {
			return err
		}

		return nil
	}, widgetClusterList)
	ctx.Views = append(ctx.Views, viewClusterList)

	//
	widgetServiceList := gui.NewWidget(widgetServiceListId, func(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
		cluster := ctx.ActiveCluster

		if err := cluster.FetchServices(); err != nil {
			return err
		}

		servicesName := func() []string {
			result := make([]string, len(cluster.Services))
			for i := range cluster.Services {
				result[i] = cluster.Services[i].Name
			}

			return result
		}()

		widget.UpdateData(servicesName)
		v, err := widget.Get(g)
		if err != nil {
			return err
		}

		_, err = g.SetCurrentView(v.Name())

		return err
	},
		1, 1, maxX/3-1, maxY-1,
	)

	widgetServiceDetail := gui.NewWidget(widgetServiceDetailId,
		func(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
			service := ctx.ActiveService
			if service == nil {
				return nil
			}

			widget.UpdateData(service.Render())

			return nil
		}, 1*maxX/3, 1, maxX-1, maxY-1,
	)

	viewServiceList := gui.NewView(viewServiceListId, func(ctx *gui.Context, g *gocui.Gui) error {
		view, err := ctx.SetCurrentView(viewServiceListId)
		if err != nil {
			return err
		}

		w, err := view.Widget(widgetServiceListId)
		if err != nil {
			return err
		}

		if err := w.Update(ctx, g); err != nil {
			return err
		}

		w, err = view.Widget(widgetServiceDetailId)
		if err != nil {
			return err
		}

		if err := w.Update(ctx, g); err != nil {
			return err
		}

		return nil
	}, widgetServiceDetail, widgetServiceList)
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
		_, err := ctx.SetCurrentView(viewClusterListId) // fallback
		return err
	}

	return ctx.CurrentView.Show(ctx, g)
}

func render(ctx *gui.Context, g *gocui.Gui) error {
	return ctx.CurrentView.Render(g)
}
