package display

import "fmt"

// Underline returns an underlined string
func Underline(str string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", str)
}
