package sheets

import (
	"sort"

	"github.com/cheat/cheat/internal/sheet"
)

// Sort organizes the cheatsheets into an alphabetically-sorted slice
func Sort(cheatsheets map[string]sheet.Sheet) []sheet.Sheet {

	// create a slice that contains the cheatsheet titles
	var titles []string
	for title := range cheatsheets {
		titles = append(titles, title)
	}

	// sort the slice of titles
	sort.Strings(titles)

	// create a slice of sorted cheatsheets
	sorted := []sheet.Sheet{}

	// iterate over the sorted slice of titles, and append cheatsheets to
	// `sorted` in an identical (alabetically sequential) order
	for _, title := range titles {
		sorted = append(sorted, cheatsheets[title])
	}

	// return the sorted slice of cheatsheets
	return sorted
}
