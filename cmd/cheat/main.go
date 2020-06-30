package main

//go:generate go run ../../build/embed.go

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/mitchellh/go-homedir"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/installer"
)

const version = "4.0.1"

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

	// identify the os-specifc paths at which configs may be located
	confpaths, err := config.Paths(runtime.GOOS, home, envvars)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	// search for the config file in the above paths
	confpath, err := config.Path(confpaths)
	if err != nil {
		// prompt the user to create a config file
		yes, err := installer.Prompt(
			"A config file was not found. Would you like to create one now? [Y/n]",
			true,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create config: %v\n", err)
			os.Exit(1)
		}

		// exit early on a negative answer
		if !yes {
			os.Exit(0)
		}

		// read the config template
		configs := configs()

		// determine the appropriate paths for config data and (optional) community
		// cheatsheets based on the user's platform
		confpath = confpaths[0]
		confdir := path.Dir(confpath)

		// create paths for community and personal cheatsheets
		community := path.Join(confdir, "/cheatsheets/community")
		personal := path.Join(confdir, "/cheatsheets/personal")

		// template the above paths into the default configs
		configs = strings.Replace(configs, "COMMUNITY_PATH", community, -1)
		configs = strings.Replace(configs, "PERSONAL_PATH", personal, -1)

		// prompt the user to download the community cheatsheets
		yes, err = installer.Prompt(
			"Would you like to download the community cheatsheets? [Y/n]",
			true,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create config: %v\n", err)
			os.Exit(1)
		}

		// clone the community cheatsheets if so instructed
		if yes {
			// clone the community cheatsheets
			if err := installer.Clone(community); err != nil {
				fmt.Fprintf(os.Stderr, "failed to create config: %v\n", err)
				os.Exit(1)
			}

			// also create a directory for personal cheatsheets
			if err := os.MkdirAll(personal, os.ModePerm); err != nil {
				fmt.Fprintf(
					os.Stderr,
					"failed to create config: failed to create directory: %s: %v\n",
					personal,
					err)
				os.Exit(1)
			}
		}

		// the config file does not exist, so we'll try to create one
		if err = config.Init(confpath, configs); err != nil {
			fmt.Fprintf(
				os.Stderr,
				"failed to create config file: %s: %v\n",
				confpath,
				err,
			)
			os.Exit(1)
		}

		fmt.Printf("Created config file: %s\n", confpath)
		fmt.Println("Please read this file for advanced configuration information.")
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
