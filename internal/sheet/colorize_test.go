package sheet

import (
	"strings"
	"testing"

	"github.com/cheat/cheat/internal/config"
)

// TestColorize asserts that syntax-highlighting is correctly applied
func TestColorize(t *testing.T) {

	// mock configs
	conf := config.Config{
		Formatter: "terminal16m",
		Style:     "solarized-dark",
	}

	// mock a sheet
	original := "echo 'foo'"
	s := Sheet{
		Text: original,
	}

	// colorize the sheet text
	s.Colorize(conf)

	// assert that the text was modified (colorization applied)
	if s.Text == original {
		t.Error("Colorize did not modify sheet text")
	}

	// assert that ANSI escape codes are present
	if !strings.Contains(s.Text, "\x1b[") && !strings.Contains(s.Text, "[0m") {
		t.Errorf("colorized text does not contain ANSI escape codes: %q", s.Text)
	}

	// assert that the original content is still present within the colorized output
	if !strings.Contains(s.Text, "echo") || !strings.Contains(s.Text, "foo") {
		t.Errorf("colorized text lost original content: %q", s.Text)
	}
}
