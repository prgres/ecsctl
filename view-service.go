package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
)

func viewServiceListShow(ctx *context.Context, g *gocui.Gui) error {
	if err := widgetServiceListShow(ctx, g); err != nil {
		return err
	}

	if err := widgetServiceDetailsShow(ctx, g); err != nil {
		return err
	}

	ctx.CurrentView = viewServiceList

	return nil
}

func viewServiceListClear(ctx *context.Context, g *gocui.Gui) error {
	// omit error because it can only return ErrUnknownView which does not bother us at this moment
	_ = g.DeleteView(widgetServiceListId)
	_ = g.DeleteView(widgetServiceDetailId)

	return nil
}
