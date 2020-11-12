package sheets

import (
	"path"
	"testing"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/mock"
)

// TestLoad asserts that sheets on valid cheatpaths can be loaded successfully
func TestLoad(t *testing.T) {

	// mock cheatpaths
	cheatpaths := []cheatpath.Cheatpath{
		{
			Name:     "community",
			Path:     path.Join(mock.Path("cheatsheets"), "community"),
			ReadOnly: true,
		},
		{
			Name:     "personal",
			Path:     path.Join(mock.Path("cheatsheets"), "personal"),
			ReadOnly: false,
		},
	}

	// load cheatsheets
	sheets, err := Load(cheatpaths)
	if err != nil {
		t.Errorf("failed to load cheatsheets: %v", err)
	}

	// assert that the correct number of sheets loaded
	// (sheet load details are tested in `sheet_test.go`)
	want := 4
	if len(sheets) != want {
		t.Errorf(
			"failed to load correct number of cheatsheets: want: %d, got: %d",
			want,
			len(sheets),
		)
	}
}

// TestLoadBadPath asserts that an error is returned if a cheatpath is invalid
func TestLoadBadPath(t *testing.T) {

	// mock a bad cheatpath
	cheatpaths := []cheatpath.Cheatpath{
		{
			Name:     "badpath",
			Path:     "/cheat/test/path/does/not/exist",
			ReadOnly: true,
		},
	}

	// attempt to load the cheatpath
	if _, err := Load(cheatpaths); err == nil {
		t.Errorf("failed to reject invalid cheatpath")
	}
}
