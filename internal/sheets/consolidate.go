// Package sheets implements functions pertaining to loading, sorting,
// filtering, and tagging cheatsheets.
package sheets

import (
	"github.com/cheat/cheat/internal/sheet"
)

// Consolidate applies cheatsheet "overrides", resolving title conflicts that
// exist among cheatpaths by preferring more local cheatsheets over less local
// cheatsheets.
func Consolidate(
	cheatpaths []map[string]sheet.Sheet,
) map[string]sheet.Sheet {

	consolidated := make(map[string]sheet.Sheet)

	for _, cheatpath := range cheatpaths {
		for title, sheet := range cheatpath {
			consolidated[title] = sheet
		}
	}

	return consolidated
}
