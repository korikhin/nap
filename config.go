package main

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/caarlos0/env/v11"
	"gopkg.in/yaml.v3"
)

// Config holds the configuration options for the application.
//
// At the moment, it is quite limited, only supporting the home folder and the
// file name of the metadata.
type Config struct {
	Home            string `env:"NAP_HOME" yaml:"home"`
	File            string `env:"NAP_FILE" yaml:"file"`
	Editor          string `env:"NAP_EDITOR" yaml:"editor"`
	DefaultLanguage string `env:"NAP_DEFAULT_LANGUAGE" yaml:"default_language"`

	Theme               string `env:"NAP_THEME" yaml:"theme"`
	BackgroundColor     string `env:"NAP_BACKGROUND" yaml:"background"`
	ForegroundColor     string `env:"NAP_FOREGROUND" yaml:"foreground"`
	BlackColor          string `env:"NAP_BLACK" yaml:"black"`
	GrayColor           string `env:"NAP_GRAY" yaml:"gray"`
	BrightGrayColor     string `env:"NAP_BRIGHT_GRAY" yaml:"bright_gray"`
	WhiteColor          string `env:"NAP_WHITE" yaml:"white"`
	PrimaryColor        string `env:"NAP_PRIMARY_COLOR" yaml:"primary_color"`
	PrimaryColorSubdued string `env:"NAP_PRIMARY_COLOR_SUBDUED" yaml:"primary_color_subdued"`
	RedColor            string `env:"NAP_RED" yaml:"red"`
	BrightRedColor      string `env:"NAP_BRIGHT_RED" yaml:"bright_red"`
	GreenColor          string `env:"NAP_GREEN" yaml:"green"`
	BrightGreenColor    string `env:"NAP_BRIGHT_GREEN" yaml:"bright_green"`
	StatusColor         string `env:"NAP_STATUS_COLOR" yaml:"status_color"`
}

func newConfig() Config {
	return Config{
		Home:                defaultHome(),
		File:                "snippets.json",
		DefaultLanguage:     defaultLanguage,
		Editor:              "", // defined at config build
		Theme:               "dracula",
		BackgroundColor:     "235",
		ForegroundColor:     "15",
		BlackColor:          "#373B41",
		GrayColor:           "235",
		BrightGrayColor:     "241",
		WhiteColor:          "#FFFFFF",
		PrimaryColor:        "#AFBEE1",
		PrimaryColorSubdued: "#64708D",
		RedColor:            "#A46060",
		BrightRedColor:      "#E49393",
		GreenColor:          "#527251",
		BrightGreenColor:    "#BCE1AF",
	}
}

// default helpers for the configuration.
// We use $XDG_DATA_HOME to avoid cluttering the user's home directory.
func defaultHome() string { return filepath.Join(xdg.DataHome, "nap") }

// defaultConfig returns the default config path
func defaultConfig() string {
	if c := os.Getenv("NAP_CONFIG"); c != "" {
		return c
	}
	if cfgPath, err := xdg.ConfigFile("nap/config.yaml"); err == nil {
		return cfgPath
	}
	return "config.yaml"
}

// readConfig returns a configuration read from the environment.
func readConfig() Config {
	config := newConfig()
	fi, err := os.Open(defaultConfig())
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return newConfig()
	}
	if fi != nil {
		defer fi.Close()
		if err := yaml.NewDecoder(fi).Decode(&config); err != nil {
			return newConfig()
		}
	}
	if err := env.Parse(&config); err != nil {
		return newConfig()
	}

	if strings.HasPrefix(config.Home, "~") {
		home, err := os.UserHomeDir()
		if err == nil {
			config.Home = filepath.Join(home, config.Home[1:])
		}
	}
	if config.Editor == "" {
		config.Editor = getEditor()
	}

	return config
}

// writeConfig returns a configuration read from the environment.
func (config Config) writeConfig() error {
	fi, err := os.Create(defaultConfig())
	if err != nil {
		return err
	}
	if fi != nil {
		defer fi.Close()
		if err := yaml.NewEncoder(fi).Encode(&config); err != nil {
			return err
		}
	}

	return nil
}
