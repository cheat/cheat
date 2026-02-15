package sheet

import (
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
	s := Sheet{
		Text: "echo 'foo'",
	}

	// colorize the sheet text
	s.Colorize(conf)

	// initialize expectations
	want := "[38;2;181;137;0mecho[0m[38;2;147;161;161m"
	want += " [0m[38;2;42;161;152m'foo'[0m"

	// assert
	if s.Text != want {
		t.Errorf("failed to colorize sheet: want: %s, got: %s", want, s.Text)
	}
}

// TestColorizeError tests the error handling in Colorize
func TestColorizeError(_ *testing.T) {
	// Create a sheet with content
	sheet := Sheet{
		Text:   "some text",
		Syntax: "invalidlexer12345", // Use an invalid lexer that might cause issues
	}

	// Create a config with invalid formatter/style
	conf := config.Config{
		Formatter: "invalidformatter",
		Style:     "invalidstyle",
	}

	// Store original text
	originalText := sheet.Text

	// Colorize should not panic even with invalid settings
	sheet.Colorize(conf)

	// The text might be unchanged if there was an error, or it might be colorized
	// We're mainly testing that it doesn't panic
	_ = sheet.Text
	_ = originalText
}
