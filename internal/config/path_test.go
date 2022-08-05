package config

import (
	"os"
	"testing"
)

// TestPathConfigNotExists asserts that `Path` identifies non-existent config
// files
func TestPathConfigNotExists(t *testing.T) {

	// package (invalid) cheatpaths
	paths := []string{"/cheat-test-conf-does-not-exist"}

	// assert
	if _, err := Path(paths); err == nil {
		t.Errorf("failed to identify non-existent config file")
	}

}

// TestPathConfigExists asserts that `Path` identifies existent config files
func TestPathConfigExists(t *testing.T) {

	// initialize a temporary config file
	confFile, err := os.CreateTemp("", "cheat-test")
	if err != nil {
		t.Errorf("failed to create temp file: %v", err)
	}

	// clean up the temp file
	defer os.Remove(confFile.Name())

	// package cheatpaths
	paths := []string{
		"/cheat-test-conf-does-not-exist",
		confFile.Name(),
	}

	// assert
	got, err := Path(paths)
	if err != nil {
		t.Errorf("failed to identify config file: %v", err)
	}
	if got != confFile.Name() {
		t.Errorf(
			"failed to return config path: want: %s, got: %s",
			confFile.Name(),
			got,
		)
	}
}
