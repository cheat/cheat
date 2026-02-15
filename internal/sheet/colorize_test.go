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

// TestColorizeDefaultSyntax asserts that when no syntax is specified, the
// default ("bash") is used and produces the same output as an explicit "bash"
func TestColorizeDefaultSyntax(t *testing.T) {

	conf := config.Config{
		Formatter: "terminal16m",
		Style:     "monokai",
	}

	// use bash-specific content that tokenizes differently across lexers
	code := "if [[ -f /etc/passwd ]]; then\n  echo \"found\" | grep -o found\nfi"

	// colorize with empty syntax (should default to "bash")
	noSyntax := Sheet{Text: code}
	noSyntax.Colorize(conf)

	// colorize with explicit "bash" syntax
	bashSyntax := Sheet{Text: code, Syntax: "bash"}
	bashSyntax.Colorize(conf)

	// both should produce the same output
	if noSyntax.Text != bashSyntax.Text {
		t.Errorf(
			"default syntax does not match explicit bash:\ndefault: %q\nexplicit: %q",
			noSyntax.Text,
			bashSyntax.Text,
		)
	}
}

// TestColorizeExplicitSyntax asserts that a specified syntax is used
func TestColorizeExplicitSyntax(t *testing.T) {

	conf := config.Config{
		Formatter: "terminal16m",
		Style:     "monokai",
	}

	// colorize as bash
	bashSheet := Sheet{Text: "def hello():\n    pass", Syntax: "bash"}
	bashSheet.Colorize(conf)

	// colorize as python
	pySheet := Sheet{Text: "def hello():\n    pass", Syntax: "python"}
	pySheet.Colorize(conf)

	// different lexers should produce different output for Python code
	if bashSheet.Text == pySheet.Text {
		t.Error("bash and python syntax produced identical output")
	}
}
