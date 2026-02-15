package cheatpath

import (
	"fmt"
)

// Filter filters all cheatpaths that are not named `name`
func Filter(paths []Path, name string) ([]Path, error) {

	// if a path of the given name exists, return it
	for _, path := range paths {
		if path.Name == name {
			return []Path{path}, nil
		}
	}

	// otherwise, return an error
	return []Path{}, fmt.Errorf("cheatpath does not exist: %s", name)
}
