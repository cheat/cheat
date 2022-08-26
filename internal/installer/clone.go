package installer

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// clone clones the community cheatsheets
func clone(path string) error {

	// clone the community cheatsheets
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/cheat/cheatsheets.git",
		Depth:    1,
		Progress: os.Stdout,
	})

	if err != nil {
		return fmt.Errorf("failed to clone cheatsheets: %v", err)
	}

	return nil
}
