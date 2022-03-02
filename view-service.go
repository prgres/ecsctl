package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/context"
)

func viewServiceListShow(ctx *context.Context, g *gocui.Gui) error {
	view, err := ctx.SetCurrentView(viewServiceListId)
	if err != nil {
		return err
	}

	w, err := view.Widget(widgetServiceListId)
	if err != nil {
		return err
	}

	if err := widgetServiceListShow(ctx, g, w); err != nil {
		return err
	}

	w, err = view.Widget(widgetServiceDetailId)
	if err != nil {
		return err
	}

	if err := widgetServiceDetailsShow(ctx, g, w); err != nil {
		return err
	}

	return nil
}

func viewServiceListClear(ctx *context.Context, g *gocui.Gui) error {
	// omit error because it can only return ErrUnknownView which does not bother us at this moment
	_ = g.DeleteView(widgetServiceListId)
	_ = g.DeleteView(widgetServiceDetailId)

	return nil
}
