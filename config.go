package main

import (
	generic "github.com/Dnnd/sway-window-switcher/domain/generic_menu"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	Powermenu generic.GenericMenu
	Launcher  string
}

func LoadConfig(sourceFile string) (*Config, error) {
	reader, err := os.Open(sourceFile)
	if err != nil {
		return nil, err
	}
	var c Config
	err = toml.NewDecoder(reader).Decode(&c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
