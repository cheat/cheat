package cheatpath

import (
	"fmt"
)

// Filter filters all cheatpaths that are not named `name`
func Filter(paths []Cheatpath, name string) ([]Cheatpath, error) {

	// if a path of the given name exists, return it
	for _, path := range paths {
		if path.Name == name {
			return []Cheatpath{path}, nil
		}
	}

	// otherwise, return an error
	return []Cheatpath{}, fmt.Errorf("cheatpath does not exist: %s", name)
}
