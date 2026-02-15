// Package cheatpath manages collections of cheat sheets organized in filesystem directories.
//
// A Cheatpath represents a directory containing cheat sheets, with associated
// metadata such as tags and read-only status. Multiple cheatpaths can be
// configured to organize sheets from different sources (personal, community, work, etc.).
//
// # Cheatpath Structure
//
// Each cheatpath has:
//   - Name: A friendly identifier (e.g., "personal", "community")
//   - Path: The filesystem path to the directory
//   - Tags: Tags automatically applied to all sheets in this path
//   - ReadOnly: Whether sheets in this path can be modified
//
// Example configuration:
//
//	cheatpaths:
//	  - name: personal
//	    path: ~/cheat
//	    tags: []
//	    readonly: false
//	  - name: community
//	    path: ~/cheat/community
//	    tags: [community]
//	    readonly: true
//
// # Directory-Scoped Cheatpaths
//
// The package supports directory-scoped cheatpaths via `.cheat` directories.
// When running cheat from a directory containing a `.cheat` subdirectory,
// that directory is temporarily added to the available cheatpaths.
//
// # Precedence and Overrides
//
// When multiple cheatpaths contain a sheet with the same name, the sheet
// from the most "local" cheatpath takes precedence. This allows users to
// override community sheets with personal versions.
//
// Key Functions
//
//   - Filter: Filters cheatpaths by name
//   - Validate: Ensures cheatpath configuration is valid
//   - Writeable: Returns the first writeable cheatpath
//
// Example Usage
//
//	// Filter cheatpaths to only "personal"
//	filtered, err := cheatpath.Filter(paths, "personal")
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Find a writeable cheatpath
//	writeable, err := cheatpath.Writeable(paths)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Validate cheatpath configuration
//	if err := cheatpath.Validate(paths); err != nil {
//	    log.Fatal(err)
//	}
package cheatpath
