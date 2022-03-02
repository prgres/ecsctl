package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
)

func viewClusterListShow(ctx *context.Context, g *gocui.Gui) error {
	w, err := viewClusterList.Widget(widgetClusterListId)
	if err != nil {
		return err
	}

	if err := widgetClusterListShow(ctx, g, w); err != nil {
		return err
	}

	ctx.CurrentView = viewClusterList

	return nil
}
