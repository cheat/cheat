package installer

import (
	"fmt"
	"os"
	"os/exec"
)

const cloneURL = "https://github.com/cheat/cheatsheets.git"

// clone clones the community cheatsheets
func clone(path string) error {

	// perform the clone in a shell
	cmd := exec.Command("git", "clone", cloneURL, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to clone cheatsheets: %v", err)
	}

	return nil
}
