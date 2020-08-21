package sheets

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/cheat/cheat/internal/sheet"
)

// TestTags asserts that cheatsheet tags are properly returned
func TestTags(t *testing.T) {

	// mock cheatsheets available on multiple cheatpaths
	cheatpaths := []map[string]sheet.Sheet{

		// mock community cheatsheets
		map[string]sheet.Sheet{
			"foo": sheet.Sheet{Title: "foo", Tags: []string{"alpha"}},
			"bar": sheet.Sheet{Title: "bar", Tags: []string{"alpha", "bravo"}},
		},

		// mock local cheatsheets
		map[string]sheet.Sheet{
			"bar": sheet.Sheet{Title: "bar", Tags: []string{"bravo", "charlie"}},
			"baz": sheet.Sheet{Title: "baz", Tags: []string{"delta"}},
		},
	}

	// consolidate the cheatsheets
	tags := Tags(cheatpaths)

	// specify the expected output
	want := []string{
		"alpha",
		"bravo",
		"charlie",
		"delta",
	}

	// assert that the cheatsheets properly consolidated
	if !reflect.DeepEqual(tags, want) {
		t.Errorf(
			"failed to return tags: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(tags),
		)
	}

}
