package cheatpath

import (
	"fmt"
)

// Writeable returns a writeable Cheatpath
func Writeable(cheatpaths []Cheatpath) (Cheatpath, error) {

	// iterate backwards over the cheatpaths
	// NB: we're going backwards because we assume that the most "local"
	// cheatpath will be specified last in the configs
	for i := len(cheatpaths) - 1; i >= 0; i-- {

		// if the cheatpath is not read-only, it is writeable, and thus returned
		if cheatpaths[i].ReadOnly == false {
			return cheatpaths[i], nil
		}

	}

	// otherwise, return an error
	return Cheatpath{}, fmt.Errorf("no writeable cheatpaths found")
}
