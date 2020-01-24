package config

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

// Path returns the config file path
func Path(sys string) (string, error) {

	var paths []string

	// if CHEAT_CONFIG_PATH is set, return it
	if os.Getenv("CHEAT_CONFIG_PATH") != "" {

		// expand ~
		expanded, err := homedir.Expand(os.Getenv("CHEAT_CONFIG_PATH"))
		if err != nil {
			return "", fmt.Errorf("failed to expand ~: %v", err)
		}

		return expanded, nil
	}

	switch sys {
	case "darwin", "linux", "freebsd":
		paths = []string{
			path.Join(os.Getenv("XDG_CONFIG_HOME"), "/cheat/conf.yml"),
			path.Join(os.Getenv("HOME"), ".config/cheat/conf.yml"),
			path.Join(os.Getenv("HOME"), ".cheat/conf.yml"),
		}
	case "windows":
		paths = []string{
			fmt.Sprintf("%s/cheat/conf.yml", os.Getenv("APPDATA")),
			fmt.Sprintf("%s/cheat/conf.yml", os.Getenv("PROGRAMDATA")),
		}
	default:
		return "", fmt.Errorf("unsupported os: %s", sys)
	}

	// check if the config file exists on any paths
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	// we can't find the config file if we make it this far
	return "", fmt.Errorf("could not locate config file")
}
