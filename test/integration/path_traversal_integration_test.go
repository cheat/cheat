package integration

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestPathTraversalIntegration tests that the cheat binary properly blocks
// path traversal attempts when invoked as a subprocess.
func TestPathTraversalIntegration(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("integration test uses Unix-specific env and tools")
	}

	// Build the cheat binary
	binPath := filepath.Join(t.TempDir(), "cheat_test")
	build := exec.Command("go", "build", "-o", binPath, "./cmd/cheat")
	build.Dir = repoRoot(t)
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build cheat: %v\nOutput: %s", err, output)
	}

	// Set up test environment
	testDir := t.TempDir()
	sheetsDir := filepath.Join(testDir, "sheets")
	os.MkdirAll(sheetsDir, 0755)

	// Create config
	config := fmt.Sprintf(`---
editor: echo
colorize: false
pager: cat
cheatpaths:
  - name: test
    path: %s
    readonly: false
`, sheetsDir)

	configPath := filepath.Join(testDir, "config.yml")
	if err := os.WriteFile(configPath, []byte(config), 0644); err != nil {
		t.Fatalf("Failed to write config: %v", err)
	}

	// Test table
	tests := []struct {
		name     string
		command  []string
		wantFail bool
		wantMsg  string
	}{
		// Blocked patterns
		{
			name:     "block parent traversal edit",
			command:  []string{"--edit", "../evil"},
			wantFail: true,
			wantMsg:  "cannot contain '..'",
		},
		{
			name:     "block absolute path edit",
			command:  []string{"--edit", "/etc/passwd"},
			wantFail: true,
			wantMsg:  "cannot be an absolute path",
		},
		{
			name:     "block home dir edit",
			command:  []string{"--edit", "~/.ssh/config"},
			wantFail: true,
			wantMsg:  "cannot start with '~'",
		},
		{
			name:     "block parent traversal remove",
			command:  []string{"--rm", "../evil"},
			wantFail: true,
			wantMsg:  "cannot contain '..'",
		},
		{
			name:     "block complex traversal",
			command:  []string{"--edit", "foo/../../bar"},
			wantFail: true,
			wantMsg:  "cannot contain '..'",
		},
		{
			name:     "block just dots",
			command:  []string{"--edit", ".."},
			wantFail: true,
			wantMsg:  "cannot contain '..'",
		},
		{
			name:     "block empty name",
			command:  []string{"--edit", ""},
			wantFail: true,
			wantMsg:  "cannot be empty",
		},
		// Allowed patterns
		{
			name:     "allow simple name",
			command:  []string{"--edit", "docker"},
			wantFail: false,
		},
		{
			name:     "allow nested name",
			command:  []string{"--edit", "lang/go"},
			wantFail: false,
		},
		{
			name:     "block hidden file",
			command:  []string{"--edit", ".gitignore"},
			wantFail: true,
			wantMsg:  "cannot start with '.'",
		},
		{
			name:     "allow current dir",
			command:  []string{"--edit", "./local"},
			wantFail: false,
		},
	}

	// Run tests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command(binPath, tc.command...)
			cmd.Env = []string{
				fmt.Sprintf("CHEAT_CONFIG_PATH=%s", configPath),
				fmt.Sprintf("HOME=%s", testDir),
			}
			output, err := cmd.CombinedOutput()

			if tc.wantFail {
				if err == nil {
					t.Errorf("Expected failure but command succeeded. Output: %s", output)
				}
				if !strings.Contains(string(output), "invalid cheatsheet name") {
					t.Errorf("Expected 'invalid cheatsheet name' error, got: %s", output)
				}
				if tc.wantMsg != "" && !strings.Contains(string(output), tc.wantMsg) {
					t.Errorf("Expected message %q in output, got: %s", tc.wantMsg, output)
				}
			} else {
				// Command might fail for other reasons (e.g., editor not found)
				// but should NOT fail with "invalid cheatsheet name"
				if strings.Contains(string(output), "invalid cheatsheet name") {
					t.Errorf("Command incorrectly blocked. Output: %s", output)
				}
			}
		})
	}
}

// TestPathTraversalRealWorld tests with more realistic scenarios
func TestPathTraversalRealWorld(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("integration test uses Unix-specific env and tools")
	}

	// This test ensures our protection works with actual file operations

	// Build cheat
	binPath := filepath.Join(t.TempDir(), "cheat_test")
	build := exec.Command("go", "build", "-o", binPath, "./cmd/cheat")
	build.Dir = repoRoot(t)
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build: %v\n%s", err, output)
	}

	// Create test structure
	testRoot := t.TempDir()
	sheetsDir := filepath.Join(testRoot, "cheatsheets")
	secretDir := filepath.Join(testRoot, "secrets")
	os.MkdirAll(sheetsDir, 0755)
	os.MkdirAll(secretDir, 0755)

	// Create a "secret" file that should not be accessible
	secretFile := filepath.Join(secretDir, "secret.txt")
	os.WriteFile(secretFile, []byte("SECRET DATA"), 0644)

	// Create config using vim in non-interactive mode
	config := fmt.Sprintf(`---
editor: vim -u NONE -n --cmd "set noswapfile" --cmd "wq"
colorize: false
pager: cat
cheatpaths:
  - name: personal
    path: %s
    readonly: false
`, sheetsDir)

	configPath := filepath.Join(testRoot, "config.yml")
	os.WriteFile(configPath, []byte(config), 0644)

	// Test 1: Try to edit a file outside cheatsheets using traversal
	cmd := exec.Command(binPath, "--edit", "../secrets/secret")
	cmd.Env = []string{
		fmt.Sprintf("CHEAT_CONFIG_PATH=%s", configPath),
		fmt.Sprintf("HOME=%s", testRoot),
	}
	output, err := cmd.CombinedOutput()

	if err == nil || !strings.Contains(string(output), "invalid cheatsheet name") {
		t.Errorf("Path traversal was not blocked! Output: %s", output)
	}

	// Test 2: Verify the secret file is still intact
	content, _ := os.ReadFile(secretFile)
	if string(content) != "SECRET DATA" {
		t.Errorf("Secret file was modified!")
	}

	// Test 3: Verify no files were created outside sheets directory
	err = filepath.Walk(testRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() &&
			path != configPath &&
			path != secretFile &&
			!strings.HasPrefix(path, sheetsDir) {
			t.Errorf("File created outside allowed directory: %s", path)
		}
		return nil
	})
	if err != nil {
		t.Errorf("Walk error: %v", err)
	}
}
