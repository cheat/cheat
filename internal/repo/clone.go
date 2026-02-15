// Package repo implements functions pertaining to the management of git repos.
package repo

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

// Clone clones the community cheatsheets repository to the specified directory
func Clone(dir string) error {

	// clone the community cheatsheets
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      "https://github.com/cheat/cheatsheets.git",
		Depth:    1,
		Progress: os.Stdout,
	})

	if err != nil {
		return fmt.Errorf("failed to clone cheatsheets: %v", err)
	}

	return nil
}
