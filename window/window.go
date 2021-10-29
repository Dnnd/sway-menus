package window

import (
	"github.com/Dnnd/sway-menus/dmenu"
	"github.com/Dnnd/sway-menus/domain/workspace"
	"github.com/Dnnd/sway-menus/swaymsg"
	ji "github.com/json-iterator/go"
)

type SwitcherSerivce struct {
	launcher dmenu.MenuPresenter
}

func NewSwitcherService(launcher dmenu.MenuPresenter) *SwitcherSerivce {
	return &SwitcherSerivce{launcher: launcher}
}

func (sw *SwitcherSerivce) Run() error {
	getTree := swaymsg.NewGetTree()
	tree, err := getTree.Send()
	if err != nil {
		return err
	}
	var root workspace.Node
	if err := ji.Unmarshal(tree, &root); err != nil {
		return err
	}

	workspaces := workspace.ExtractWorkspaces(&root)
	menuEntryNum, err := sw.launcher.Show(workspaces.ToMenuData())

	if err != nil {
		return err
	}

	leafToShow, err := workspaces.Find(menuEntryNum)
	if err != nil {
		return err
	}

	_, err = swaymsg.NewFocusWindowMessage(leafToShow.Id).Send()
	return err
}
