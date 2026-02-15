// Package sheet provides functionality for parsing and managing individual cheat sheets.
//
// A sheet represents a single cheatsheet file containing helpful commands, notes,
// or documentation. Sheets can include optional YAML frontmatter for metadata
// such as tags and syntax highlighting preferences.
//
// # Sheet Format
//
// Sheets are plain text files that may begin with YAML frontmatter:
//
//	---
//	syntax: bash
//	tags: [networking, linux, ssh]
//	---
//	# Connect to remote server
//	ssh user@hostname
//
//	# Copy files over SSH
//	scp local_file user@hostname:/remote/path
//
// The frontmatter is optional. If omitted, the sheet will use default values.
//
// # Core Types
//
// The Sheet type contains:
//   - Title: The sheet's name (derived from filename)
//   - Path: Full filesystem path to the sheet
//   - Text: The content of the sheet (without frontmatter)
//   - Tags: Categories assigned to the sheet
//   - Syntax: Language hint for syntax highlighting
//   - ReadOnly: Whether the sheet can be modified
//
// Key Functions
//
//   - New: Creates a new Sheet from a file path
//   - Parse: Extracts frontmatter and content from sheet text
//   - Search: Searches sheet content using regular expressions
//   - Colorize: Applies syntax highlighting to sheet content
//
// # Syntax Highlighting
//
// The package integrates with the Chroma library to provide syntax highlighting.
// Supported languages include bash, python, go, javascript, and many others.
// The syntax can be specified in the frontmatter or auto-detected.
//
// Example Usage
//
//	// Load a sheet from disk
//	s, err := sheet.New("/path/to/sheet", []string{"personal"}, false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Search for content
//	matches, err := s.Search("ssh", false)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Apply syntax highlighting
//	colorized, err := s.Colorize(config)
//	if err != nil {
//	    log.Fatal(err)
//	}
package sheet
