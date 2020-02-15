package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cheat/cheat/internal/config"
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
		fmt.Printf("No cheatsheet found for '%s'.\n", cheatsheet)
		os.Exit(0)
	}

	// apply colorization if requested
	if conf.Color(opts) {
		sheet.Colorize(conf)
	}

	// display the cheatsheet
	fmt.Print(sheet.Text)
}
