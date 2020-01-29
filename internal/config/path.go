package config

import (
	"fmt"
	"os"
)

// Path returns the config file path
func Path(paths []string) (string, error) {

	// check if the config file exists on any paths
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	// we can't find the config file if we make it this far
	return "", fmt.Errorf("could not locate config file")
}
