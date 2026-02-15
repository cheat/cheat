// Package main serves as the executable entrypoint.
package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/completions"
	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/installer"
)

const version = "5.0.0"

var rootCmd = &cobra.Command{
	Use:   "cheat [cheatsheet]",
	Short: "Create and view interactive cheatsheets on the command-line",
	Long: `cheat allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind *nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.`,
	Example: `  To initialize a config file:
    mkdir -p ~/.config/cheat && cheat --init > ~/.config/cheat/conf.yml

  To view the tar cheatsheet:
    cheat tar

  To edit (or create) the foo cheatsheet:
    cheat -e foo

  To edit (or create) the foo/bar cheatsheet on the "work" cheatpath:
    cheat -p work -e foo/bar

  To view all cheatsheet directories:
    cheat -d

  To list all available cheatsheets:
    cheat -l

  To briefly list all cheatsheets whose titles match "apt":
    cheat -b apt

  To list all tags in use:
    cheat -T

  To list available cheatsheets that are tagged as "personal":
    cheat -l -t personal

  To search for "ssh" among all cheatsheets, and colorize matches:
    cheat -c -s ssh

  To search (by regex) for cheatsheets that contain an IP address:
    cheat -c -r -s '(?:[0-9]{1,3}\.){3}[0-9]{1,3}'

  To remove (delete) the foo/bar cheatsheet:
    cheat --rm foo/bar

  To view the configuration file path:
    cheat --conf

  To generate shell completions (bash, zsh, fish, powershell):
    cheat --completion bash`,
	RunE:              run,
	Args:              cobra.MaximumNArgs(1),
	SilenceErrors:     true,
	SilenceUsage:      true,
	ValidArgsFunction: completions.Cheatsheets,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	f := rootCmd.Flags()

	// bool flags
	f.BoolP("all", "a", false, "Search among all cheatpaths")
	f.BoolP("brief", "b", false, "List cheatsheets without file paths")
	f.BoolP("colorize", "c", false, "Colorize output")
	f.BoolP("directories", "d", false, "List cheatsheet directories")
	f.Bool("init", false, "Write a default config file to stdout")
	f.BoolP("list", "l", false, "List cheatsheets")
	f.BoolP("regex", "r", false, "Treat search <phrase> as a regex")
	f.BoolP("tags", "T", false, "List all tags in use")
	f.BoolP("version", "v", false, "Print the version number")
	f.Bool("conf", false, "Display the config file path")

	// string flags
	f.StringP("edit", "e", "", "Edit `cheatsheet`")
	f.StringP("path", "p", "", "Return only sheets found on cheatpath `name`")
	f.StringP("search", "s", "", "Search cheatsheets for `phrase`")
	f.StringP("tag", "t", "", "Return only sheets matching `tag`")
	f.String("rm", "", "Remove (delete) `cheatsheet`")
	f.String("completion", "", "Generate shell completion script (`shell`: bash, zsh, fish, powershell)")

	// register flag completion functions
	rootCmd.RegisterFlagCompletionFunc("tag", completions.Tags)
	rootCmd.RegisterFlagCompletionFunc("path", completions.Paths)
	rootCmd.RegisterFlagCompletionFunc("edit", completions.Cheatsheets)
	rootCmd.RegisterFlagCompletionFunc("rm", completions.Cheatsheets)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	f := cmd.Flags()

	// handle --init early (no config needed)
	if initFlag, _ := f.GetBool("init"); initFlag {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get user home directory: %v\n", err)
			os.Exit(1)
		}
		envvars := config.EnvVars()
		cmdInit(home, envvars)
		os.Exit(0)
	}

	// handle --version early
	if versionFlag, _ := f.GetBool("version"); versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	// handle --completion early (no config needed)
	if f.Changed("completion") {
		shell, _ := f.GetString("completion")
		return completions.Generate(cmd, shell, os.Stdout)
	}

	// get the user's home directory
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to get user home directory: %v\n", err)
		os.Exit(1)
	}

	// read the envvars into a map of strings
	envvars := config.EnvVars()

	// identify the os-specific paths at which configs may be located
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

		// choose a confpath
		confpath = confpaths[0]

		// run the installer
		if err := installer.Run(configs(), confpath); err != nil {
			fmt.Fprintf(os.Stderr, "failed to run installer: %v\n", err)
			os.Exit(1)
		}

		// notify the user and exit
		fmt.Printf("Created config file: %s\n", confpath)
		fmt.Println("Please read this file for advanced configuration information.")
		os.Exit(0)
	}

	// initialize the configs
	conf, err := config.New(confpath, true)
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
	if f.Changed("path") {
		pathVal, _ := f.GetString("path")
		conf.Cheatpaths, err = cheatpath.Filter(
			conf.Cheatpaths,
			pathVal,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid option --path: %v\n", err)
			os.Exit(1)
		}
	}

	// determine which command to execute
	confFlag, _ := f.GetBool("conf")
	dirFlag, _ := f.GetBool("directories")
	listFlag, _ := f.GetBool("list")
	briefFlag, _ := f.GetBool("brief")
	tagsFlag, _ := f.GetBool("tags")
	tagVal, _ := f.GetString("tag")

	switch {
	case confFlag:
		cmdConf(cmd, args, conf)

	case dirFlag:
		cmdDirectories(cmd, args, conf)

	case f.Changed("edit"):
		cmdEdit(cmd, args, conf)

	case listFlag, briefFlag:
		cmdList(cmd, args, conf)

	case tagsFlag:
		cmdTags(cmd, args, conf)

	case f.Changed("search"):
		cmdSearch(cmd, args, conf)

	case f.Changed("rm"):
		cmdRemove(cmd, args, conf)

	case len(args) > 0:
		cmdView(cmd, args, conf)

	case tagVal != "":
		cmdList(cmd, args, conf)

	default:
		return cmd.Help()
	}

	return nil
}
