// Package display handles output formatting and presentation for the cheat application.
//
// The display package provides utilities for:
//   - Writing output to stdout or a pager
//   - Formatting text with indentation
//   - Creating faint (dimmed) text for de-emphasis
//   - Managing colored output
//
// # Pager Integration
//
// The package integrates with system pagers (less, more, etc.) to handle
// long output. If a pager is configured and the output is to a terminal,
// content is automatically piped through the pager.
//
// # Text Formatting
//
// Various formatting utilities are provided:
//   - Faint: Creates dimmed text using ANSI escape codes
//   - Indent: Adds consistent indentation to text blocks
//   - Write: Intelligent output that uses stdout or pager as appropriate
//
// Example Usage
//
//	// Write output, using pager if configured
//	if err := display.Write(output, config); err != nil {
//	    log.Fatal(err)
//	}
//
//	// Create faint text for de-emphasis
//	fainted := display.Faint("(read-only)", config)
//
//	// Indent a block of text
//	indented := display.Indent(text, "  ")
//
// # Color Support
//
// The package respects the colorization settings from the config.
// When colorization is disabled, formatting functions like Faint
// return unmodified text.
//
// # Terminal Detection
//
// The package uses isatty to detect if output is to a terminal,
// which affects decisions about using a pager and applying colors.
package display
