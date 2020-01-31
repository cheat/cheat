package config

import (
	"fmt"
	"path"

	"github.com/mitchellh/go-homedir"
)

// Paths returns config file paths that are appropriate for the operating
// system
func Paths(sys string, envvars map[string]string) ([]string, error) {

	// if CHEAT_CONFIG_PATH is set, expand ~ and return it
	if confpath, ok := envvars["CHEAT_CONFIG_PATH"]; ok {

		// expand ~
		expanded, err := homedir.Expand(confpath)
		if err != nil {
			return []string{}, fmt.Errorf("failed to expand ~: %v", err)
		}

		return []string{expanded}, nil
	}

	switch sys {
	case "darwin", "linux", "freebsd":
		return []string{
			path.Join(envvars["XDG_CONFIG_HOME"], "/cheat/conf.yml"),
			path.Join(envvars["HOME"], ".config/cheat/conf.yml"),
			path.Join(envvars["HOME"], ".cheat/conf.yml"),
		}, nil
	case "windows":
		return []string{
			path.Join(envvars["APPDATA"], "/cheat/conf.yml"),
			path.Join(envvars["PROGRAMDATA"], "/cheat/conf.yml"),
		}, nil
	default:
		return []string{}, fmt.Errorf("unsupported os: %s", sys)
	}
}
