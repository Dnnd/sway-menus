package main

import (
	"bufio"
	"github.com/Dnnd/sway-window-switcher/dmenu/wofi"
	"github.com/Dnnd/sway-window-switcher/domain/node"
	"github.com/Dnnd/sway-window-switcher/domain/workspace"
	"github.com/Dnnd/sway-window-switcher/swaymsg"
	ji "github.com/json-iterator/go"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var root node.Node
	err := ji.NewDecoder(r).Decode(&root)
	if err != nil {
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
	_, err = swaymsg.NewFocusWindowMessage(leaveToShow.Id).Send()
	if err != nil {
		log.Fatal(err)
	}
}
