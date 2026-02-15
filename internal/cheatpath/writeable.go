package cheatpath

import (
	"fmt"
)

// Writeable returns a writeable Path
func Writeable(cheatpaths []Path) (Path, error) {

	// iterate backwards over the cheatpaths
	// NB: we're going backwards because we assume that the most "local"
	// cheatpath will be specified last in the configs
	for i := len(cheatpaths) - 1; i >= 0; i-- {
		// if the cheatpath is not read-only, it is writeable, and thus returned
		if !cheatpaths[i].ReadOnly {
			return cheatpaths[i], nil
		}
	}

	// otherwise, return an error
	return Path{}, fmt.Errorf("no writeable cheatpaths found")
}
