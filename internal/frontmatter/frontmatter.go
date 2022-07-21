package frontmatter

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"gopkg.in/yaml.v1"
)

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// Frontmatter encapsulates cheatsheet frontmatter data
type Frontmatter struct {
	Tags   []string
	Syntax string
}

// Parse parses cheatsheet frontmatter
func Parse(markdown string) (string, Frontmatter, error) {

	// specify the frontmatter delimiter
	delim := "---"

	// initialize a frontmatter struct
	var fm Frontmatter

	// if the markdown does not contain frontmatter, pass it through unmodified
	if !strings.HasPrefix(markdown, delim) {
		return TrimEmptyLines(markdown), fm, nil
	}

	// otherwise, split the frontmatter and cheatsheet text
	parts := strings.SplitN(markdown, delim, 3)

	// return an error if the frontmatter parses into the wrong number of parts
	if len(parts) != 3 {
		return markdown, fm, fmt.Errorf("failed to delimit frontmatter")
	}

	// return an error if the YAML cannot be unmarshalled
	if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
		return markdown, fm, fmt.Errorf("failed to unmarshal frontmatter: %v", err)
	}

	return strings.TrimSpace(parts[2]), fm, nil
}

// TrimEmptyLines returns a slice of the string s, with all leading
// and trailing lines removed, which consist only of white space as
// defined by Unicode.
func TrimEmptyLines(s string) string {
	start := 0
	newlinePos := -1
	for ; start < len(s); start++ {
		c := s[start]
		if c >= utf8.RuneSelf {
			return strings.TrimFunc(s[start:], unicode.IsSpace)
		}
		if c == '\n' {
			newlinePos = start
			continue
		}
		if asciiSpace[c] == 0 {
			if newlinePos >= 0 {
				start = newlinePos + 1
			} else {
				start = 0
			}
			break
		}
	}

	stop := len(s)
	newlinePos = -1
	for ; stop > start; stop-- {
		c := s[stop-1]
		if c >= utf8.RuneSelf {
			return strings.TrimFunc(s[start:stop], unicode.IsSpace)
		}
		if c == '\n' {
			newlinePos = stop
			continue
		}
		if asciiSpace[c] == 0 {
			if newlinePos >= 0 {
				stop = newlinePos
			} else {
				stop = len(s)
			}
			break
		}
	}

	return s[start:stop]
}
