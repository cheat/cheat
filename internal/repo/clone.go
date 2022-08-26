package repo

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// Clone clones the repo available at `url`
func Clone(url string) error {

	// clone the community cheatsheets
	_, err := git.PlainClone(url, false, &git.CloneOptions{
		URL:      "https://github.com/cheat/cheatsheets.git",
		Depth:    1,
		Progress: os.Stdout,
	})

	if err != nil {
		return fmt.Errorf("failed to clone cheatsheets: %v", err)
	}

	return nil
}
