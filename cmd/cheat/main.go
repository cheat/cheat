package main

//go:generate go run ../../build/embed.go

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/docopt/docopt-go"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/config"
)

const version = "3.5.0"

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

	// read the envvars into a map of strings
	envvars := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envvars[pair[0]] = pair[1]
	}

	// load the os-specifc paths at which the config file may be located
	confpaths, err := config.Paths(runtime.GOOS, envvars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// search for the config file in the above paths
	confpath, err := config.Path(confpaths)
	if err != nil {

		// the config file does not exist, so we'll try to create one
		if err = config.Init(confpaths[0], configs()); err != nil {
			fmt.Fprintf(
				os.Stderr,
				"failed to create config file: %s: %v\n",
				confpaths[0],
				err,
			)
			os.Exit(1)
		}

		confpath = confpaths[0]

		fmt.Printf("Created config file: %s\n", confpath)
		fmt.Println("Please edit this file now to configure cheat.")
		os.Exit(0)
	}

	// initialize the configs
	conf, err := config.New(opts, confpath, true)
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

	case opts["--tags"].(bool):
		cmd = cmdTags

	case opts["--search"] != nil:
		cmd = cmdSearch

	case opts["--rm"] != nil:
		cmd = cmdRemove

	case opts["<cheatsheet>"] != nil:
		cmd = cmdView

	default:
		fmt.Println(usage())
		os.Exit(0)
	}

	// execute the command
	cmd(opts, conf)
}
