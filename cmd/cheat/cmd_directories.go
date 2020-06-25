package main

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
)

// cmdDirectories lists the configured cheatpaths.
func cmdDirectories(opts map[string]interface{}, conf config.Config) {

	// initialize a tabwriter to produce cleanly columnized output
	var out bytes.Buffer
	w := tabwriter.NewWriter(&out, 0, 0, 1, ' ', 0)

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
	display.Display(out.String(), conf)
}
