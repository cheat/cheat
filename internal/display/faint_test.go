package display

import "testing"

// TestFaint asserts that Faint applies faint formatting
func TestFaint(t *testing.T) {

	// case: apply colorization
	want := "\033[2mfoo\033[0m"
	got := Faint("foo", true)
	if want != got {
		t.Errorf("failed to faint: want: %s, got: %s", want, got)
	}

	// case: do not apply colorization
	want = "foo"
	got = Faint("foo", false)
	if want != got {
		t.Errorf("failed to faint: want: %s, got: %s", want, got)
	}
}
