package cheatpath

import (
	"testing"
)

// TestFilterSuccess asserts that the proper cheatpath is returned when the
// requested cheatpath exists
func TestFilterSuccess(t *testing.T) {

	// init cheatpaths
	paths := []Cheatpath{
		Cheatpath{Name: "foo"},
		Cheatpath{Name: "bar"},
		Cheatpath{Name: "baz"},
	}

	// filter the paths
	paths, err := Filter(paths, "bar")
	if err != nil {
		t.Errorf("failed to filter paths: %v", err)
	}

	// assert that the expected path was returned
	if len(paths) != 1 {
		t.Errorf(
			"failed to return correct path count: want: 1, got: %d",
			len(paths),
		)
	}

	if paths[0].Name != "bar" {
		t.Errorf("failed to return correct path: want: bar, got: %s", paths[0].Name)
	}
}

// TestFilterFailure asserts that an error is returned when a non-existent
// cheatpath is requested
func TestFilterFailure(t *testing.T) {

	// init cheatpaths
	paths := []Cheatpath{
		Cheatpath{Name: "foo"},
		Cheatpath{Name: "bar"},
		Cheatpath{Name: "baz"},
	}

	// filter the paths
	_, err := Filter(paths, "qux")
	if err == nil {
		t.Errorf("failed to return an error on non-existent cheatpath")
	}
}
