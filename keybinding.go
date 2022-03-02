package main

import "github.com/jroimartin/gocui"

func keybindings(g *gocui.Gui) error {
	/* --- GLOBAL --- */
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, nextLine); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, prevLine); err != nil {
		return err
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, back); err != nil {
		return err
	}

	/* --- viewClusterListId --- */
	if err := g.SetKeybinding(widgetClusterListId, gocui.MouseLeft, gocui.ModNone, widgetClusterListClickMouse); err != nil {
		return err
	}

	if err := g.SetKeybinding(widgetClusterListId, gocui.KeyEnter, gocui.ModNone, widgetClusterListClick); err != nil {
		return err
	}

	/* --- viewServiceListId --- */
	if err := g.SetKeybinding(widgetServiceListId, gocui.KeyEnter, gocui.ModNone, widgetServiceListClick); err != nil {
		return err
	}

	return nil
}
