package config

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	cp "github.com/cheat/cheat/internal/cheatpath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

// Config encapsulates configuration parameters
type Config struct {
	Colorize   bool           `yaml:"colorize"`
	Editor     string         `yaml:"editor"`
	Cheatpaths []cp.Cheatpath `yaml:"cheatpaths"`
	Style      string         `yaml:"style"`
	Formatter  string         `yaml:"formatter"`
	Pager      string         `yaml:"pager"`
}

// New returns a new Config struct
func New(opts map[string]interface{}, confPath string, resolve bool) (Config, error) {

	// read the config file
	buf, err := os.ReadFile(confPath)
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file: %v", err)
	}

	// initialize a config object
	conf := Config{}

	// unmarshal the yaml
	err = yaml.UnmarshalStrict(buf, &conf)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshal yaml: %v", err)
	}

	// if a .cheat directory exists locally, append it to the cheatpaths
	cwd, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get cwd: %v", err)
	}

	local := filepath.Join(cwd, ".cheat")
	if _, err := os.Stat(local); err == nil {
		path := cp.Cheatpath{
			Name:     "cwd",
			Path:     local,
			ReadOnly: false,
			Tags:     []string{},
		}

		conf.Cheatpaths = append(conf.Cheatpaths, path)
	}

	// process cheatpaths
	for i, cheatpath := range conf.Cheatpaths {

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
				return Config{}, fmt.Errorf(
					"failed to resolve symlink: %s: %v",
					expanded,
					err,
				)
			}

			expanded = evaled
		}

		conf.Cheatpaths[i].Path = expanded
	}

	// if an editor was not provided in the configs, look to envvars
	if conf.Editor == "" {
		if os.Getenv("VISUAL") != "" {
			conf.Editor = os.Getenv("VISUAL")
		} else if os.Getenv("EDITOR") != "" {
			conf.Editor = os.Getenv("EDITOR")
		} else if runtime.GOOS == "windows" {
			conf.Editor = "notepad"
		} else {
			// try to fall back to `nano`
			path, err := exec.LookPath("nano")
			if err != nil {
				return Config{}, fmt.Errorf("failed to locate nano: %s", err)
			}

			// use `nano` if we found it
			if path != "" {
				conf.Editor = "nano"
				// otherwise, give up
			} else {
				return Config{}, fmt.Errorf("no editor set")
			}
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

	// attempt to fall back to `PAGER` if a pager is not specified in configs
	conf.Pager = strings.TrimSpace(conf.Pager)
	if conf.Pager == "" {
		// look for `pager`, `less`, and `more` on the system PATH
		pagerPath, _ := exec.LookPath("pager")
		lessPath, _ := exec.LookPath("less")
		morePath, _ := exec.LookPath("more")

		// search first for a `PAGER` envvar
		if os.Getenv("PAGER") != "" {
			conf.Pager = os.Getenv("PAGER")

			// search for `pager`
		} else if pagerPath != "" {
			conf.Pager = pagerPath

			// search for `less`
		} else if lessPath != "" {
			conf.Pager = lessPath

			// search for `more`
			//
			// XXX: this causes issues on some Linux systems. See:
			// https://github.com/cheat/cheat/issues/681#issuecomment-1201842334
			//
			// By checking for `more` last, we're hoping to at least mitigate
			// the frequency of this occurrence, because `pager` and `less` are
			// likely to be available on most systems on which a user is likely
			// to have installed `cheat`.
		} else if morePath != "" {
			conf.Pager = morePath
		}
	}

	return conf, nil
}
