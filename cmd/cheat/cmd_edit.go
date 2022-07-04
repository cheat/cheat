package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/sheets"
)

// cmdEdit opens a cheatsheet for editing (or creates it if it doesn't exist).
func cmdEdit(opts map[string]interface{}, conf config.Config) {

	cheatsheet := opts["--edit"].(string)

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

	// consolidate the cheatsheets found on all paths into a single map of
	// `title` => `sheet` (ie, allow more local cheatsheets to override less
	// local cheatsheets)
	consolidated := sheets.Consolidate(cheatsheets)

	// the file path of the sheet to edit
	var editpath string

	// determine if the sheet exists
	sheet, ok := consolidated[cheatsheet]

	// if the sheet exists and is not read-only, edit it in place
	if ok && !sheet.ReadOnly {
		editpath = sheet.Path

		// if the sheet exists but is read-only, copy it before editing
	} else if ok && sheet.ReadOnly {
		// compute the new edit path
		// begin by getting a writeable cheatpath
		writepath, err := cheatpath.Writeable(conf.Cheatpaths)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get writeable path: %v\n", err)
			os.Exit(1)
		}

		// compute the new edit path
		editpath = filepath.Join(writepath.Path, sheet.Title)

		// create any necessary subdirectories
		dirs := filepath.Dir(editpath)
		if dirs != "." {
			if err := os.MkdirAll(dirs, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "failed to create directory: %s, %v\n", dirs, err)
				os.Exit(1)
			}
		}

		// copy the sheet to the new edit path
		err = sheet.Copy(editpath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to copy cheatsheet: %v\n", err)
			os.Exit(1)
		}

		// if the sheet does not exist, create it
	} else {
		// compute the new edit path
		// begin by getting a writeable cheatpath
		writepath, err := cheatpath.Writeable(conf.Cheatpaths)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to get writeable path: %v\n", err)
			os.Exit(1)
		}

		// compute the new edit path
		editpath = filepath.Join(writepath.Path, cheatsheet)

		// create any necessary subdirectories
		dirs := filepath.Dir(editpath)
		if dirs != "." {
			if err := os.MkdirAll(dirs, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "failed to create directory: %s, %v\n", dirs, err)
				os.Exit(1)
			}
		}
	}

	// split `conf.Editor` into parts to separate the editor's executable from
	// any arguments it may have been passed. If this is not done, the nearby
	// call to `exec.Command` will fail.
	parts := strings.Fields(conf.Editor)
	editor := parts[0]
	args := append(parts[1:], editpath)

	// edit the cheatsheet
	cmd := exec.Command(editor, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to edit cheatsheet: %v\n", err)
		os.Exit(1)
	}
}
