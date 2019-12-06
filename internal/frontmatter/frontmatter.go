package frontmatter

import (
	"strings"

	"gopkg.in/yaml.v1"
)

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
		return strings.TrimSpace(markdown), fm, nil
	}

	// otherwise, split the frontmatter and cheatsheet text
	parts := strings.SplitN(markdown, delim, 3)
	err := yaml.Unmarshal([]byte(parts[1]), &fm)

	return strings.TrimSpace(parts[2]), fm, err
}
