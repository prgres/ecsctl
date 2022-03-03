package main

import (
	"github.com/jroimartin/gocui"
)

// func widgetClusterListShow(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
// 	clustersName := func() []string {
// 		result := make([]string, len(ctx.ClustersData))
// 		for i := range ctx.ClustersData {
// 			result[i] = ctx.ClustersData[i].Name
// 		}

// 		return result
// 	}()

// 	widget.UpdateData(clustersName)
// 	v, err := widget.Get(g)
// 	if err != nil {
// 		return nil
// 	}

// 	_, _ = g.SetCurrentView(v.Name())
// 	return nil
// }

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

	nextView, err := ctx.View(viewServiceListId)
	if err != nil {
		return err
	}

	if nextView.Show(ctx, g); err != nil {
		return err
	}

	return g.DeleteView(v.Name())
}
