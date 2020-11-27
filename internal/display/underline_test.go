package display

import (
	"testing"
)

// TestUnderline asserts that Underline applies underline formatting
func TestUnderline(t *testing.T) {
	want := "\033[4mfoo\033[0m"
	got := Underline("foo")
	if want != got {
		t.Errorf("failed to underline: want: %s, got: %s", want, got)
	}
}
