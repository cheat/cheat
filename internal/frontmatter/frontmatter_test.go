package frontmatter

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
	text, fm, err := Parse(markdown)

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
	text, fm, err := Parse(markdown)

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
	text, _, err := Parse(markdown)

	// assert that an error was returned
	if err == nil {
		t.Error("failed to error on invalid frontmatter")
	}

	// assert that the "raw" markdown was returned
	if text != markdown {
		t.Errorf("failed to parse text: want: %s, got: %s", markdown, text)
	}
}

// TestTrimEmptyLines assert that leading and trailing empty lines are removed
func TestTrimEmptyLines(t *testing.T) {

	// define for readability of the tests
	blank := "\x20"
	linefeed := "\x0A"

	testCases := []struct {
		input string
		want  string
	}{
		// nothing to be trimmed
		{"", ""},
		{"a", "a"},
		{" non empty line ", " non empty line "},
		{" non empty line " + linefeed, " non empty line " + linefeed},
		{" non empty line " + linefeed + " another non empty line ", " non empty line " + linefeed + " another non empty line "},
		{" non empty line " + linefeed + " another non empty line " + linefeed, " non empty line " + linefeed + " another non empty line " + linefeed},
		// trim leading empty lines
		{linefeed + " non empty line ", " non empty line "},
		{blank + linefeed + " non empty line ", " non empty line "},
		{blank + linefeed + linefeed + " non empty line ", " non empty line "},
		{linefeed + blank + linefeed + " non empty line ", " non empty line "},
		{blank + linefeed + blank + linefeed + " non empty line ", " non empty line "},
		{linefeed + " non empty line " + linefeed + " another non empty line ", " non empty line " + linefeed + " another non empty line "},
		{blank, ""},
		{linefeed, ""},
		{blank + linefeed, ""},
		{linefeed + blank, ""},
		{blank + linefeed + blank, ""},
		{linefeed + blank + linefeed, ""},
		// trim trailing empty lines
		{" non empty line " + linefeed + blank, " non empty line " + linefeed},
		{" non empty line " + linefeed + blank + linefeed, " non empty line " + linefeed},
		{" non empty line " + linefeed + blank + linefeed + blank, " non empty line " + linefeed},
	}

	for _, testCase := range testCases {
		// parse the input
		text, _, err := Parse(testCase.input)

		// assert expectations
		if err != nil {
			t.Errorf("failed to parse input: %v", err)
		}

		// assert that the wanted output was returned
		if text != testCase.want {
			t.Errorf("failed to parse text: want: [%s], got: [%s]", testCase.want, text)
		}
	}
}
