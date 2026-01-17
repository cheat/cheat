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
	def, defErr := exec.LookPath("editor") // default `editor` wrapper
	if defErr != nil {
		def = "" // Reset to empty string if not found
	}
	nano, nanoErr := exec.LookPath("nano")
	if nanoErr != nil {
		nano = "" // Reset to empty string if not found
	}
	vim, vimErr := exec.LookPath("vim")
	if vimErr != nil {
		vim = "" // Reset to empty string if not found
	}

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
	return "", fmt.Errorf("no editor found: please set the EDITOR or VISUAL environment variable, or ensure 'nano' or 'vim' are in your PATH")
}
