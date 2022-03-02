package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
)

func viewClusterListShow(ctx *context.Context, g *gocui.Gui) error {
	if err := widgetClusterListShow(ctx, g); err != nil {
		return err
	}

	ctx.CurrentView = viewClusterList

	return nil
}
