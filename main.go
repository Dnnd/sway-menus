package main

import (
	"flag"
	"github.com/Dnnd/sway-window-switcher/dmenu"
	"github.com/Dnnd/sway-window-switcher/dmenu/rofi"
	"github.com/Dnnd/sway-window-switcher/dmenu/wofi"
	"github.com/Dnnd/sway-window-switcher/logind"
	"github.com/Dnnd/sway-window-switcher/powermenu"
	"github.com/Dnnd/sway-window-switcher/window"
	"log"
	"os"
)

func main() {

	mode := flag.String("mode", "window", "mode")
	var configFile string
	flag.StringVar(&configFile, "config", "", "path to config file")
	flag.Parse()
	if configFile == "" {
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome == "" {
			xdgConfigHome = os.Getenv("HOME") + "/.config"
		}
		var configDir = xdgConfigHome + "/sway_menus"
		configFile = configDir + "/config.toml"
	}
	config, err := LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var menuFactory dmenu.MenuFactory
	if config.Launcher == "rofi" {
		menuFactory = rofi.MenuFactory{}
	} else if config.Launcher == "wofi" {
		menuFactory = wofi.MenuFactory{}
	} else {
		log.Fatal("unknown launcher")
	}

	if *mode == "window" {
		switcher := window.NewSwitcherService(menuFactory.NewMenuPresenter())
		if err := switcher.Run(); err != nil {
			log.Fatal(err)
		}
	} else if *mode == "powermenu" {
		logindManager, err := logind.ConnectToLogindManager()
		if err != nil {
			log.Fatal(err)
		}
		powermenuService := powermenu.NewPowerMenuService(config.Powermenu, &logindManager, menuFactory.NewMenuPresenter())
		if err := powermenuService.Run(); err != nil {
			log.Fatal(err)
		}
	}

}
