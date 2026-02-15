package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
	"github.com/cheat/cheat/internal/sheet"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdList lists all available cheatsheets.
func cmdList(cmd *cobra.Command, args []string, conf config.Config) {

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

	// instead of "consolidating" all of the cheatsheets (ie, overwriting global
	// sheets with local sheets), here we simply want to create a slice
	// containing all sheets.
	flattened := []sheet.Sheet{}
	for _, pathsheets := range cheatsheets {
		for _, s := range pathsheets {
			flattened = append(flattened, s)
		}
	}

	// sort the "flattened" sheets alphabetically
	sort.Slice(flattened, func(i, j int) bool {
		return flattened[i].Title < flattened[j].Title
	})

	// filter if <cheatsheet> was specified
	if len(args) > 0 {

		// initialize a slice of filtered sheets
		filtered := []sheet.Sheet{}

		// initialize our filter pattern
		pattern := "(?i)" + args[0]

		// compile the regex
		reg, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to compile regexp: %s, %v\n", pattern, err)
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

	// return exit code 2 if no cheatsheets are available
	if len(flattened) == 0 {
		os.Exit(2)
	}

	// initialize a tabwriter to produce cleanly columnized output
	var out bytes.Buffer
	w := tabwriter.NewWriter(&out, 0, 0, 1, ' ', 0)

	// generate sorted, columnized output
	briefFlag, _ := cmd.Flags().GetBool("brief")
	if briefFlag {
		fmt.Fprintln(w, "title:\ttags:")
		for _, sheet := range flattened {
			fmt.Fprintf(w, "%s\t%s\n", sheet.Title, strings.Join(sheet.Tags, ","))
		}
	} else {
		fmt.Fprintln(w, "title:\tfile:\ttags:")
		for _, sheet := range flattened {
			fmt.Fprintf(w, "%s\t%s\t%s\n", sheet.Title, sheet.Path, strings.Join(sheet.Tags, ","))
		}
	}

	// write columnized output to stdout
	w.Flush()
	display.Write(out.String(), conf)
}
