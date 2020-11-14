package main

import (
	"github.com/Dnnd/sway-window-switcher/dmenu/wofi"
	"github.com/Dnnd/sway-window-switcher/domain/node"
	"github.com/Dnnd/sway-window-switcher/domain/workspace"
	"github.com/Dnnd/sway-window-switcher/swaymsg"
	ji "github.com/json-iterator/go"
	"log"
)

func main() {
	getTree := swaymsg.NewGetTree()
	tree, err := getTree.Send()
	if err != nil {
		log.Fatal(err)
	}
	var root node.Node
	if err := ji.Unmarshal(tree, &root); err != nil {
		log.Fatal(err)
	}

	workspaces := workspace.ExtractWorkspaces(&root)

	dmenu := wofi.NewWofiDmenuCommand(workspaces.ToDmenu())
	dmenuEntry, err := dmenu.Show()
	if err != nil {
		log.Fatal(err)
	}
	ordinal, err := wofi.ParseWofiDmenuEntryNumber(dmenuEntry)
	if err != nil {
		log.Fatal(err)
	}
	leaveToShow, err := workspaces.Find(ordinal)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = swaymsg.NewFocusWindowMessage(leaveToShow.Id).Send(); err != nil {
		log.Fatal(err)
	}
}
