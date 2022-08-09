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

	// look for `nano` and `vim` on the `PATH`
	def, _ := exec.LookPath("editor") // default `editor` wrapper
	nano, _ := exec.LookPath("nano")
	vim, _ := exec.LookPath("vim")

	// set editor priority
	editors := []string{
		os.Getenv("VISUAL"),
		os.Getenv("EDITOR"),
		def,
		nano,
		vim,
	}

	// return the first editor that was found per the priority above
	for _, editor := range editors {
		if editor != "" {
			return editor, nil
		}
	}

	// return an error if no path is found
	return "", fmt.Errorf("no editor set")
}
