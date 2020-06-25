package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
	"github.com/cheat/cheat/internal/sheet"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdList lists all available cheatsheets.
func cmdList(opts map[string]interface{}, conf config.Config) {

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

	// instead of "consolidating" all of the cheatsheets (ie, overwriting global
	// sheets with local sheets), here we simply want to create a slice
	// containing all sheets.
	flattened := []sheet.Sheet{}
	for _, pathSheets := range cheatsheets {
		for _, s := range pathSheets {
			flattened = append(flattened, s)
		}
	}

	// sort the "flattened" sheets alphabetically
	sort.Slice(flattened, func(i, j int) bool {
		return flattened[i].Title < flattened[j].Title
	})

	// filter if <cheatsheet> was specified
	// NB: our docopt specification is misleading here. When used in conjunction
	// with `-l`, `<cheatsheet>` is really a pattern against which to filter
	// sheet titles.
	if opts["<cheatsheet>"] != nil {

		// initialize a slice of filtered sheets
		filtered := []sheet.Sheet{}

		// initialize our filter pattern
		pattern := "(?i)" + opts["<cheatsheet>"].(string)

		// compile the regex
		reg, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintln(
				os.Stderr,
				fmt.Sprintf("failed to compile regexp: %s, %v", pattern, err),
			)
			os.Exit(1)
		}

		// iterate over each cheatsheet, and pass-through those which match the
		// filter pattern
		for _, s := range flattened {
			if reg.MatchString(s.Title) {
				filtered = append(filtered, s)
			}
		}

		flattened = filtered
	}

	// exit early if no cheatsheets are available
	if len(flattened) == 0 {
		os.Exit(0)
	}

	// initialize a tabwriter to produce cleanly columnized output
	var out bytes.Buffer
	w := tabwriter.NewWriter(&out, 0, 0, 1, ' ', 0)

	// write a header row
	fmt.Fprintln(w, "title:\tfile:\ttags:")

	// generate sorted, columnized output
	for _, sheet := range flattened {
		fmt.Fprintln(w, fmt.Sprintf(
			"%s\t%s\t%s",
			sheet.Title,
			sheet.Path,
			strings.Join(sheet.Tags, ","),
		))
	}

	// write columnized output to stdout
	w.Flush()
	display.Display(out.String(), conf)
}
