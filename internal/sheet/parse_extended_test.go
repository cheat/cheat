package sheet

import (
	"runtime"
	"testing"
)

// TestParseWindowsLineEndings tests parsing with Windows line endings
func TestParseWindowsLineEndings(t *testing.T) {
	// Only test Windows line endings on Windows
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows line ending test on non-Windows platform")
	}

	// stub our cheatsheet content with Windows line endings
	markdown := "---\r\nsyntax: go\r\ntags: [ test ]\r\n---\r\nTo foo the bar: baz"

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
}

// TestParseInvalidYAML tests parsing with invalid YAML in frontmatter
func TestParseInvalidYAML(t *testing.T) {
	// stub our cheatsheet content with invalid YAML
	markdown := `---
syntax: go
tags: [ test
  unclosed bracket
---
To foo the bar: baz`

	// parse the frontmatter
	_, _, err := parse(markdown)

	// assert that an error was returned for invalid YAML
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}
