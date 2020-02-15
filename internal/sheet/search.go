package sheet

import (
	"regexp"
	"strings"
)

// Search returns lines within a sheet's Text that match the search regex
func (s *Sheet) Search(reg *regexp.Regexp) string {

	// record matches
	matches := []string{}

	// search through the cheatsheet's text line by line
	// TODO: searching line-by-line is surely the "naive" approach. Revisit this
	// later with an eye for performance improvements.
	for _, line := range strings.Split(s.Text, "\n") {

		// exit early if the line doesn't match the regex
		if !reg.MatchString(line) {
			continue
		}

		// record the match
		matches = append(matches, line)
	}

	return strings.Join(matches, "\n")
}
