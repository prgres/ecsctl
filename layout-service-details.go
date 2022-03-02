package main

import (
	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/context"
	"github.com/prgres/ecsctl/widget"
)

func widgetServiceDetailsShow(ctx *context.Context, g *gocui.Gui, widget *widget.Widget) error {
	service := ctx.ActiveService
	if service == nil {
		return nil
	}

	widget.UpdateData(service.Render())

	return nil
}
