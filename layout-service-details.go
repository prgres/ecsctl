package main

import (
	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/context"
)

func layoutServiceDetailsShow(ctx *context.Context, g *gocui.Gui) error {
	service := ctx.ActiveService
	_, err := layoutServiceDetail.Render(g, service.Render())
	if err != nil {
		return err
	}

	return nil
}
