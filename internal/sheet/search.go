package sheet

import (
	"regexp"
	"strings"
)

// Search returns lines within a sheet's Text that match the search regex
func (s *Sheet) Search(reg *regexp.Regexp) string {

	// record matches
	matches := ""

	// search through the cheatsheet's text line by line
	for _, line := range strings.Split(s.Text, "\n\n") {

		// exit early if the line doesn't match the regex
		if reg.MatchString(line) {
			matches += line + "\n\n"
		}
	}

	return strings.TrimSpace(matches)
}
