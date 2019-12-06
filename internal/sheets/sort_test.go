package sheets

import (
	"testing"

	"github.com/cheat/cheat/internal/sheet"
)

// TestSort asserts that Sort properly sorts sheets
func TestSort(t *testing.T) {

	// mock a map of cheatsheets
	sheets := map[string]sheet.Sheet{
		"foo": sheet.Sheet{Title: "foo"},
		"bar": sheet.Sheet{Title: "bar"},
		"baz": sheet.Sheet{Title: "baz"},
	}

	// sort the sheets
	sorted := Sort(sheets)

	// assert that the sheets sorted properly
	want := []string{"bar", "baz", "foo"}

	for i, got := range sorted {
		if got.Title != want[i] {
			t.Errorf(
				"sort returned incorrect value: want: %s, got: %s",
				want[i],
				got.Title,
			)
		}
	}
}
