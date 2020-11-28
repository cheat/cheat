package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
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

	// iterate over each cheatpath
	out := ""
	for _, pathcheats := range cheatsheets {

		// sort the cheatsheets alphabetically, and search for matches
		for _, sheet := range sheets.Sort(pathcheats) {

			// if <cheatsheet> was provided, constrain the search only to
			// matching cheatsheets
			if opts["<cheatsheet>"] != nil && sheet.Title != opts["<cheatsheet>"] {
				continue
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

			// display the cheatsheet title and path
			out += fmt.Sprintf("%s %s\n",
				display.Underline(sheet.Title),
				display.Faint(fmt.Sprintf("(%s)", sheet.CheatPath), conf),
			)

			// indent each line of content
			out += display.Indent(sheet.Text) + "\n"
		}
	}

	// trim superfluous newlines
	out = strings.TrimSpace(out)

	// display the output
	// NB: resist the temptation to call `display.Display` multiple times in
	// the loop above. That will not play nicely with the paginator.
	display.Write(out, conf)
}
