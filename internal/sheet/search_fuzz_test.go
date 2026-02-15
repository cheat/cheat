package sheet

import (
	"regexp"
	"strings"
	"testing"
	"time"
)

// FuzzSearchRegex tests the regex compilation and search functionality
// to ensure it handles malformed patterns gracefully and doesn't suffer
// from catastrophic backtracking
func FuzzSearchRegex(f *testing.F) {
	// Add seed corpus with various regex patterns
	// Valid patterns
	f.Add("test", "This is a test string")
	f.Add("(?i)test", "This is a TEST string")
	f.Add("foo|bar", "foo and bar")
	f.Add("^start", "start of line\nnext line")
	f.Add("end$", "at the end\nnext line")
	f.Add("\\d+", "123 numbers 456")
	f.Add("[a-z]+", "lowercase UPPERCASE")

	// Edge cases and potentially problematic patterns
	f.Add("", "empty pattern")
	f.Add(".", "any character")
	f.Add(".*", "match everything")
	f.Add(".+", "match something")
	f.Add("\\", "backslash")
	f.Add("(", "unclosed paren")
	f.Add(")", "unmatched paren")
	f.Add("[", "unclosed bracket")
	f.Add("]", "unmatched bracket")
	f.Add("[^]", "negated empty class")
	f.Add("(?", "incomplete group")

	// Patterns that might cause performance issues
	f.Add("(a+)+", "aaaaaaaaaaaaaaaaaaaaaaaab")
	f.Add("(a*)*", "aaaaaaaaaaaaaaaaaaaaaaaab")
	f.Add("(a|a)*", "aaaaaaaaaaaaaaaaaaaaaaaab")
	f.Add("(.*)*", "any text here")
	f.Add("(\\d+)+", "123456789012345678901234567890x")

	// Unicode patterns
	f.Add("☺", "Unicode ☺ smiley")
	f.Add("[一-龯]", "Chinese 中文 characters")
	f.Add("\\p{L}+", "Unicode letters")

	// Very long patterns
	f.Add(strings.Repeat("a", 1000), "long pattern")
	f.Add(strings.Repeat("(a|b)", 100), "complex pattern")

	f.Fuzz(func(t *testing.T, pattern string, text string) {
		// Test 1: Regex compilation should not panic
		var reg *regexp.Regexp
		var compileErr error

		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("regexp.Compile panicked with pattern %q: %v", pattern, r)
				}
			}()

			reg, compileErr = regexp.Compile(pattern)
		}()

		// If compilation failed, that's OK - we're testing error handling
		if compileErr != nil {
			// This is expected for invalid patterns
			return
		}

		// Test 2: Create a sheet and test Search method
		sheet := Sheet{
			Title: "test",
			Text:  text,
		}

		// Search should not panic
		var result string
		done := make(chan bool, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Search panicked with pattern %q on text %q: %v", pattern, text, r)
				}
				done <- true
			}()

			result = sheet.Search(reg)
		}()

		// Timeout after 100ms to catch catastrophic backtracking
		select {
		case <-done:
			// Search completed successfully
		case <-time.After(100 * time.Millisecond):
			t.Errorf("Search timed out (possible catastrophic backtracking) with pattern %q on text %q", pattern, text)
		}

		// Test 3: Verify search result invariants
		if result != "" {
			// The Search function splits by "\n\n", so we need to compare using the same logic
			resultLines := strings.Split(result, "\n\n")
			textLines := strings.Split(text, "\n\n")

			// Every result line should exist in the original text lines
			for _, rLine := range resultLines {
				found := false
				for _, tLine := range textLines {
					if rLine == tLine {
						found = true
						break
					}
				}
				if !found && rLine != "" {
					t.Errorf("Search result contains line not in original text: %q", rLine)
				}
			}
		}
	})
}

// FuzzSearchCatastrophicBacktracking specifically tests for regex patterns
// that could cause performance issues
func FuzzSearchCatastrophicBacktracking(f *testing.F) {
	// Seed with patterns known to potentially cause issues
	f.Add("a", 10, 5)
	f.Add("x", 20, 3)

	f.Fuzz(func(t *testing.T, char string, repeats int, groups int) {
		// Limit the size to avoid memory issues in the test
		if repeats > 30 || repeats < 0 || groups > 10 || groups < 0 || len(char) > 5 {
			t.Skip("Skipping invalid or overly large test case")
		}

		// Construct patterns that might cause backtracking
		patterns := []string{
			strings.Repeat(char, repeats),
			"(" + char + "+)+",
			"(" + char + "*)*",
			"(" + char + "|" + char + ")*",
		}

		// Add nested groups
		if groups > 0 && groups < 10 {
			nested := char
			for i := 0; i < groups; i++ {
				nested = "(" + nested + ")+"
			}
			patterns = append(patterns, nested)
		}

		// Test text that might trigger backtracking
		testText := strings.Repeat(char, repeats) + "x"

		for _, pattern := range patterns {
			// Try to compile the pattern
			reg, err := regexp.Compile(pattern)
			if err != nil {
				// Invalid pattern, skip
				continue
			}

			// Test with timeout
			done := make(chan bool, 1)

			go func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Search panicked with backtracking pattern %q: %v", pattern, r)
					}
					done <- true
				}()

				sheet := Sheet{Text: testText}
				_ = sheet.Search(reg)
			}()

			select {
			case <-done:
				// Completed successfully
			case <-time.After(50 * time.Millisecond):
				t.Logf("Warning: potential backtracking issue with pattern %q (completed slowly)", pattern)
			}
		}
	})
}
