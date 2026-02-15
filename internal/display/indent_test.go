package display

import "testing"

// TestIndent asserts that Indent prepends a tab to each line
func TestIndent(t *testing.T) {
	got := Indent("foo\nbar\nbaz")
	want := "\tfoo\n\tbar\n\tbaz\n"
	if got != want {
		t.Errorf("failed to indent: want: %s, got: %s", want, got)
	}
}

// TestIndentTrimsWhitespace asserts that Indent trims leading and trailing
// whitespace before indenting
func TestIndentTrimsWhitespace(t *testing.T) {
	got := Indent("  foo\nbar\nbaz  \n")
	want := "\tfoo\n\tbar\n\tbaz\n"
	if got != want {
		t.Errorf("failed to trim and indent: want: %q, got: %q", want, got)
	}
}
