package sheet

import (
	"regexp"
	"strings"
)

// Search searches for regexp matches in a cheatsheet's text, and optionally
// colorizes matching strings.
func (s *Sheet) Search(reg *regexp.Regexp) []Match {

	// record matches
	matches := []Match{}

	// search through the cheatsheet's text line by line
	// TODO: searching line-by-line is surely the "naive" approach. Revisit this
	// later with an eye for performance improvements.
	for _, line := range strings.Split(s.Text, "\n") {

		// exit early if the line doesn't match the regex
		if !reg.MatchString(line) {
			continue
		}

		// init the match
		m := Match{
			Text: strings.TrimSpace(line),
		}

		// record the match
		matches = append(matches, m)
	}

	return matches
}
