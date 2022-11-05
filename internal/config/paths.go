package config

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// Paths returns config file paths that are appropriate for the operating
// system
func Paths(
	sys string,
	home string,
	envvars map[string]string,
) ([]string, error) {

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

	// darwin/linux/unix
	case "aix", "android", "darwin", "dragonfly", "freebsd", "illumos", "ios",
		"linux", "netbsd", "openbsd", "plan9", "solaris":
		paths := []string{}

		// don't include the `XDG_CONFIG_HOME` path if that envvar is not set
		if xdgpath, ok := envvars["XDG_CONFIG_HOME"]; ok {
			paths = append(paths, filepath.Join(xdgpath, "cheat", "conf.yml"))
		}

		paths = append(paths, []string{
			filepath.Join(home, ".config", "cheat", "conf.yml"),
			filepath.Join(home, ".cheat", "conf.yml"),
			"/etc/cheat/conf.yml",
		}...)

		return paths, nil

	// windows
	case "windows":
		return []string{
			filepath.Join(envvars["APPDATA"], "cheat", "conf.yml"),
			filepath.Join(envvars["PROGRAMDATA"], "cheat", "conf.yml"),
		}, nil
	default:
		return []string{}, fmt.Errorf("unsupported os: %s", sys)
	}
}
