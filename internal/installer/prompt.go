// Package installer implements functions that provide a first-time
// installation wizard.
package installer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Prompt prompts the user for a answer
func Prompt(prompt string, def bool) (bool, error) {

	// initialize a line reader
	reader := bufio.NewReader(os.Stdin)

	// display the prompt
	fmt.Printf("%s: ", prompt)

	// read the answer
	ans, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to parse input: %v", err)
	}

	// normalize the answer
	ans = strings.ToLower(strings.TrimSpace(ans))

	// return the appropriate response
	switch ans {
	case "y":
		return true, nil
	case "":
		return def, nil
	default:
		return false, nil
	}
}
