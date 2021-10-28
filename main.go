package main

import (
	"flag"
	"github.com/Dnnd/sway-window-switcher/dmenu"
	"github.com/Dnnd/sway-window-switcher/dmenu/rofi"
	"github.com/Dnnd/sway-window-switcher/dmenu/wofi"
	"github.com/Dnnd/sway-window-switcher/domain"
	"github.com/Dnnd/sway-window-switcher/swaymsg"
	ji "github.com/json-iterator/go"
	"log"
)

func main() {
	launcher := flag.String("launcher", "rofi", "program to display menu")
	flag.Parse()
	var menuFactory dmenu.MenuFactory

	if *launcher == "rofi" {
		menuFactory = rofi.SwitchWorkspacesMenuFactory{}
	} else if *launcher == "wofi" {
		menuFactory = wofi.SwitchWorkspacesMenuFactory{}
	} else {
		log.Fatal("unknown launcher")
	}

	getTree := swaymsg.NewGetTree()
	tree, err := getTree.Send()
	if err != nil {
		log.Fatal(err)
	}
	var root domain.Node
	if err := ji.Unmarshal(tree, &root); err != nil {
		log.Fatal(err)
	}

	workspaces := domain.ExtractWorkspaces(&root)

	menu := menuFactory.NewMenu()
	menuEntry, err := menu.Show(workspaces.ToMenuData())

	if err != nil {
		log.Fatal(err)
	}
	ordinal, err := domain.ParseDmenuEntryNumber(menuEntry)
	if err != nil {
		log.Fatal(err)
	}
	leafToShow, err := workspaces.Find(ordinal)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = swaymsg.NewFocusWindowMessage(leafToShow.Id).Send(); err != nil {
		log.Fatal(err)
	}
}
