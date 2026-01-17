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
		fmt.Fprintf(os.Stderr, "failed to load cheatsheets from configured paths: %v\n", err)
		os.Exit(1)
	}

	// filter cheatcheats by tag if --tag was provided
	if opts["--tag"] != nil {
		cheatsheets = sheets.Filter(
			cheatsheets,
			strings.Split(opts["--tag"].(string), ","),
		)
		if len(cheatsheets) == 0 {
			fmt.Fprintf(os.Stderr, "no cheatsheets found matching the specified tags\n")
			os.Exit(1)
		}
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
			fmt.Fprintf(os.Stderr, "failed to find a writable cheatpath for editing: %v\n", err)
			os.Exit(1)
		}

		// compute the new edit path
		editpath = filepath.Join(writepath.Path, sheet.Title)

		// create any necessary subdirectories
		dirs := filepath.Dir(editpath)
		if dirs != "." {
			if err := os.MkdirAll(dirs, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "failed to create directory structure for copied cheatsheet: %s, error: %v\n", dirs, err)
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
			fmt.Fprintf(os.Stderr, "failed to find a writable cheatpath for new cheatsheet: %v\n", err)
			os.Exit(1)
		}

		// compute the new edit path
		editpath = filepath.Join(writepath.Path, cheatsheet)

		// create any necessary subdirectories
		dirs := filepath.Dir(editpath)
		if dirs != "." {
			if err := os.MkdirAll(dirs, 0755); err != nil {
				fmt.Fprintf(os.Stderr, "failed to create directory structure for new cheatsheet: %s, error: %v\n", dirs, err)
				os.Exit(1)
			}
		}
	}

	// split `conf.Editor` into parts to separate the editor's executable from
	// any arguments it may have been passed. If this is not done, the nearby
	// call to `exec.Command` will fail.
	editor, args := parseEditorCommand(conf.Editor, editpath)

	// edit the cheatsheet
	cmd := exec.Command(editor, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute editor '%s': %v\n", editor, err)
		os.Exit(1)
	}
}

// parseEditorCommand parses the editor command string into the executable and arguments
// This handles properly quoted paths and arguments that may contain spaces
func parseEditorCommand(editorCmd string, editpath string) (string, []string) {
	// Handle the case where editorCmd is empty
	if editorCmd == "" {
		return "", []string{editpath}
	}

	var args []string
	var currentArg string
	var inQuotes bool
	var quoteChar rune

	for _, r := range editorCmd {
		switch {
		case r == '"' || r == '\'':
			if inQuotes && r == quoteChar {
				inQuotes = false
			} else if !inQuotes {
				inQuotes = true
				quoteChar = r
			} else {
				currentArg += string(r)
			}
		case r == ' ' && !inQuotes:
			if currentArg != "" {
				args = append(args, currentArg)
				currentArg = ""
			}
		default:
			currentArg += string(r)
		}
	}

	// Add the last argument if it exists
	if currentArg != "" {
		args = append(args, currentArg)
	}

	// If no arguments were parsed, use the entire string as the editor
	editor := args[0]
	args = append(args[1:], editpath)

	return editor, args
}
