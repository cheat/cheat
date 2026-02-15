package repo

import (
	"os"
	"path/filepath"
	"testing"
)

// TestClone tests the Clone function
func TestClone(t *testing.T) {
	// This test requires network access, so we'll only test error cases
	// that don't require actual cloning

	t.Run("clone to read-only directory", func(t *testing.T) {
		if os.Getuid() == 0 {
			t.Skip("Cannot test read-only directory as root")
		}

		// Create a temporary directory
		tempDir, err := os.MkdirTemp("", "cheat-clone-test-*")
		if err != nil {
			t.Fatalf("failed to create temp dir: %v", err)
		}
		defer os.RemoveAll(tempDir)

		// Create a read-only subdirectory
		readOnlyDir := filepath.Join(tempDir, "readonly")
		if err := os.Mkdir(readOnlyDir, 0555); err != nil {
			t.Fatalf("failed to create read-only dir: %v", err)
		}

		// Attempt to clone to read-only directory
		targetDir := filepath.Join(readOnlyDir, "cheatsheets")
		err = Clone(targetDir)

		// Should fail because we can't write to read-only directory
		if err == nil {
			t.Error("expected error when cloning to read-only directory, got nil")
		}
	})

	t.Run("clone to invalid path", func(t *testing.T) {
		// Try to clone to a path with null bytes (invalid on most filesystems)
		err := Clone("/tmp/invalid\x00path")
		if err == nil {
			t.Error("expected error with invalid path, got nil")
		}
	})
}
