package frontmatter

import (
	"fmt"
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
