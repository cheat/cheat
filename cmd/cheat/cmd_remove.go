package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdRemove removes (deletes) a cheatsheet.
func cmdRemove(opts map[string]interface{}, conf config.Config) {

	cheatsheet := opts["--rm"].(string)

	// validate the cheatsheet name
	if err := cheatpath.ValidateSheetName(cheatsheet); err != nil {
		fmt.Fprintf(os.Stderr, "invalid cheatsheet name: %v\n", err)
		os.Exit(1)
	}

	// load the cheatsheets
	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to list cheatsheets: %v\n", err)
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
		fmt.Fprintf(os.Stderr, "No cheatsheet found for '%s'.\n", cheatsheet)
		os.Exit(2)
	}

	// fail early if the sheet is read-only
	if sheet.ReadOnly {
		fmt.Fprintf(os.Stderr, "cheatsheet '%s' is read-only.\n", cheatsheet)
		os.Exit(1)
	}

	// otherwise, attempt to delete the sheet
	if err := os.Remove(sheet.Path); err != nil {
		fmt.Fprintf(os.Stderr, "failed to delete sheet: %s, %v\n", sheet.Title, err)
		os.Exit(1)
	}
}
