package display

import (
	"testing"
)

// TestFaint asserts that Faint applies faint formatting
func TestFaint(t *testing.T) {
	want := "\033[2mfoo\033[0m"
	got := Faint("foo")
	if want != got {
		t.Errorf("failed to faint: want: %s, got: %s", want, got)
	}
}
