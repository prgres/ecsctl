package main

import (
	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/gui"
)

func viewClusterList(g *gocui.Gui) *gui.View {
	maxX, maxY := g.Size()

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

	return viewClusterList
}
