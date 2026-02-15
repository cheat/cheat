package sheets

import (
	"path"
	"testing"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/mocks"
)

// TestLoad asserts that sheets on valid cheatpaths can be loaded successfully
func TestLoad(t *testing.T) {

	// mock cheatpaths
	cheatpaths := []cheatpath.Path{
		{
			Name:     "community",
			Path:     path.Join(mocks.Path("cheatsheets"), "community"),
			ReadOnly: true,
		},
		{
			Name:     "personal",
			Path:     path.Join(mocks.Path("cheatsheets"), "personal"),
			ReadOnly: false,
		},
	}

	// load cheatsheets
	cheatpathSheets, err := Load(cheatpaths)
	if err != nil {
		t.Errorf("failed to load cheatsheets: %v", err)
	}

	// assert that the correct number of sheets loaded
	// (sheet load details are tested in `sheet_test.go`)
	totalSheets := 0
	for _, sheets := range cheatpathSheets {
		totalSheets += len(sheets)
	}

	// we expect 4 total sheets (2 from community, 2 from personal)
	// hidden files and files with extensions are excluded
	want := 4
	if totalSheets != want {
		t.Errorf(
			"failed to load correct number of cheatsheets: want: %d, got: %d",
			want,
			totalSheets,
		)
	}
}

// TestLoadBadPath asserts that an error is returned if a cheatpath is invalid
func TestLoadBadPath(t *testing.T) {

	// mock a bad cheatpath
	cheatpaths := []cheatpath.Path{
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
