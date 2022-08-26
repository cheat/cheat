package sheet

import (
	"testing"
)

// TestHasFrontmatter asserts that markdown is properly parsed when it contains
// frontmatter
func TestHasFrontmatter(t *testing.T) {

	// stub our cheatsheet content
	markdown := `---
syntax: go
tags: [ test ]
---
To foo the bar: baz`

	// parse the frontmatter
	fm, text, err := parse(markdown)

	// assert expectations
	if err != nil {
		t.Errorf("failed to parse markdown: %v", err)
	}

	want := "To foo the bar: baz"
	if text != want {
		t.Errorf("failed to parse text: want: %s, got: %s", want, text)
	}

	want = "go"
	if fm.Syntax != want {
		t.Errorf("failed to parse syntax: want: %s, got: %s", want, fm.Syntax)
	}

	want = "test"
	if fm.Tags[0] != want {
		t.Errorf("failed to parse tags: want: %s, got: %s", want, fm.Tags[0])
	}
	if len(fm.Tags) != 1 {
		t.Errorf("failed to parse tags: want: len 0, got: len %d", len(fm.Tags))
	}
}

// TestHasFrontmatter asserts that markdown is properly parsed when it does not
// contain frontmatter
func TestHasNoFrontmatter(t *testing.T) {

	// stub our cheatsheet content
	markdown := "To foo the bar: baz"

	// parse the frontmatter
	fm, text, err := parse(markdown)

	// assert expectations
	if err != nil {
		t.Errorf("failed to parse markdown: %v", err)
	}

	if text != markdown {
		t.Errorf("failed to parse text: want: %s, got: %s", markdown, text)
	}

	if fm.Syntax != "" {
		t.Errorf("failed to parse syntax: want: '', got: %s", fm.Syntax)
	}

	if len(fm.Tags) != 0 {
		t.Errorf("failed to parse tags: want: len 0, got: len %d", len(fm.Tags))
	}
}

// TestHasInvalidFrontmatter asserts that markdown is properly parsed when it
// contains invalid frontmatter
func TestHasInvalidFrontmatter(t *testing.T) {

	// stub our cheatsheet content (with invalid frontmatter)
	markdown := `---
syntax: go
tags: [ test ]
To foo the bar: baz`

	// parse the frontmatter
	_, text, err := parse(markdown)

	// assert that an error was returned
	if err == nil {
		t.Error("failed to error on invalid frontmatter")
	}

	// assert that the "raw" markdown was returned
	if text != markdown {
		t.Errorf("failed to parse text: want: %s, got: %s", markdown, text)
	}
}
