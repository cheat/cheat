//go:build integration
// +build integration

package repo

import (
	"os"
	"path/filepath"
	"testing"
)

// TestCloneIntegration performs a real clone operation to verify functionality
// Run with: go test -tags=integration ./internal/repo -v -run TestCloneIntegration
func TestCloneIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "cheat-clone-integration-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	destDir := filepath.Join(tmpDir, "cheatsheets")

	t.Logf("Cloning to: %s", destDir)

	// Perform the actual clone
	err = Clone(destDir)
	if err != nil {
		t.Fatalf("Clone() failed: %v", err)
	}

	// Verify the clone succeeded
	info, err := os.Stat(destDir)
	if err != nil {
		t.Fatalf("destination directory not created: %v", err)
	}

	if !info.IsDir() {
		t.Fatal("destination is not a directory")
	}

	// Check for .git directory
	gitDir := filepath.Join(destDir, ".git")
	if _, err := os.Stat(gitDir); err != nil {
		t.Error(".git directory not found")
	}

	// Check for some expected cheatsheets
	expectedFiles := []string{
		"bash", // bash cheatsheet should exist
		"git",  // git cheatsheet should exist
		"ls",   // ls cheatsheet should exist
	}

	foundCount := 0
	for _, file := range expectedFiles {
		path := filepath.Join(destDir, file)
		if _, err := os.Stat(path); err == nil {
			foundCount++
		}
	}

	if foundCount < 2 {
		t.Errorf("expected at least 2 common cheatsheets, found %d", foundCount)
	}

	t.Log("Clone integration test passed!")

	// Test cloning to existing directory (should fail)
	err = Clone(destDir)
	if err == nil {
		t.Error("expected error when cloning to existing repository, got nil")
	} else {
		t.Logf("Expected error when cloning to existing dir: %v", err)
	}
}
