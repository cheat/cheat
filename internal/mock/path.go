package mock

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

// Path returns the absolute path to the specified mock file.
func Path(filename string) string {

	// determine the path of this file during runtime
	_, thisfile, _, _ := runtime.Caller(0)

	// compute the config path
	file, err := filepath.Abs(
		path.Join(
			filepath.Dir(thisfile),
			"../../mocks",
			filename,
		),
	)
	if err != nil {
		panic(fmt.Errorf("failed to resolve config path: %v", err))
	}

	return file
}
