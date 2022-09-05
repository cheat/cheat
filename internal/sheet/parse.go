package sheet

import (
	"fmt"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse parses cheatsheet frontmatter
func parse(markdown string) (frontmatter, string, error) {

	// determine the appropriate line-break for the platform
	linebreak := "\n"
	if runtime.GOOS == "windows" {
		linebreak = "\r\n"
	}

	// specify the frontmatter delimiter
	delim := fmt.Sprintf("---%s", linebreak)

	// initialize a frontmatter struct
	var fm frontmatter

	// if the markdown does not contain frontmatter, pass it through unmodified
	if !strings.HasPrefix(markdown, delim) {
		return fm, markdown, nil
	}

	// otherwise, split the frontmatter and cheatsheet text
	parts := strings.SplitN(markdown, delim, 3)

	// return an error if the frontmatter parses into the wrong number of parts
	if len(parts) != 3 {
		return fm, markdown, fmt.Errorf("failed to delimit frontmatter")
	}

	// return an error if the YAML cannot be unmarshalled
	if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
		return fm, markdown, fmt.Errorf("failed to unmarshal frontmatter: %v", err)
	}

	return fm, parts[2], nil
}
