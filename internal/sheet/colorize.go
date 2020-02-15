package sheet

import (
	"bytes"

	"github.com/cheat/cheat/internal/config"

	"github.com/alecthomas/chroma/quick"
)

// Colorize applies syntax-highlighting to a cheatsheet's Text.
func (s *Sheet) Colorize(conf config.Config) {

	// if the syntax was not specified, default to bash
	lex := s.Syntax
	if lex == "" {
		lex = "bash"
	}

	// write colorized text into a buffer
	var buf bytes.Buffer
	err := quick.Highlight(
		&buf,
		s.Text,
		lex,
		conf.Formatter,
		conf.Style,
	)

	// if colorization somehow failed, do nothing
	if err != nil {
		return
	}

	// otherwise, swap the cheatsheet's Text with its colorized equivalent
	s.Text = buf.String()
}
