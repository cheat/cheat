package completions

import (
	"runtime"

	"github.com/mitchellh/go-homedir"

	"github.com/cheat/cheat/internal/config"
)

// loadConfig loads the cheat configuration for use in completion functions.
// It returns an error rather than exiting, since completions should degrade
// gracefully.
func loadConfig() (config.Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return config.Config{}, err
	}

	envvars := config.EnvVars()

	confpaths, err := config.Paths(runtime.GOOS, home, envvars)
	if err != nil {
		return config.Config{}, err
	}

	confpath, err := config.Path(confpaths)
	if err != nil {
		return config.Config{}, err
	}

	conf, err := config.New(confpath, true)
	if err != nil {
		return config.Config{}, err
	}

	return conf, nil
}
