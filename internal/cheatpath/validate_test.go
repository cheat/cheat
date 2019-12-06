package cheatpath

import (
	"testing"
)

// TestValidateValid asserts that valid cheatpaths validate successfully
func TestValidateValid(t *testing.T) {

	// initialize a valid cheatpath
	cheatpath := Cheatpath{
		Name:     "foo",
		Path:     "/foo",
		ReadOnly: false,
		Tags:     []string{},
	}

	// assert that no errors are returned
	if err := cheatpath.Validate(); err != nil {
		t.Errorf("failed to validate valid cheatpath: %v", err)
	}
}

// TestValidateMissingName asserts that paths that are missing a name fail to
// validate
func TestValidateMissingName(t *testing.T) {

	// initialize a valid cheatpath
	cheatpath := Cheatpath{
		Path:     "/foo",
		ReadOnly: false,
		Tags:     []string{},
	}

	// assert that no errors are returned
	if err := cheatpath.Validate(); err == nil {
		t.Errorf("failed to invalidate cheatpath without name")
	}
}

// TestValidateMissingPath asserts that paths that are missing a path fail to
// validate
func TestValidateMissingPath(t *testing.T) {

	// initialize a valid cheatpath
	cheatpath := Cheatpath{
		Name:     "foo",
		ReadOnly: false,
		Tags:     []string{},
	}

	// assert that no errors are returned
	if err := cheatpath.Validate(); err == nil {
		t.Errorf("failed to invalidate cheatpath without path")
	}
}
