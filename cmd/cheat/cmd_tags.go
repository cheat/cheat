package main

import (
	"fmt"
	"os"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdTags lists all tags in use.
func cmdTags(opts map[string]interface{}, conf config.Config) {

	// load the cheatsheets
	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("failed to list cheatsheets: %v", err))
		os.Exit(1)
	}

	// write sheet tags to stdout
	for _, tag := range sheets.Tags(cheatsheets) {
		fmt.Println(tag)
	}
}
