package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/installer"
)

// cmdInit displays an example config file.
func cmdInit(home string, envvars map[string]string) {

	// identify the os-specific paths at which configs may be located
	confpaths, err := config.Paths(runtime.GOOS, home, envvars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config paths: %v\n", err)
		os.Exit(1)
	}

	confpath := confpaths[0]

	// expand template placeholders and comment out community cheatpath
	configs := installer.ExpandTemplate(configs(), confpath)
	configs = installer.CommentCommunity(configs, confpath)

	// output the templated configs
	fmt.Println(configs)
}
