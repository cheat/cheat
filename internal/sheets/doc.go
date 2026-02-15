// Package sheets manages collections of cheat sheets across multiple cheatpaths.
//
// The sheets package provides functionality to:
//   - Load sheets from multiple cheatpaths
//   - Consolidate duplicate sheets (with precedence rules)
//   - Filter sheets by tags
//   - Sort sheets alphabetically
//   - Extract unique tags across all sheets
//
// # Loading Sheets
//
// Sheets are loaded recursively from cheatpath directories, excluding:
//   - Hidden files (starting with .)
//   - Files in .git directories
//   - Files with extensions (sheets have no extension)
//
// # Consolidation
//
// When multiple cheatpaths contain sheets with the same name, consolidation
// rules apply based on the order of cheatpaths. Sheets from earlier paths
// override those from later paths, allowing personal sheets to override
// community sheets.
//
// Example:
//
//	cheatpaths:
//	  1. personal: ~/cheat
//	  2. community: ~/cheat/community
//
//	If both contain "git", the version from "personal" is used.
//
// # Filtering
//
// Sheets can be filtered by:
//   - Tags: Include only sheets with specific tags
//   - Cheatpath: Include only sheets from specific paths
//
// Key Functions
//
//   - Load: Loads all sheets from the given cheatpaths
//   - Filter: Filters sheets by tag
//   - Consolidate: Merges sheets from multiple paths with precedence
//   - Sort: Sorts sheets alphabetically by title
//   - Tags: Extracts all unique tags from sheets
//
// Example Usage
//
//	// Load sheets from all cheatpaths
//	allSheets, err := sheets.Load(cheatpaths)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Consolidate to handle duplicates
//	consolidated := sheets.Consolidate(allSheets)
//
//	// Filter by tag
//	filtered := sheets.Filter(consolidated, "networking")
//
//	// Sort alphabetically
//	sheets.Sort(filtered)
//
//	// Get all unique tags
//	tags := sheets.Tags(consolidated)
package sheets
