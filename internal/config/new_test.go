package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewTrimsWhitespace(t *testing.T) {
	// Create a temporary config file with whitespace in editor and pager
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yml")

	configContent := `---
editor: "  vim -c 'set number'  "
pager: "  less -R  "
style: monokai
formatter: terminal
cheatpaths:
  - name: personal
    path: ~/cheat
    tags: []
    readonly: false
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	// Load the config
	conf, err := New(map[string]interface{}{}, configPath, false)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Verify editor is trimmed
	expectedEditor := "vim -c 'set number'"
	if conf.Editor != expectedEditor {
		t.Errorf("editor not properly trimmed: got %q, want %q", conf.Editor, expectedEditor)
	}

	// Verify pager is trimmed
	expectedPager := "less -R"
	if conf.Pager != expectedPager {
		t.Errorf("pager not properly trimmed: got %q, want %q", conf.Pager, expectedPager)
	}
}

func TestNewEmptyEditorFallback(t *testing.T) {
	// Skip if required environment variables would interfere
	oldVisual := os.Getenv("VISUAL")
	oldEditor := os.Getenv("EDITOR")
	os.Unsetenv("VISUAL")
	os.Unsetenv("EDITOR")
	defer func() {
		os.Setenv("VISUAL", oldVisual)
		os.Setenv("EDITOR", oldEditor)
	}()

	// Create a config with whitespace-only editor
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yml")

	configContent := `---
editor: "   "
pager: less
style: monokai
formatter: terminal
cheatpaths:
  - name: personal
    path: ~/cheat
    tags: []
    readonly: false
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	// Load the config
	conf, err := New(map[string]interface{}{}, configPath, false)
	if err != nil {
		// It's OK if this fails due to no editor being found
		// The important thing is it doesn't panic
		return
	}

	// If it succeeded, editor should not be empty (fallback was used)
	if conf.Editor == "" {
		t.Error("editor should not be empty after fallback")
	}
}

func TestNewWhitespaceOnlyPager(t *testing.T) {
	// Create a config with whitespace-only pager
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yml")

	configContent := `---
editor: vim
pager: "   "
style: monokai
formatter: terminal
cheatpaths:
  - name: personal
    path: ~/cheat
    tags: []
    readonly: false
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	// Load the config
	conf, err := New(map[string]interface{}{}, configPath, false)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Pager should be empty after trimming
	if conf.Pager != "" {
		t.Errorf("pager should be empty after trimming whitespace: got %q", conf.Pager)
	}
}
