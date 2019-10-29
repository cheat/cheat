package main

//go:generate go run ../../build/embed.go

import (
	"fmt"
	"os"
	"runtime"

	"github.com/docopt/docopt-go"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/config"
)

const version = "3.0.3"

func main() {

	// initialize options
	opts, err := docopt.Parse(usage(), nil, true, version, false)
	if err != nil {
		// panic here, because this should never happen
		panic(fmt.Errorf("docopt failed to parse: %v", err))
	}

	// if --init was passed, we don't want to attempt to load a config file.
	// Instead, just execute cmd_init and exit
	if opts["--init"] != nil && opts["--init"] == true {
		cmdInit()
		os.Exit(0)
	}

	// load the config file
	confpath, err := config.Path(runtime.GOOS)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not locate config file")
		os.Exit(1)
	}

	// initialize the configs
	conf, err := config.New(opts, confpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// assert that the configs are valid
	if err := conf.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// filter the cheatpaths if --path was passed
	if opts["--path"] != nil {
		conf.Cheatpaths, err = cheatpath.Filter(
			conf.Cheatpaths,
			opts["--path"].(string),
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid option --path: %v\n", err)
			os.Exit(1)
		}
	}

	// determine which command to execute
	var cmd func(map[string]interface{}, config.Config)

	switch {
	case opts["--directories"].(bool):
		cmd = cmdDirectories

	case opts["--edit"] != nil:
		cmd = cmdEdit

	case opts["--list"].(bool):
		cmd = cmdList

	case opts["--search"] != nil:
		cmd = cmdSearch

	case opts["<cheatsheet>"] != nil:
		cmd = cmdView

	default:
		fmt.Println(usage())
		os.Exit(0)
	}

	// execute the command
	cmd(opts, conf)
}
