package main

import (
	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/gui"
)

/* --- keybinding func --- */
var prevLineMouseClick = ""

func widgetClusterListClickMouse(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()
	_, cy := v.Cursor()

	clusterId, err := v.Line(cy)
	if err != nil {
		return nil
	}

	if prevLineMouseClick == clusterId {
		if _, err := ctx.SetActiveClusterId(clusterId); err != nil {
			return err
		}

		return ctx.CurrentView.Show(ctx, g)
	}

	if prevLineMouseClick == "" {
		prevLineMouseClick = clusterId
	}

	return nil
}

func widgetClusterListClick(g *gocui.Gui, v *gocui.View) error {
	ctx := _ctx.Context()
	_, cy := v.Cursor()

	clusterId, err := v.Line(cy)
	if err != nil {
		return nil
	}

	if _, err := ctx.SetActiveClusterId(clusterId); err != nil {
		return err
	}

	nextView, err := ctx.View(gui.ViewServiceListId)
	if err != nil {
		return err
	}

	if nextView.Show(ctx, g); err != nil {
		return err
	}

	return g.DeleteView(v.Name())
}
