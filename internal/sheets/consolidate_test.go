package sheets

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/cheat/cheat/internal/sheet"
)

// TestConsolidate asserts that cheatsheets are properly consolidated
func TestConsolidate(t *testing.T) {

	// mock cheatsheets available on multiple cheatpaths
	cheatpaths := []map[string]sheet.Sheet{

		// mock community cheatsheets
		map[string]sheet.Sheet{
			"foo": sheet.Sheet{Title: "foo", Path: "community/foo"},
			"bar": sheet.Sheet{Title: "bar", Path: "community/bar"},
		},

		// mock local cheatsheets
		map[string]sheet.Sheet{
			"bar": sheet.Sheet{Title: "bar", Path: "local/bar"},
			"baz": sheet.Sheet{Title: "baz", Path: "local/baz"},
		},
	}

	// consolidate the cheatsheets
	consolidated := Consolidate(cheatpaths)

	// specify the expected output
	want := map[string]sheet.Sheet{
		"foo": sheet.Sheet{Title: "foo", Path: "community/foo"},
		"bar": sheet.Sheet{Title: "bar", Path: "local/bar"},
		"baz": sheet.Sheet{Title: "baz", Path: "local/baz"},
	}

	// assert that the cheatsheets properly consolidated
	if !reflect.DeepEqual(consolidated, want) {
		t.Errorf(
			"failed to consolidate cheatpaths: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(consolidated),
		)
	}

}
