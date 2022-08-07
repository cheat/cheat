package config

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Editor attempts to locate an editor that's appropriate for the environment.
func Editor() (string, error) {

	// default to `notepad.exe` on Windows
	if runtime.GOOS == "windows" {
		return "notepad", nil
	}

	// look for `nano` on the `PATH`
	nano, _ := exec.LookPath("nano")

	// search for `$VISUAL`, `$EDITOR`, and then `nano`, in that order
	for _, editor := range []string{os.Getenv("VISUAL"), os.Getenv("EDITOR"), nano} {
		if editor != "" {
			return editor, nil
		}
	}

	// return an error if no path is found
	return "", fmt.Errorf("no editor set")
}
