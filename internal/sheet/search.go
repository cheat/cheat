package sheet

import (
	"regexp"
	"strings"

	"github.com/mgutz/ansi"
)

// Search searches for regexp matches in a cheatsheet's text, and optionally
// colorizes matching strings.
func (s *Sheet) Search(reg *regexp.Regexp, colorize bool) []Match {

	// record matches
	matches := []Match{}

	// search through the cheatsheet's text line by line
	// TODO: searching line-by-line is surely the "naive" approach. Revisit this
	// later with an eye for performance improvements.
	for linenum, line := range strings.Split(s.Text, "\n") {

		// exit early if the line doesn't match the regex
		if !reg.MatchString(line) {
			continue
		}

		// init the match
		m := Match{
			Line: linenum + 1,
			Text: strings.TrimSpace(line),
		}

		// colorize the matching text if so configured
		if colorize {
			m.Text = reg.ReplaceAllStringFunc(m.Text, func(matched string) string {
				return ansi.Color(matched, "red+b")
			})
		}

		// record the match
		matches = append(matches, m)
	}

	return matches
}
