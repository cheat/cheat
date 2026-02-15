package sheets

import (
	"sort"
	"unicode/utf8"

	"github.com/cheat/cheat/internal/sheet"
)

// Tags returns a slice of all tags in use in any sheet
func Tags(cheatpaths []map[string]sheet.Sheet) []string {

	// create a map of all tags in use in any sheet
	tags := make(map[string]bool)

	// iterate over all tags on all sheets on all cheatpaths
	for _, path := range cheatpaths {
		for _, sheet := range path {
			for _, tag := range sheet.Tags {
				// Skip invalid UTF-8 tags to prevent downstream issues
				if utf8.ValidString(tag) {
					tags[tag] = true
				}
			}
		}
	}

	// restructure the map into a slice
	sorted := []string{}
	for tag := range tags {
		sorted = append(sorted, tag)
	}

	// sort the slice
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})

	return sorted
}
