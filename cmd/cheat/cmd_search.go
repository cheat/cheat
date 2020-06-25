package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
	"github.com/cheat/cheat/internal/sheet"
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

	// if <cheatsheet> was provided, search that single sheet only
	if opts["<cheatsheet>"] != nil {

		cheatsheet := opts["<cheatsheet>"].(string)

		// assert that the cheatsheet exists
		s, ok := consolidated[cheatsheet]
		if !ok {
			fmt.Printf("No cheatsheet found for '%s'.\n", cheatsheet)
			os.Exit(2)
		}

		consolidated = map[string]sheet.Sheet{
			cheatsheet: s,
		}
	}

	// sort the cheatsheets alphabetically, and search for matches
	for _, sheet := range sheets.Sort(consolidated) {

		// assume that we want to perform a case-insensitive search for <phrase>
		pattern := "(?i)" + phrase

		// unless --regex is provided, in which case we pass the regex unaltered
		if opts["--regex"] == true {
			pattern = phrase
		}

		// compile the regex
		reg, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("failed to compile regexp: %s, %v", pattern, err))
			os.Exit(1)
		}

		// `Search` will return text entries that match the search terms. We're
		// using it here to overwrite the prior cheatsheet Text, filtering it to
		// only what is relevant
		sheet.Text = sheet.Search(reg)

		// if the sheet did not match the search, ignore it and move on
		if sheet.Text == "" {
			continue
		}

		// if colorization was requested, apply it here
		if conf.Color(opts) {
			sheet.Colorize(conf)
		}

		// output the cheatsheet title
		out := fmt.Sprintf("%s:\n", sheet.Title)

		// indent each line of content with two spaces
		for _, line := range strings.Split(sheet.Text, "\n") {
			out += fmt.Sprintf("  %s\n", line)
		}

		// display the output
		display.Display(out, conf)
	}
}
