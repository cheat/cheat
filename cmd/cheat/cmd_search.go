package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdSearch searches for strings in cheatsheets.
func cmdSearch(opts map[string]interface{}, conf config.Config) {

	phrase := opts["--search"].(string)

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

	// sort the cheatsheets alphabetically, and search for matches
	for _, sheet := range sheets.Sort(consolidated) {

		// colorize output?
		colorize := false
		if conf.Colorize == true || opts["--colorize"] == true {
			colorize = true
		}

		// assume that we want to perform a case-insensitive search for <phrase>
		pattern := "(?i)" + phrase

		// unless --regex is provided, in which case we pass the regex unaltered
		if opts["--regex"] == true {
			pattern = phrase
		}

		// compile the regex
		reg, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Errorf("failed to compile regexp: %s, %v", pattern, err)
			os.Exit(1)
		}

		// search the sheet
		matches := sheet.Search(reg, colorize)

		// display the results
		if len(matches) > 0 {
			fmt.Printf("%s:\n", sheet.Title)
			for _, m := range matches {
				fmt.Printf("  %d: %s\n", m.Line, m.Text)
			}
			fmt.Print("\n")
		}
	}

}
