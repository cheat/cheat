package config

import (
	"fmt"
	"path"

	"github.com/mitchellh/go-homedir"
)

// Paths returns config file paths that are appropriate for the operating
// system
func Paths(sys string, envvars map[string]string) ([]string, error) {

	// get the user's home directory
	home, err := homedir.Dir()
	if err != nil {
		return []string{}, fmt.Errorf("failed to get user home directory: %v", err)
	}

	// if `CHEAT_CONFIG_PATH` is set, expand ~ and return it
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
		paths := []string{}

		// don't include the `XDG_CONFIG_HOME` path if that envvar is not set
		if xdgpath, ok := envvars["XDG_CONFIG_HOME"]; ok {
			paths = append(paths, path.Join(xdgpath, "/cheat/conf.yml"))
		}

		// if `XDG_CONFIG_HOME` is not set, search the user's home directory
		paths = append(paths, []string{
			path.Join(home, ".config/cheat/conf.yml"),
			path.Join(home, ".cheat/conf.yml"),
		}...)

		return paths, nil
	case "windows":
		return []string{
			path.Join(envvars["APPDATA"], "/cheat/conf.yml"),
			path.Join(envvars["PROGRAMDATA"], "/cheat/conf.yml"),
		}, nil
	default:
		return []string{}, fmt.Errorf("unsupported os: %s", sys)
	}
}
