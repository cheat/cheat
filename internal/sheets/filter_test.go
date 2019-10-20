package sheets

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/cheat/cheat/internal/sheet"
)

// TestFilterSingleTag asserts that Filter properly filters results when passed
// a single tag
func TestFilterSingleTag(t *testing.T) {

	// mock cheatsheets available on multiple cheatpaths
	cheatpaths := []map[string]sheet.Sheet{

		map[string]sheet.Sheet{
			"foo": sheet.Sheet{Title: "foo", Tags: []string{"alpha", "bravo"}},
			"bar": sheet.Sheet{Title: "bar", Tags: []string{"bravo", "charlie"}},
		},

		map[string]sheet.Sheet{
			"baz": sheet.Sheet{Title: "baz", Tags: []string{"alpha", "bravo"}},
			"bat": sheet.Sheet{Title: "bat", Tags: []string{"bravo", "charlie"}},
		},
	}

	// filter the cheatsheets
	filtered := Filter(cheatpaths, []string{"bravo"})

	// assert that the expect results were returned
	want := []map[string]sheet.Sheet{
		map[string]sheet.Sheet{
			"foo": sheet.Sheet{Title: "foo", Tags: []string{"alpha", "bravo"}},
			"bar": sheet.Sheet{Title: "bar", Tags: []string{"bravo", "charlie"}},
		},

		map[string]sheet.Sheet{
			"baz": sheet.Sheet{Title: "baz", Tags: []string{"alpha", "bravo"}},
			"bat": sheet.Sheet{Title: "bat", Tags: []string{"bravo", "charlie"}},
		},
	}

	if !reflect.DeepEqual(filtered, want) {
		t.Errorf(
			"failed to return expected results: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(filtered),
		)
	}
}

// TestFilterSingleTag asserts that Filter properly filters results when passed
// multiple tags
func TestFilterMultiTag(t *testing.T) {

	// mock cheatsheets available on multiple cheatpaths
	cheatpaths := []map[string]sheet.Sheet{

		map[string]sheet.Sheet{
			"foo": sheet.Sheet{Title: "foo", Tags: []string{"alpha", "bravo"}},
			"bar": sheet.Sheet{Title: "bar", Tags: []string{"bravo", "charlie"}},
		},

		map[string]sheet.Sheet{
			"baz": sheet.Sheet{Title: "baz", Tags: []string{"alpha", "bravo"}},
			"bat": sheet.Sheet{Title: "bat", Tags: []string{"bravo", "charlie"}},
		},
	}

	// filter the cheatsheets
	filtered := Filter(cheatpaths, []string{"alpha", "bravo"})

	// assert that the expect results were returned
	want := []map[string]sheet.Sheet{
		map[string]sheet.Sheet{
			"foo": sheet.Sheet{Title: "foo", Tags: []string{"alpha", "bravo"}},
		},

		map[string]sheet.Sheet{
			"baz": sheet.Sheet{Title: "baz", Tags: []string{"alpha", "bravo"}},
		},
	}

	if !reflect.DeepEqual(filtered, want) {
		t.Errorf(
			"failed to return expected results: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(filtered),
		)
	}
}
