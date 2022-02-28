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
	if len(g.Views()) == 1 && g.Views()[0].Name() == viewClusterId {
		return quit(g, v)
	}

	g.DeleteView(g.Views()[len(g.Views())-1].Name())
	return nil
}
