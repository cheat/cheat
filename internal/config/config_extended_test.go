package config

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cheat/cheat/internal/mock"
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
	err = os.WriteFile(invalidYAML, []byte("invalid: yaml: content:\n  - no closing"), 0644)
	if err != nil {
		t.Fatalf("failed to write invalid yaml: %v", err)
	}

	// Attempt to load invalid YAML
	_, err = New(map[string]interface{}{}, invalidYAML, false)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

// TestConfigLocalCheatpath tests local .cheat directory detection
func TestConfigLocalCheatpath(t *testing.T) {
	// Create a temporary directory to act as working directory
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

	// Save current working directory
	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	defer os.Chdir(oldCwd)

	// Change to temp directory
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("failed to change dir: %v", err)
	}

	// Create .cheat directory
	localCheat := filepath.Join(tempDir, ".cheat")
	err = os.Mkdir(localCheat, 0755)
	if err != nil {
		t.Fatalf("failed to create .cheat dir: %v", err)
	}

	// Load config
	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	// Check that local cheatpath was added
	found := false
	for _, cp := range conf.Cheatpaths {
		if cp.Name == "cwd" && cp.Path == localCheat {
			found = true
			break
		}
	}

	if !found {
		t.Error("local .cheat directory was not added to cheatpaths")
	}
}

// TestConfigLocalCheatpathInParent tests that .cheat in a parent directory is found
func TestConfigLocalCheatpathInParent(t *testing.T) {
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

	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	defer os.Chdir(oldCwd)

	// Create .cheat in the root of the temp dir
	localCheat := filepath.Join(tempDir, ".cheat")
	if err := os.Mkdir(localCheat, 0755); err != nil {
		t.Fatalf("failed to create .cheat dir: %v", err)
	}

	// Create a subdirectory and cd into it
	subDir := filepath.Join(tempDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create sub dir: %v", err)
	}
	if err := os.Chdir(subDir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	found := false
	for _, cp := range conf.Cheatpaths {
		if cp.Name == "cwd" && cp.Path == localCheat {
			found = true
			break
		}
	}
	if !found {
		t.Error("parent .cheat directory was not added to cheatpaths")
	}
}

// TestConfigLocalCheatpathNearestWins tests that the nearest .cheat wins
func TestConfigLocalCheatpathNearestWins(t *testing.T) {
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

	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	defer os.Chdir(oldCwd)

	// Create .cheat at root
	if err := os.Mkdir(filepath.Join(tempDir, ".cheat"), 0755); err != nil {
		t.Fatalf("failed to create root .cheat dir: %v", err)
	}

	// Create sub/.cheat (the nearer one)
	subDir := filepath.Join(tempDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create sub dir: %v", err)
	}
	nearCheat := filepath.Join(subDir, ".cheat")
	if err := os.Mkdir(nearCheat, 0755); err != nil {
		t.Fatalf("failed to create near .cheat dir: %v", err)
	}

	// cd into sub/deep/
	deepDir := filepath.Join(subDir, "deep")
	if err := os.Mkdir(deepDir, 0755); err != nil {
		t.Fatalf("failed to create deep dir: %v", err)
	}
	if err := os.Chdir(deepDir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	found := false
	for _, cp := range conf.Cheatpaths {
		if cp.Name == "cwd" {
			if cp.Path != nearCheat {
				t.Errorf("expected nearest .cheat %s, got %s", nearCheat, cp.Path)
			}
			found = true
			break
		}
	}
	if !found {
		t.Error("no cwd cheatpath found")
	}
}

// TestConfigNoLocalCheatpath tests that no cwd cheatpath is added when no .cheat exists
func TestConfigNoLocalCheatpath(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	defer os.Chdir(oldCwd)

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	for _, cp := range conf.Cheatpaths {
		if cp.Name == "cwd" {
			t.Error("cwd cheatpath should not be added when no .cheat exists")
		}
	}
}

// TestConfigLocalCheatpathFileSkipped tests that a .cheat file (not dir) is skipped
func TestConfigLocalCheatpathFileSkipped(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	defer os.Chdir(oldCwd)

	// Create .cheat as a file, not a directory
	if err := os.WriteFile(filepath.Join(tempDir, ".cheat"), []byte("not a dir"), 0644); err != nil {
		t.Fatalf("failed to create .cheat file: %v", err)
	}

	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	for _, cp := range conf.Cheatpaths {
		if cp.Name == "cwd" {
			t.Error("cwd cheatpath should not be added for a .cheat file")
		}
	}
}

// TestConfigDefaults tests default values
func TestConfigDefaults(t *testing.T) {
	// Load empty config
	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
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
	conf, err := New(map[string]interface{}{}, configFile, true)
	if err != nil {
		t.Errorf("failed to load config: %v", err)
	}

	// Verify symlink was resolved
	if len(conf.Cheatpaths) > 0 && conf.Cheatpaths[0].Path != targetDir {
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
	conf, err := New(map[string]interface{}{}, configFile, true)
	if err != nil {
		t.Errorf("expected no error for broken symlink (should skip), got: %v", err)
	}
	if len(conf.Cheatpaths) != 0 {
		t.Errorf("expected broken cheatpath to be filtered out, got %d cheatpaths", len(conf.Cheatpaths))
	}
}

// TestConfigTildeExpansionError tests tilde expansion error handling
func TestConfigTildeExpansionError(t *testing.T) {
	// This is tricky to test without mocking homedir.Expand
	// We'll create a config with an invalid home reference
	tempDir, err := os.MkdirTemp("", "cheat-config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create config with user that likely doesn't exist
	configContent := `---
editor: vim
cheatpaths:
  - name: test
    path: ~nonexistentuser12345/cheat
    readonly: true
`
	configFile := filepath.Join(tempDir, "config.yml")
	err = os.WriteFile(configFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to write config: %v", err)
	}

	// Load config - this may or may not fail depending on the system
	// but we're testing that it doesn't panic
	_, _ = New(map[string]interface{}{}, configFile, false)
}

// TestConfigGetCwdError tests error handling when os.Getwd fails
func TestConfigGetCwdError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Windows does not allow removing the current directory")
	}

	// This is difficult to test without being able to break os.Getwd
	// We'll create a scenario where the current directory is removed

	// Create and enter a temp directory
	tempDir, err := os.MkdirTemp("", "cheat-config-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	defer os.Chdir(oldCwd)

	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("failed to change dir: %v", err)
	}

	// Remove the directory we're in
	err = os.RemoveAll(tempDir)
	if err != nil {
		t.Fatalf("failed to remove temp dir: %v", err)
	}

	// Now os.Getwd should fail
	_, err = New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	// This might not fail on all systems, so we just ensure no panic
	_ = err
}
