package main

import (
	"github.com/jroimartin/gocui"

	"github.com/prgres/ecsctl/gui"
)

func viewServiceListClear(ctx *gui.Context, g *gocui.Gui) error {
	// omit error because it can only return ErrUnknownView which does not bother us at this moment
	_ = g.DeleteView(widgetServiceListId)
	_ = g.DeleteView(widgetServiceDetailId)

	return nil
}
