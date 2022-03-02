package main

import (
	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/context"
)

func widgetServiceDetailsShow(ctx *context.Context, g *gocui.Gui) error {
	service := ctx.ActiveService
	if service == nil {
		return nil
	}

	widgetServiceDetail.UpdateData(service.Render())

	return nil
}
