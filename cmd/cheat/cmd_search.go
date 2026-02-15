package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdSearch searches for strings in cheatsheets.
func cmdSearch(cmd *cobra.Command, args []string, conf config.Config) {

	phrase, _ := cmd.Flags().GetString("search")
	colorize, _ := cmd.Flags().GetBool("colorize")
	useRegex, _ := cmd.Flags().GetBool("regex")

	// load the cheatsheets
	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to list cheatsheets: %v\n", err)
		os.Exit(1)
	}
	if cmd.Flags().Changed("tag") {
		tagVal, _ := cmd.Flags().GetString("tag")
		cheatsheets = sheets.Filter(
			cheatsheets,
			strings.Split(tagVal, ","),
		)
	}

	// prepare the search pattern
	pattern := "(?i)" + phrase

	// unless --regex is provided, in which case we pass the regex unaltered
	if useRegex {
		pattern = phrase
	}

	// compile the regex once, outside the loop
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to compile regexp: %s, %v\n", pattern, err)
		os.Exit(1)
	}

	// iterate over each cheatpath
	out := ""
	for _, pathcheats := range cheatsheets {

		// sort the cheatsheets alphabetically, and search for matches
		for _, sheet := range sheets.Sort(pathcheats) {

			// if <cheatsheet> was provided, constrain the search only to
			// matching cheatsheets
			if len(args) > 0 && sheet.Title != args[0] {
				continue
			}

			// `Search` will return text entries that match the search terms.
			// We're using it here to overwrite the prior cheatsheet Text,
			// filtering it to only what is relevant.
			sheet.Text = sheet.Search(reg)

			// if the sheet did not match the search, ignore it and move on
			if sheet.Text == "" {
				continue
			}

			// if colorization was requested, apply it here
			if conf.Color(colorize) {
				sheet.Colorize(conf)
			}

			// display the cheatsheet body
			out += fmt.Sprintf(
				"%s %s\n%s\n",
				// append the cheatsheet title
				sheet.Title,
				// append the cheatsheet path
				display.Faint(fmt.Sprintf("(%s)", sheet.CheatPath), conf.Color(colorize)),
				// indent each line of content
				display.Indent(sheet.Text),
			)
		}
	}

	// trim superfluous newlines
	out = strings.TrimSpace(out)

	// display the output
	// NB: resist the temptation to call `display.Write` multiple times in the
	// loop above. That will not play nicely with the paginator.
	display.Write(out, conf)
}
