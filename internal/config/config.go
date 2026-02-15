// Package config implements functions pertaining to configuration management.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	cp "github.com/cheat/cheat/internal/cheatpath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

// Config encapsulates configuration parameters
type Config struct {
	Colorize   bool           `yaml:"colorize"`
	Editor     string         `yaml:"editor"`
	Cheatpaths []cp.Cheatpath `yaml:"cheatpaths"`
	Style      string         `yaml:"style"`
	Formatter  string         `yaml:"formatter"`
	Pager      string         `yaml:"pager"`
	Path       string
}

// New returns a new Config struct
func New(_ map[string]interface{}, confPath string, resolve bool) (Config, error) {

	// read the config file
	buf, err := os.ReadFile(confPath)
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file: %v", err)
	}

	// initialize a config object
	conf := Config{}

	// store the config path
	conf.Path = confPath

	// unmarshal the yaml
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshal yaml: %v", err)
	}

	// if a .cheat directory exists in the current directory or any ancestor,
	// append it to the cheatpaths
	cwd, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get cwd: %v", err)
	}

	if local := findLocalCheatpath(cwd); local != "" {
		path := cp.Cheatpath{
			Name:     "cwd",
			Path:     local,
			ReadOnly: false,
			Tags:     []string{},
		}
		conf.Cheatpaths = append(conf.Cheatpaths, path)
	}

	// process cheatpaths
	var validPaths []cp.Cheatpath
	for _, cheatpath := range conf.Cheatpaths {

		// expand ~ in config paths
		expanded, err := homedir.Expand(cheatpath.Path)
		if err != nil {
			return Config{}, fmt.Errorf("failed to expand ~: %v", err)
		}

		// follow symlinks
		//
		// NB: `resolve` is an ugly kludge that exists for the sake of unit-tests.
		// It's necessary because `EvalSymlinks` will error if the symlink points
		// to a non-existent location on the filesystem. When unit-testing,
		// however, we don't want to have dependencies on the filesystem. As such,
		// `resolve` is a switch that allows us to turn off symlink resolution when
		// running the config tests.
		if resolve {
			evaled, err := filepath.EvalSymlinks(expanded)
			if err != nil {
				// if the path simply doesn't exist, warn and skip it
				if os.IsNotExist(err) {
					fmt.Fprintf(os.Stderr,
						"WARNING: cheatpath '%s' does not exist, skipping\n",
						expanded,
					)
					continue
				}
				return Config{}, fmt.Errorf(
					"failed to resolve symlink: %s: %v",
					expanded,
					err,
				)
			}

			expanded = evaled
		}

		cheatpath.Path = expanded
		validPaths = append(validPaths, cheatpath)
	}
	conf.Cheatpaths = validPaths

	// determine the editor: env vars override the config file value,
	// following standard Unix convention (see #589)
	if v := os.Getenv("VISUAL"); v != "" {
		conf.Editor = v
	} else if v := os.Getenv("EDITOR"); v != "" {
		conf.Editor = v
	} else {
		conf.Editor = strings.TrimSpace(conf.Editor)
	}

	// if an editor was still not determined, attempt to choose one
	// that's appropriate for the environment
	if conf.Editor == "" {
		if conf.Editor, err = Editor(); err != nil {
			return Config{}, err
		}
	}

	// if a chroma style was not provided, set a default
	if conf.Style == "" {
		conf.Style = "bw"
	}

	// if a chroma formatter was not provided, set a default
	if conf.Formatter == "" {
		conf.Formatter = "terminal"
	}

	// load the pager
	conf.Pager = strings.TrimSpace(conf.Pager)

	return conf, nil
}

// findLocalCheatpath walks upward from dir looking for a .cheat directory.
// It returns the path to the first .cheat directory found, or an empty string
// if none exists. This mirrors the discovery pattern used by git for .git
// directories.
func findLocalCheatpath(dir string) string {
	for {
		candidate := filepath.Join(dir, ".cheat")
		if info, err := os.Stat(candidate); err == nil && info.IsDir() {
			return candidate
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
