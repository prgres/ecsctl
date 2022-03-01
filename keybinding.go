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
	if err := g.SetKeybinding(viewClusterListId, gocui.MouseLeft, gocui.ModNone, layoutClustersClickMouse); err != nil {
		return err
	}

	if err := g.SetKeybinding(viewClusterListId, gocui.KeyEnter, gocui.ModNone, layoutClustersClick); err != nil {
		return err
	}

	/* --- viewServiceListId --- */
	if err := g.SetKeybinding(viewServiceListId, gocui.KeyEnter, gocui.ModNone, layoutServiceListClick); err != nil {
		return err
	}

	return nil
}
