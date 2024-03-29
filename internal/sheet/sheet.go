// Package sheet implements functions pertaining to parsing, searching, and
// displaying cheatsheets.
package sheet

import (
	"fmt"
	"os"
	"sort"
)

// Frontmatter encapsulates cheatsheet frontmatter data
type frontmatter struct {
	Tags   []string
	Syntax string
}

// Sheet encapsulates sheet information
type Sheet struct {
	Title     string
	CheatPath string
	Path      string
	Text      string
	Tags      []string
	Syntax    string
	ReadOnly  bool
}

// New initializes a new Sheet
func New(
	title string,
	cheatpath string,
	path string,
	tags []string,
	readOnly bool,
) (Sheet, error) {

	// read the cheatsheet file
	markdown, err := os.ReadFile(path)
	if err != nil {
		return Sheet{}, fmt.Errorf("failed to read file: %s, %v", path, err)
	}

	// parse the raw cheatsheet text
	fm, text, err := parse(string(markdown))
	if err != nil {
		return Sheet{}, fmt.Errorf("failed to parse front-matter: %v", err)
	}

	// merge the sheet-specific tags into the cheatpath tags
	tags = append(tags, fm.Tags...)

	// sort strings so they pretty-print nicely
	sort.Strings(tags)

	// initialize and return a sheet
	return Sheet{
		Title:     title,
		CheatPath: cheatpath,
		Path:      path,
		Text:      text,
		Tags:      tags,
		Syntax:    fm.Syntax,
		ReadOnly:  readOnly,
	}, nil
}
