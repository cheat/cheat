package display

import "fmt"

// Faint returns an faint string
func Faint(str string) string {
	return fmt.Sprintf(fmt.Sprintf("\033[2m%s\033[0m", str))
}
