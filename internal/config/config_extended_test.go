package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cheat/cheat/mocks"
)

// TestConfigYAMLErrors tests YAML parsing errors
func TestConfigYAMLErrors(t *testing.T) {
	// Create a temporary file with invalid YAML
	tempDir, err := os.MkdirTemp("", "cheat-config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	invalidYAML := filepath.Join(tempDir, "invalid.yml")
	err = os.WriteFile(invalidYAML, []byte("cheatpaths: [{unclosed\n"), 0644)
	if err != nil {
		t.Fatalf("failed to write invalid yaml: %v", err)
	}

	// Attempt to load invalid YAML
	_, err = New(invalidYAML, false)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

// TestConfigDefaults tests default values
func TestConfigDefaults(t *testing.T) {
	// Load empty config
	conf, err := New(mocks.Path("conf/empty.yml"), false)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	// Check defaults
	if conf.Style != "bw" {
		t.Errorf("expected default style 'bw', got %s", conf.Style)
	}

	if conf.Formatter != "terminal" {
		t.Errorf("expected default formatter 'terminal', got %s", conf.Formatter)
	}
}

// TestConfigSymlinkResolution tests symlink resolution
func TestConfigSymlinkResolution(t *testing.T) {
	// Create temp directory structure
	tempDir, err := os.MkdirTemp("", "cheat-config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Resolve symlinks in temp dir path (macOS /var -> /private/var)
	tempDir, err = filepath.EvalSymlinks(tempDir)
	if err != nil {
		t.Fatalf("failed to resolve temp dir symlinks: %v", err)
	}

	// Create target directory
	targetDir := filepath.Join(tempDir, "target")
	err = os.Mkdir(targetDir, 0755)
	if err != nil {
		t.Fatalf("failed to create target dir: %v", err)
	}

	// Create symlink
	linkPath := filepath.Join(tempDir, "link")
	err = os.Symlink(targetDir, linkPath)
	if err != nil {
		t.Fatalf("failed to create symlink: %v", err)
	}

	// Create config with symlink path
	configContent := `---
editor: vim
cheatpaths:
  - name: test
    path: ` + linkPath + `
    readonly: true
`
	configFile := filepath.Join(tempDir, "config.yml")
	err = os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	// Load config with symlink resolution
	conf, err := New(configFile, true)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	// Verify symlink was resolved
	if len(conf.Cheatpaths) == 0 {
		t.Fatal("expected at least one cheatpath, got none")
	}
	if conf.Cheatpaths[0].Path != targetDir {
		t.Errorf("expected symlink to be resolved to %s, got %s", targetDir, conf.Cheatpaths[0].Path)
	}
}

// TestConfigBrokenSymlink tests broken symlink handling
func TestConfigBrokenSymlink(t *testing.T) {
	// Create temp directory
	tempDir, err := os.MkdirTemp("", "cheat-config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create broken symlink
	linkPath := filepath.Join(tempDir, "broken-link")
	err = os.Symlink("/nonexistent/path", linkPath)
	if err != nil {
		t.Fatalf("failed to create symlink: %v", err)
	}

	// Create config with broken symlink
	configContent := `---
editor: vim
cheatpaths:
  - name: test
    path: ` + linkPath + `
    readonly: true
`
	configFile := filepath.Join(tempDir, "config.yml")
	err = os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	// Load config with symlink resolution should skip the broken cheatpath
	// (warn to stderr) rather than hard-error
	conf, err := New(configFile, true)
	if err != nil {
		t.Errorf("expected no error for broken symlink (should skip), got: %v", err)
	}
	if len(conf.Cheatpaths) != 0 {
		t.Errorf("expected broken cheatpath to be filtered out, got %d cheatpaths", len(conf.Cheatpaths))
	}
}
