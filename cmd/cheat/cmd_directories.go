package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/cheat/cheat/internal/config"
)

// cmdDirectories lists the configured cheatpaths.
func cmdDirectories(opts map[string]interface{}, conf config.Config) {

	// initialize a tabwriter to produce cleanly columnized output
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)

	// generate sorted, columnized output
	for _, path := range conf.Cheatpaths {
		fmt.Fprintln(w, fmt.Sprintf(
			"%s:\t%s",
			path.Name,
			path.Path,
		))
	}

	// write columnized output to stdout
	w.Flush()
}
