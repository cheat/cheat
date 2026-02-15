// Package installer implements functions that provide a first-time
// installation wizard.
package installer

import (
	"fmt"
	"os"
	"strings"
)

// Prompt prompts the user for a answer
func Prompt(prompt string, def bool) (bool, error) {

	// display the prompt
	fmt.Printf("%s: ", prompt)

	// read one byte at a time until newline to avoid buffering past the
	// end of the current line, which would consume input intended for
	// subsequent Prompt calls on the same stdin
	var line []byte
	buf := make([]byte, 1)
	for {
		n, err := os.Stdin.Read(buf)
		if n > 0 {
			if buf[0] == '\n' {
				break
			}
			if buf[0] != '\r' {
				line = append(line, buf[0])
			}
		}
		if err != nil {
			if len(line) > 0 {
				break
			}
			return false, fmt.Errorf("failed to prompt: %v", err)
		}
	}

	// normalize the answer
	ans := strings.ToLower(strings.TrimSpace(string(line)))

	// return the appropriate response
	switch ans {
	case "y":
		return true, nil
	case "":
		return def, nil
	default:
		return false, nil
	}
}
