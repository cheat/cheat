package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/cheat/cheat/internal/config"
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

	// exit early if no cheatsheets are available
	if len(flattened) == 0 {
		os.Exit(0)
	}

	// initialize a tabwriter to produce cleanly columnized output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	// generate sorted, columnized output
	fmt.Fprintln(w, "title:\tfile:\ttags:")
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
}
