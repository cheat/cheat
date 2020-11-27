package display

import "testing"

// TestIndent asserts that Indent prepends a tab to each line
func TestIndent(t *testing.T) {
	got := Indent("foo\nbar\nbaz")
	want := "\tfoo\n\tbar\n\tbaz"
	if got != want {
		t.Errorf("failed to indent: want: %s, got: %s", want, got)
	}
}
