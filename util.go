package main

import "github.com/jroimartin/gocui"

func nextLine(g *gocui.Gui, v *gocui.View) error {
	cx, cy := v.Cursor()
	nextLine, err := v.Line(cy + 1)
	if err != nil {
		nextLine = ""
	}
	if nextLine == "" {
		return nil
	}

	if err := v.SetCursor(cx, cy+1); err != nil {
		ox, oy := v.Origin()
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}

	return nil
}

func prevLine(g *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	cx, cy := v.Cursor()
	if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func back(g *gocui.Gui, v *gocui.View) error {
	switch {
	case len(g.Views()) == 0:
		return quit(g, v)

	case g.CurrentView().Name() == viewClusterListId:
		return quit(g, v)

	case g.CurrentView().Name() == viewServiceListId:
		// omit error because it can only return ErrUnknownView which does not bother us at this moment
		_ = g.DeleteView(viewServiceListId)
		_ = g.DeleteView(viewServiceDetailId)

		return layoutClusterListShow(_ctx.Context(), g)

	default:
		return nil
	}
}
