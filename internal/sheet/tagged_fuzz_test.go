package sheet

import (
	"strings"
	"testing"
)

// FuzzTagged tests the Tagged function with potentially malicious tag inputs
//
// Threat model: An attacker crafts a malicious cheatsheet with specially
// crafted tags that could cause issues when a user searches/filters by tags.
// This is particularly relevant for shared community cheatsheets.
func FuzzTagged(f *testing.F) {
	// Add seed corpus with potentially problematic inputs
	// These represent tags an attacker might use in a malicious cheatsheet
	f.Add("normal", "normal")
	f.Add("", "")
	f.Add(" ", " ")
	f.Add("\n", "\n")
	f.Add("\r\n", "\r\n")
	f.Add("\x00", "\x00")                         // Null byte
	f.Add("../../etc/passwd", "../../etc/passwd") // Path traversal attempt
	f.Add("'; DROP TABLE sheets;--", "sql")       // SQL injection attempt
	f.Add("<script>alert('xss')</script>", "xss") // XSS attempt
	f.Add("${HOME}", "${HOME}")                   // Environment variable
	f.Add("$(whoami)", "$(whoami)")               // Command substitution
	f.Add("`date`", "`date`")                     // Command substitution
	f.Add("\\x41\\x42", "\\x41\\x42")             // Escape sequences
	f.Add("%00", "%00")                           // URL encoded null
	f.Add("tag\nwith\nnewlines", "tag")
	f.Add(strings.Repeat("a", 10000), "a") // Very long tag
	f.Add("🎉", "🎉")                        // Unicode
	f.Add("\U0001F4A9", "\U0001F4A9")      // Unicode poop emoji
	f.Add("tag with spaces", "tag with spaces")
	f.Add("TAG", "tag") // Case sensitivity check
	f.Add("tag", "TAG") // Case sensitivity check

	f.Fuzz(func(t *testing.T, sheetTag string, searchTag string) {
		// Create a sheet with the potentially malicious tag
		sheet := Sheet{
			Title: "test",
			Tags:  []string{sheetTag},
		}

		// The Tagged function should never panic regardless of input
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Tagged panicked with sheetTag=%q, searchTag=%q: %v",
						sheetTag, searchTag, r)
				}
			}()

			result := sheet.Tagged(searchTag)

			// Verify the result is consistent with a simple string comparison
			expected := false
			for _, tag := range sheet.Tags {
				if tag == searchTag {
					expected = true
					break
				}
			}

			if result != expected {
				t.Errorf("Tagged returned %v but expected %v for sheetTag=%q, searchTag=%q",
					result, expected, sheetTag, searchTag)
			}

			// Additional invariant: Tagged should be case-sensitive
			if sheetTag != searchTag && result {
				t.Errorf("Tagged matched different strings: sheetTag=%q, searchTag=%q",
					sheetTag, searchTag)
			}
		}()

		// Test with multiple tags including the fuzzed one
		sheetMulti := Sheet{
			Title: "test",
			Tags:  []string{"safe1", sheetTag, "safe2", sheetTag}, // Duplicate tags
		}

		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Tagged panicked with multiple tags including %q: %v",
						sheetTag, r)
				}
			}()

			_ = sheetMulti.Tagged(searchTag)
		}()
	})
}
