// Package front provides YAML frontmatter unmarshalling.
package front

import (
	"bytes"

	"gopkg.in/yaml.v1"
)

// Delimiter.
var delim = []byte("---")

// Unmarshal parses YAML frontmatter and returns the content. When no
// frontmatter delimiters are present the original content is returned.
func Unmarshal(b []byte, v interface{}) (content []byte, err error) {
	if !bytes.HasPrefix(b, delim) {
		return b, nil
	}

	parts := bytes.SplitN(b, delim, 3)
	content = parts[2]
	err = yaml.Unmarshal(parts[1], v)
	return
}
