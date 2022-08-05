package display

import (
	"fmt"

	"github.com/cheat/cheat/internal/config"
)

// Faint returns an faint string
func Faint(str string, conf config.Config) string {
	// make `str` faint only if colorization has been requested
	if conf.Colorize {
		return fmt.Sprintf("\033[2m%s\033[0m", str)
	}

	// otherwise, return the string unmodified
	return str
}
