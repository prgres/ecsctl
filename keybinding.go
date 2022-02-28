package main

import "github.com/jroimartin/gocui"

func keybindings(g *gocui.Gui) error {
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

	if err := g.SetKeybinding(viewClusterId, gocui.MouseLeft, gocui.ModNone, showServicesLayoutMouse); err != nil {
		return err
	}

	if err := g.SetKeybinding(viewClusterId, gocui.KeyEnter, gocui.ModNone, showServicesLayout); err != nil {
		return err
	}

	return nil
}
