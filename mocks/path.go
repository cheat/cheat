// Package mocks provides test fixture data and helpers for unit tests.
package mocks

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Path returns the absolute path to the specified mock file within
// the mocks/ directory.
func Path(filename string) string {
	_, thisfile, _, _ := runtime.Caller(0)

	file, err := filepath.Abs(
		filepath.Join(filepath.Dir(thisfile), filename),
	)
	if err != nil {
		panic(fmt.Errorf("failed to resolve mock path: %v", err))
	}

	return file
}
