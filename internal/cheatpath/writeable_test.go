package cheatpath

import (
	"testing"
)

// TestWriteableOK asserts that Writeable returns the appropriate cheatpath
// when a writeable cheatpath exists
func TestWriteableOK(t *testing.T) {

	// initialize some cheatpaths
	cheatpaths := []Cheatpath{
		Cheatpath{Path: "/foo", ReadOnly: true},
		Cheatpath{Path: "/bar", ReadOnly: false},
		Cheatpath{Path: "/baz", ReadOnly: true},
	}

	// get the writeable cheatpath
	got, err := Writeable(cheatpaths)

	// assert that no errors were returned
	if err != nil {
		t.Errorf("failed to get cheatpath: %v", err)
	}

	// assert that the path is correct
	if got.Path != "/bar" {
		t.Errorf("incorrect cheatpath returned: got: %s", got.Path)
	}
}

// TestWriteableNotOK asserts that Writeable returns an error when no writeable
// cheatpaths exist
func TestWriteableNotOK(t *testing.T) {

	// initialize some cheatpaths
	cheatpaths := []Cheatpath{
		Cheatpath{Path: "/foo", ReadOnly: true},
		Cheatpath{Path: "/bar", ReadOnly: true},
		Cheatpath{Path: "/baz", ReadOnly: true},
	}

	// get the writeable cheatpath
	_, err := Writeable(cheatpaths)

	// assert that no errors were returned
	if err == nil {
		t.Errorf("failed to return an error when no writeable paths found")
	}
}
