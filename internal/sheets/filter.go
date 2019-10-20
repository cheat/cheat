package sheets

import (
	"strings"

	"github.com/cheat/cheat/internal/sheet"
)

// Filter filters cheatsheets that do not match `tag(s)`
func Filter(
	cheatpaths []map[string]sheet.Sheet,
	tags []string,
) []map[string]sheet.Sheet {

	// buffer a map of filtered cheatsheets
	filtered := make([]map[string]sheet.Sheet, 0, len(cheatpaths))

	// iterate over each cheatpath
	for _, cheatsheets := range cheatpaths {

		// create a map of cheatsheets for each cheatpath. The filtering will be
		// applied to each cheatpath individually.
		pathFiltered := make(map[string]sheet.Sheet)

		// iterate over each cheatsheet that exists on each cheatpath
		for title, sheet := range cheatsheets {

			// assume that the sheet should be kept (ie, should not be filtered)
			keep := true

			// iterate over each tag. If the sheet does not match *all* tags, filter
			// it out.
			for _, tag := range tags {
				if !sheet.Tagged(strings.TrimSpace(tag)) {
					keep = false
				}
			}

			// if the sheet does match all tags, it passes the filter
			if keep {
				pathFiltered[title] = sheet
			}
		}

		// the sheets on this individual cheatpath have now been filtered. Now,
		// store those alongside the sheets on the other cheatpaths that also made
		// it passed the filter.
		filtered = append(filtered, pathFiltered)
	}

	// return the filtered cheatsheets on all paths
	return filtered
}
