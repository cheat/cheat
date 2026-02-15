// Package display implement functions pertaining to writing formatted
// cheatsheet content to stdout, or alternatively the system pager.
package display

import "fmt"

// Faint returns a faintly-colored string that's used to de-prioritize text
// written to stdout
func Faint(str string, colorize bool) string {
	// make `str` faint only if colorization has been requested
	if colorize {
		return fmt.Sprintf("\033[2m%s\033[0m", str)
	}

	// otherwise, return the string unmodified
	return str
}
