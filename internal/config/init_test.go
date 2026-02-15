package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestInit asserts that configs are properly initialized
func TestInit(t *testing.T) {

	// initialize a temporary config file
	confFile, err := os.CreateTemp("", "cheat-test")
	if err != nil {
		t.Errorf("failed to create temp file: %v", err)
	}

	// clean up the temp file
	defer os.Remove(confFile.Name())

	// initialize the config file
	conf := "mock config data"
	if err = Init(confFile.Name(), conf); err != nil {
		t.Errorf("failed to init config file: %v", err)
	}

	// read back the config file contents
	bytes, err := os.ReadFile(confFile.Name())
	if err != nil {
		t.Errorf("failed to read config file: %v", err)
	}

	// assert that the contents were written correctly
	got := string(bytes)
	if got != conf {
		t.Errorf("failed to write configs: want: %s, got: %s", conf, got)
	}
}

// TestInitCreateDirectory tests that Init creates the directory if it doesn't exist
func TestInitCreateDirectory(t *testing.T) {
	// Create a temp directory
	tempDir, err := os.MkdirTemp("", "cheat-init-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Path to a config file in a non-existent subdirectory
	confPath := filepath.Join(tempDir, "subdir", "conf.yml")

	// Initialize the config file
	conf := "test config"
	if err = Init(confPath, conf); err != nil {
		t.Errorf("failed to init config file: %v", err)
	}

	// Verify the directory was created
	if _, err := os.Stat(filepath.Dir(confPath)); os.IsNotExist(err) {
		t.Error("Init did not create the directory")
	}

	// Verify the file was created with correct content
	bytes, err := os.ReadFile(confPath)
	if err != nil {
		t.Errorf("failed to read config file: %v", err)
	}
	if string(bytes) != conf {
		t.Errorf("config content mismatch: got %q, want %q", string(bytes), conf)
	}
}

// TestInitWriteError tests error handling when file write fails
func TestInitWriteError(t *testing.T) {
	// Skip this test if running as root (can write anywhere)
	if os.Getuid() == 0 {
		t.Skip("Cannot test write errors as root")
	}

	// Try to write to a read-only directory
	err := Init("/dev/null/impossible/path/conf.yml", "test")
	if err == nil {
		t.Error("expected error when writing to invalid path, got nil")
	}
	if err != nil && !strings.Contains(err.Error(), "failed to create") {
		t.Errorf("expected 'failed to create' error, got: %v", err)
	}
}

// TestInitExistingFile tests that Init overwrites existing files
func TestInitExistingFile(t *testing.T) {
	// Create a temp file
	tempFile, err := os.CreateTemp("", "cheat-init-existing-*")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write initial content
	initialContent := "initial content"
	if err := os.WriteFile(tempFile.Name(), []byte(initialContent), 0644); err != nil {
		t.Fatalf("failed to write initial content: %v", err)
	}

	// Initialize with new content
	newContent := "new config content"
	if err = Init(tempFile.Name(), newContent); err != nil {
		t.Errorf("failed to init over existing file: %v", err)
	}

	// Verify the file was overwritten
	bytes, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Errorf("failed to read config file: %v", err)
	}
	if string(bytes) != newContent {
		t.Errorf("config not overwritten: got %q, want %q", string(bytes), newContent)
	}
}
