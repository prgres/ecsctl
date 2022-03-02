package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
)

func viewClusterListShow(ctx *context.Context, g *gocui.Gui) error {
	view, err := ctx.SetCurrentView(viewClusterListId)
	if err != nil {
		return err
	}

	w, err := view.Widget(widgetClusterListId)
	if err != nil {
		return err
	}

	if err := widgetClusterListShow(ctx, g, w); err != nil {
		return err
	}

	return nil
}
