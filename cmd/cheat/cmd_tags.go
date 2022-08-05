package main

import (
	"fmt"
	"os"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/display"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdTags lists all tags in use.
func cmdTags(opts map[string]interface{}, conf config.Config) {

	// load the cheatsheets
	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to list cheatsheets: %v\n", err)
		os.Exit(1)
	}

	// assemble the output
	out := ""
	for _, tag := range sheets.Tags(cheatsheets) {
		out += fmt.Sprintln(tag)
	}

	// display the output
	display.Write(out, conf)
}
