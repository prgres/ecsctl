package main

import (
	"github.com/jroimartin/gocui"
	"github.com/prgres/ecsctl/gui"
)

func viewClusterList(g *gocui.Gui) *gui.View {
	maxX, maxY := g.Size()

	widgetLogoBanner := gui.NewWidget(gui.WidgetLogoBannerId, func(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
		widget.UpdateData([]string{
			"EEEEEEEEEEEEEEEEEEEEEE      CCCCCCCCCCCCC  SSSSSSSSSSSSSSS        CCCCCCCCCCCCTTTTTTTTTTTTTTTTTTTTTTLLLLLLLLLLL             ",
			"E::::::::::::::::::::E   CCC::::::::::::CSS:::::::::::::::S    CCC::::::::::::T:::::::::::::::::::::L:::::::::L             ",
			"E::::::::::::::::::::E CC:::::::::::::::S:::::SSSSSS::::::S  CC:::::::::::::::T:::::::::::::::::::::L:::::::::L             ",
			"EE::::::EEEEEEEEE::::EC:::::CCCCCCCC::::S:::::S     SSSSSSS C:::::CCCCCCCC::::T:::::TT:::::::TT:::::LL:::::::LL             ",
			"  E:::::E       EEEEEC:::::C       CCCCCS:::::S            C:::::C       CCCCCTTTTTT  T:::::T  TTTTTT L:::::L               ",
			"  E:::::E           C:::::C             S:::::S           C:::::C                     T:::::T         L:::::L               ",
			"  E::::::EEEEEEEEEE C:::::C              S::::SSSS        C:::::C                     T:::::T         L:::::L               ",
			"  E:::::::::::::::E C:::::C               SS::::::SSSSS   C:::::C                     T:::::T         L:::::L               ",
			"  E:::::::::::::::E C:::::C                 SSS::::::::SS C:::::C                     T:::::T         L:::::L               ",
			"  E::::::EEEEEEEEEE C:::::C                    SSSSSS::::SC:::::C                     T:::::T         L:::::L               ",
			"  E:::::E           C:::::C                         S:::::C:::::C                     T:::::T         L:::::L               ",
			"  E:::::E       EEEEEC:::::C       CCCCCC           S:::::SC:::::C       CCCCCC       T:::::T         L:::::L         LLLLLL",
			"EE::::::EEEEEEEE:::::EC:::::CCCCCCCC::::SSSSSSS     S:::::S C:::::CCCCCCCC::::C     TT:::::::TT     LL:::::::LLLLLLLLL:::::L",
			"E::::::::::::::::::::E CC:::::::::::::::S::::::SSSSSS:::::S  CC:::::::::::::::C     T:::::::::T     L::::::::::::::::::::::L",
			"E::::::::::::::::::::E   CCC::::::::::::S:::::::::::::::SS     CCC::::::::::::C     T:::::::::T     L::::::::::::::::::::::L",
			"EEEEEEEEEEEEEEEEEEEEEE      CCCCCCCCCCCCCSSSSSSSSSSSSSSS          CCCCCCCCCCCCC     TTTTTTTTTTT     LLLLLLLLLLLLLLLLLLLLLLLL",
		})

		return nil
	}, maxX/8, 7*maxX/8, 1, maxY/4)

	//TODO: better way to edit
	viewWidgetLogoBanner, err := widgetLogoBanner.Get(g)
	if err != nil {
		panic(err)
	}

	viewWidgetLogoBanner.Autoscroll = false
	viewWidgetLogoBanner.Highlight = false
	viewWidgetLogoBanner.Editable = false
	viewWidgetLogoBanner.Frame = false
	//

	widgetClusterList := gui.NewWidget(gui.WidgetClusterListId, func(ctx *gui.Context, g *gocui.Gui, widget *gui.Widget) error {
		clustersName := func() []string {
			result := make([]string, len(ctx.ClustersData))
			for i := range ctx.ClustersData {
				result[i] = ctx.ClustersData[i].Name
			}

			return result
		}()

		widget.UpdateData(clustersName)
		v, err := widget.Get(g)
		if err != nil {
			return nil
		}

		_, _ = g.SetCurrentView(v.Name())
		return nil
	}, maxX/4, 3*maxX/4, maxY/4, 3*maxY/4)

	viewClusterList := gui.NewView(gui.ViewClusterListId, func(ctx *gui.Context, g *gocui.Gui) error {
		view, err := ctx.SetCurrentView(gui.ViewClusterListId)
		if err != nil {
			return err
		}

		w, err := view.Widget(gui.WidgetClusterListId)
		if err != nil {
			return err
		}

		if err := w.Update(ctx, g); err != nil {
			return err
		}

		return nil
	}, widgetClusterList, widgetLogoBanner)

	return viewClusterList
}

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
