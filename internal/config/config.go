// Package config implements functions pertaining to configuration management.
package config

import (
	cp "github.com/cheat/cheat/internal/cheatpath"
)

// Config encapsulates configuration parameters
type Config struct {
	Colorize   bool      `yaml:"colorize"`
	Editor     string    `yaml:"editor"`
	Cheatpaths []cp.Path `yaml:"cheatpaths"`
	Style      string    `yaml:"style"`
	Formatter  string    `yaml:"formatter"`
	Pager      string    `yaml:"pager"`
	Path       string
}
