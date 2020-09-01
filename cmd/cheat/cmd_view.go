package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdView displays a cheatsheet for viewing.
func cmdView(opts map[string]interface{}, conf config.Config) {

	cheatsheet := opts["<cheatsheet>"].(string)

	// load the cheatsheets
	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("failed to list cheatsheets: %v", err))
		os.Exit(1)
	}

	// filter cheatcheats by tag if --tag was provided
	if opts["--tag"] != nil {
		cheatsheets = sheets.Filter(
			cheatsheets,
			strings.Split(opts["--tag"].(string), ","),
		)
	}

	// consolidate the cheatsheets found on all paths into a single map of
	// `title` => `sheet` (ie, allow more local cheatsheets to override less
	// local cheatsheets)
	consolidated := sheets.Consolidate(cheatsheets)

	// fail early if the requested cheatsheet does not exist
	sheet, ok := consolidated[cheatsheet]
	if !ok {
		if opts["--tag"] != nil {
			fmt.Printf("No cheatsheet found for '%s'.\n", cheatsheet)
			os.Exit(2)
		} else {
			path := opts["<cheatsheet>"]
			opts["--tag"], opts["<cheatsheet>"] = splitTagAndPath(path.(string))
			if opts["<cheatsheet>"] == "" || opts["--tag"] == "" {
				fmt.Printf("No cheatsheet found for '%s'.\n", cheatsheet)
				os.Exit(2)
			}
			cmdView(opts, conf)
			return
		}
	}

	// apply colorization if requested
	if conf.Color(opts) {
		sheet.Colorize(conf)
	}

	// display the cheatsheet
	display.Display(sheet.Text, conf)
}

func splitTagAndPath(path string) (string, string) {
	s := strings.Split(path, string(filepath.Separator))
	if len(s) > 0 {
		return s[0], filepath.Join(s[1:]...)
	}
	return "", ""
}
