package cheatpath

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// FuzzValidateSheetName tests the ValidateSheetName function with fuzzing
// to ensure it properly prevents path traversal and other security issues
func FuzzValidateSheetName(f *testing.F) {
	// Add seed corpus with various valid and malicious inputs
	// Valid names
	f.Add("docker")
	f.Add("docker/compose")
	f.Add("lang/go/slice")
	f.Add("my-cheat_sheet")
	f.Add("file.txt")
	f.Add("a")
	f.Add("123")

	// Path traversal attempts
	f.Add("..")
	f.Add("../etc/passwd")
	f.Add("foo/../bar")
	f.Add("foo/../../etc/passwd")
	f.Add("..\\windows\\system32")
	f.Add("foo\\..\\..\\windows")

	// Encoded traversal attempts
	f.Add("%2e%2e")
	f.Add("%2e%2e%2f")
	f.Add("..%2f")
	f.Add("%2e.")
	f.Add(".%2e")
	f.Add("\x2e\x2e")
	f.Add("\\x2e\\x2e")

	// Unicode and special characters
	f.Add("€test")
	f.Add("test€")
	f.Add("中文")
	f.Add("🎉emoji")
	f.Add("\x00null")
	f.Add("test\x00null")
	f.Add("\nnewline")
	f.Add("test\ttab")

	// Absolute paths
	f.Add("/etc/passwd")
	f.Add("C:\\Windows\\System32")
	f.Add("\\\\server\\share")
	f.Add("//server/share")

	// Home directory
	f.Add("~")
	f.Add("~/config")
	f.Add("~user/file")

	// Hidden files
	f.Add(".hidden")
	f.Add("dir/.hidden")
	f.Add(".git/config")

	// Edge cases
	f.Add("")
	f.Add(" ")
	f.Add("  ")
	f.Add("\t")
	f.Add(".")
	f.Add("./")
	f.Add("./file")
	f.Add(".../")
	f.Add("...")
	f.Add("....")

	// Very long names
	f.Add(strings.Repeat("a", 255))
	f.Add(strings.Repeat("a/", 100) + "file")
	f.Add(strings.Repeat("../", 50) + "etc/passwd")

	f.Fuzz(func(t *testing.T, input string) {
		// The function should never panic
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("ValidateSheetName panicked with input %q: %v", input, r)
				}
			}()

			err := ValidateSheetName(input)

			// Security invariants that must always hold
			if err == nil {
				// If validation passed, verify security properties

				// Should not contain ".." for path traversal
				if strings.Contains(input, "..") {
					t.Errorf("validation passed but input contains '..': %q", input)
				}

				// Should not be empty
				if input == "" {
					t.Error("validation passed for empty input")
				}

				// Should not start with ~ (home directory)
				if strings.HasPrefix(input, "~") {
					t.Errorf("validation passed but input starts with '~': %q", input)
				}

				// Base filename should not start with .
				parts := strings.Split(input, "/")
				if len(parts) > 0 {
					lastPart := parts[len(parts)-1]
					if strings.HasPrefix(lastPart, ".") && lastPart != "." {
						t.Errorf("validation passed but filename starts with '.': %q", input)
					}
				}

				// Additional check: result should be valid UTF-8
				if !utf8.ValidString(input) {
					// While the function doesn't explicitly check this,
					// we want to ensure it handles invalid UTF-8 gracefully
					t.Logf("validation passed for invalid UTF-8: %q", input)
				}
			}
		}()
	})
}

// FuzzValidateSheetNamePathTraversal specifically targets path traversal bypasses
func FuzzValidateSheetNamePathTraversal(f *testing.F) {
	// Seed corpus focusing on path traversal variations
	f.Add("..", "/", "")
	f.Add("", "..", "/")
	f.Add("a", "b", "c")

	f.Fuzz(func(t *testing.T, prefix string, middle string, suffix string) {
		// Construct various path traversal attempts
		inputs := []string{
			prefix + ".." + suffix,
			prefix + "/.." + suffix,
			prefix + "\\.." + suffix,
			prefix + middle + ".." + suffix,
			prefix + "../" + middle + suffix,
			prefix + "..%2f" + suffix,
			prefix + "%2e%2e" + suffix,
			prefix + "%2e%2e%2f" + suffix,
		}

		for _, input := range inputs {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("ValidateSheetName panicked with constructed input %q: %v", input, r)
					}
				}()

				err := ValidateSheetName(input)

				// If the input contains literal "..", it must be rejected
				if strings.Contains(input, "..") && err == nil {
					t.Errorf("validation incorrectly passed for input containing '..': %q", input)
				}
			}()
		}
	})
}
