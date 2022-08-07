package config

import (
	"os"
	"os/exec"
)

// Pager attempts to locate a pager that's appropriate for the environment.
func Pager() string {

	// if $PAGER is set, return the corresponding pager
	if os.Getenv("PAGER") != "" {
		return os.Getenv("PAGER")
	}

	// Otherwise, search for `pager`, `less`, and `more` on the `$PATH`. If
	// none are found, return an empty pager.
	for _, pager := range []string{"pager", "less", "more"} {
		if path, err := exec.LookPath(pager); err != nil {
			return path
		}
	}

	// default to no pager
	return ""
}
