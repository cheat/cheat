package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/cheat/cheat/internal/config"
)

// cmdInit displays an example config file.
func cmdInit() {

	// get the user's home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get user home directory: %v\n", err)
		os.Exit(1)
	}

	// read the envvars into a map of strings
	envvars := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envvars[pair[0]] = pair[1]
	}

	// load the config template
	configs := configs()

	// identify the os-specifc paths at which configs may be located
	confpaths, err := config.Paths(runtime.GOOS, home, envvars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config paths: %v\n", err)
		os.Exit(1)
	}

	// determine the appropriate paths for config data and (optional) community
	// cheatsheets based on the user's platform
	confpath := confpaths[0]
	confdir := path.Dir(confpath)

	// create paths for community and personal cheatsheets
	community := path.Join(confdir, "/cheatsheets/community")
	personal := path.Join(confdir, "/cheatsheets/personal")

	// template the above paths into the default configs
	configs = strings.Replace(configs, "COMMUNITY_PATH", community, -1)
	configs = strings.Replace(configs, "PERSONAL_PATH", personal, -1)

	// output the templated configs
	fmt.Println(configs)
}
