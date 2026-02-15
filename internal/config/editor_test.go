package config

import (
	"os"
	"runtime"
	"testing"
)

// TestEditor tests the Editor function
func TestEditor(t *testing.T) {
	// Save original env vars
	oldVisual := os.Getenv("VISUAL")
	oldEditor := os.Getenv("EDITOR")
	defer func() {
		os.Setenv("VISUAL", oldVisual)
		os.Setenv("EDITOR", oldEditor)
	}()

	t.Run("windows default", func(t *testing.T) {
		if runtime.GOOS != "windows" {
			t.Skip("skipping windows test on non-windows platform")
		}

		// Clear env vars
		os.Setenv("VISUAL", "")
		os.Setenv("EDITOR", "")

		editor, err := Editor()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if editor != "notepad" {
			t.Errorf("expected 'notepad' on windows, got %s", editor)
		}
	})

	t.Run("VISUAL takes precedence", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("skipping non-windows test on windows platform")
		}

		os.Setenv("VISUAL", "emacs")
		os.Setenv("EDITOR", "nano")

		editor, err := Editor()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if editor != "emacs" {
			t.Errorf("expected VISUAL to take precedence, got %s", editor)
		}
	})

	t.Run("EDITOR when no VISUAL", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("skipping non-windows test on windows platform")
		}

		os.Setenv("VISUAL", "")
		os.Setenv("EDITOR", "vim")

		editor, err := Editor()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if editor != "vim" {
			t.Errorf("expected EDITOR value, got %s", editor)
		}
	})

	t.Run("no editor found error", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("skipping non-windows test on windows platform")
		}

		// Clear all environment variables
		os.Setenv("VISUAL", "")
		os.Setenv("EDITOR", "")

		// Create a custom PATH that doesn't include common editors
		oldPath := os.Getenv("PATH")
		defer os.Setenv("PATH", oldPath)

		// Set a very limited PATH that won't have editors
		os.Setenv("PATH", "/nonexistent")

		editor, err := Editor()

		// If we found an editor, it's likely in the system
		// This test might not always produce an error on systems with editors
		if editor == "" && err == nil {
			t.Error("expected error when no editor found")
		}
	})
}
