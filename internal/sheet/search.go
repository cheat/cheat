package sheet

import (
	"regexp"
	"strings"
)

// Search returns lines within a sheet's Text that match the search regex
func (s *Sheet) Search(reg *regexp.Regexp) string {

	// record matches
	var matches []string

	// search through the cheatsheet's text line by line
	for _, line := range strings.Split(s.Text, "\n\n") {

		// save matching lines
		if reg.MatchString(line) {
			matches = append(matches, line)
		}
	}

	// Join matches with the same delimiter used for splitting
	return strings.Join(matches, "\n\n")
}
