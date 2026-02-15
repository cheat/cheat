package cheatpath

import (
	"fmt"
	"path/filepath"
	"strings"
)

// ValidateSheetName ensures that a cheatsheet name does not contain
// directory traversal sequences or other potentially dangerous patterns.
func ValidateSheetName(name string) error {
	// Reject empty names
	if name == "" {
		return fmt.Errorf("cheatsheet name cannot be empty")
	}

	// Reject names containing directory traversal
	if strings.Contains(name, "..") {
		return fmt.Errorf("cheatsheet name cannot contain '..'")
	}

	// Reject absolute paths
	if filepath.IsAbs(name) {
		return fmt.Errorf("cheatsheet name cannot be an absolute path")
	}

	// Reject names that start with ~ (home directory expansion)
	if strings.HasPrefix(name, "~") {
		return fmt.Errorf("cheatsheet name cannot start with '~'")
	}

	// Reject hidden files (files that start with a dot)
	// We don't display hidden files, so we shouldn't create them
	filename := filepath.Base(name)
	if strings.HasPrefix(filename, ".") {
		return fmt.Errorf("cheatsheet name cannot start with '.' (hidden files are not supported)")
	}

	return nil
}
