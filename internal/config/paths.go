package config

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

// Paths returns config file paths that are appropriate for the operating
// system
func Paths(sys string) ([]string, error) {

	// if CHEAT_CONFIG_PATH is set, return it
	if os.Getenv("CHEAT_CONFIG_PATH") != "" {

		// expand ~
		expanded, err := homedir.Expand(os.Getenv("CHEAT_CONFIG_PATH"))
		if err != nil {
			return []string{}, fmt.Errorf("failed to expand ~: %v", err)
		}

		return []string{expanded}, nil
	}

	switch sys {
	case "darwin", "linux", "freebsd":
		return []string{
			path.Join(os.Getenv("XDG_CONFIG_HOME"), "/cheat/conf.yml"),
			path.Join(os.Getenv("HOME"), ".config/cheat/conf.yml"),
			path.Join(os.Getenv("HOME"), ".cheat/conf.yml"),
		}, nil
	case "windows":
		return []string{
			fmt.Sprintf("%s/cheat/conf.yml", os.Getenv("APPDATA")),
			fmt.Sprintf("%s/cheat/conf.yml", os.Getenv("PROGRAMDATA")),
		}, nil
	default:
		return []string{}, fmt.Errorf("unsupported os: %s", sys)
	}
}
