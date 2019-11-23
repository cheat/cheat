package mock

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Path returns the absolute path to the specified mock file.
func Path(filename string) string {
	// determine the path of this file during runtime
	_, thisfile, _, _ := runtime.Caller(0)

	breadcrumbs := []string{
		filepath.Dir(thisfile),
		filepath.Clean("../../mocks"),
		filepath.Clean(filename),
	}

	// compute the config path
	combinedBreadcrumbs := strings.Join(breadcrumbs, string(os.PathSeparator))
	file, err := filepath.Abs(combinedBreadcrumbs)
	if err != nil {
		panic(fmt.Errorf("failed to resolve config path: %v", err))
	}

	return file
}
