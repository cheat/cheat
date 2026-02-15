package sheet

import (
	"strings"
	"testing"
)

// FuzzParse tests the parse function with fuzzing to uncover edge cases
// and potential panics in YAML frontmatter parsing
func FuzzParse(f *testing.F) {
	// Add seed corpus with various valid and edge case inputs
	// Valid frontmatter
	f.Add("---\nsyntax: go\n---\nContent")
	f.Add("---\ntags: [a, b]\n---\n")
	f.Add("---\nsyntax: bash\ntags: [linux, shell]\n---\n#!/bin/bash\necho hello")

	// No frontmatter
	f.Add("No frontmatter here")
	f.Add("")
	f.Add("Just plain text\nwith multiple lines")

	// Edge cases with delimiters
	f.Add("---")
	f.Add("---\n")
	f.Add("---\n---")
	f.Add("---\n---\n")
	f.Add("---\n---\n---")
	f.Add("---\n---\n---\n---")
	f.Add("------\n------")

	// Invalid YAML
	f.Add("---\n{invalid yaml\n---\n")
	f.Add("---\nsyntax: \"unclosed quote\n---\n")
	f.Add("---\ntags: [a, b,\n---\n")

	// Windows line endings
	f.Add("---\r\nsyntax: go\r\n---\r\nContent")
	f.Add("---\r\n---\r\n")

	// Mixed line endings
	f.Add("---\nsyntax: go\r\n---\nContent")
	f.Add("---\r\nsyntax: go\n---\r\nContent")

	// Unicode and special characters
	f.Add("---\ntags: [emoji, 🎉]\n---\n")
	f.Add("---\nsyntax: 中文\n---\n")
	f.Add("---\ntags: [\x00, \x01]\n---\n")

	// Very long inputs
	f.Add("---\ntags: [" + strings.Repeat("a,", 1000) + "a]\n---\n")
	f.Add("---\n" + strings.Repeat("field: value\n", 1000) + "---\n")

	// Nested structures
	f.Add("---\ntags:\n  - nested\n  - list\n---\n")
	f.Add("---\nmeta:\n  author: test\n  version: 1.0\n---\n")

	f.Fuzz(func(t *testing.T, input string) {
		// The parse function should never panic, regardless of input
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("parse panicked with input %q: %v", input, r)
				}
			}()

			fm, text, err := parse(input)

			// Verify invariants
			if err == nil {
				// If parsing succeeded, validate the result

				// The returned text should be a suffix of the input
				// (either the whole input if no frontmatter, or the part after frontmatter)
				if !strings.HasSuffix(input, text) && text != input {
					t.Errorf("returned text %q is not a valid suffix of input %q", text, input)
				}

				// If input starts with delimiter and has valid frontmatter,
				// text should be shorter than input
				if strings.HasPrefix(input, "---\n") || strings.HasPrefix(input, "---\r\n") {
					if len(fm.Tags) > 0 || fm.Syntax != "" {
						// We successfully parsed frontmatter, so text should be shorter
						if len(text) >= len(input) {
							t.Errorf("text length %d should be less than input length %d when frontmatter is parsed",
								len(text), len(input))
						}
					}
				}

				// Note: Tags can be nil when frontmatter is not present or empty
				// This is expected behavior in Go for uninitialized slices
			} else {
				// If parsing failed, the original input should be returned as text
				if text != input {
					t.Errorf("on error, text should equal input: got %q, want %q", text, input)
				}
			}
		}()
	})
}

// FuzzParseDelimiterHandling specifically tests delimiter edge cases
func FuzzParseDelimiterHandling(f *testing.F) {
	// Seed corpus focusing on delimiter variations
	f.Add("---", "content")
	f.Add("", "---")
	f.Add("---", "---")
	f.Add("", "")

	f.Fuzz(func(t *testing.T, prefix string, suffix string) {
		// Build input with controllable parts around delimiters
		inputs := []string{
			prefix + "---\n" + suffix,
			prefix + "---\r\n" + suffix,
			prefix + "---\n---\n" + suffix,
			prefix + "---\r\n---\r\n" + suffix,
			prefix + "---\n" + "yaml: data\n" + "---\n" + suffix,
		}

		for _, input := range inputs {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("parse panicked with constructed input: %v", r)
					}
				}()

				_, _, _ = parse(input)
			}()
		}
	})
}
