package display

import (
	"testing"

	"github.com/cheat/cheat/internal/config"
)

// TestFaint asserts that Faint applies faint formatting
func TestFaint(t *testing.T) {

	// case: apply colorization
	conf := config.Config{Colorize: true}
	want := "\033[2mfoo\033[0m"
	got := Faint("foo", conf)
	if want != got {
		t.Errorf("failed to faint: want: %s, got: %s", want, got)
	}

	// case: do not apply colorization
	conf.Colorize = false
	want = "foo"
	got = Faint("foo", conf)
	if want != got {
		t.Errorf("failed to faint: want: %s, got: %s", want, got)
	}
}
