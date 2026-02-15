//go:build integration

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// BenchmarkSearchCommand benchmarks the actual cheat search command
func BenchmarkSearchCommand(b *testing.B) {
	// Build the cheat binary in .tmp (using absolute path)
	rootDir, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		b.Fatalf("Failed to get root dir: %v", err)
	}
	tmpDir := filepath.Join(rootDir, ".tmp", "bench-test")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}

	cheatBin := filepath.Join(tmpDir, "cheat-bench")

	// Clean up the binary when done
	b.Cleanup(func() {
		os.Remove(cheatBin)
	})

	cmd := exec.Command("go", "build", "-o", cheatBin, "./cmd/cheat")
	cmd.Dir = rootDir
	if output, err := cmd.CombinedOutput(); err != nil {
		b.Fatalf("Failed to build cheat: %v\nOutput: %s", err, output)
	}

	// Set up test environment in .tmp
	configDir := filepath.Join(tmpDir, "config")
	cheatsheetDir := filepath.Join(configDir, "cheatsheets", "community")

	// Clone community cheatsheets (or reuse if already exists)
	if _, err := os.Stat(cheatsheetDir); os.IsNotExist(err) {
		b.Logf("Cloning community cheatsheets to %s...", cheatsheetDir)
		_, err := git.PlainClone(cheatsheetDir, false, &git.CloneOptions{
			URL:           "https://github.com/cheat/cheatsheets.git",
			Depth:         1,
			SingleBranch:  true,
			ReferenceName: plumbing.ReferenceName("refs/heads/master"),
			Progress:      nil,
		})
		if err != nil {
			b.Fatalf("Failed to clone cheatsheets: %v", err)
		}
	}

	// Create a minimal config file
	configFile := filepath.Join(configDir, "conf.yml")
	configContent := fmt.Sprintf(`---
cheatpaths:
  - name: community
    path: %s
    tags: [ community ]
    readonly: true
`, cheatsheetDir)

	if err := os.MkdirAll(configDir, 0755); err != nil {
		b.Fatalf("Failed to create config dir: %v", err)
	}
	if err := os.WriteFile(configFile, []byte(configContent), 0644); err != nil {
		b.Fatalf("Failed to write config: %v", err)
	}

	// Set environment to use our config
	env := append(os.Environ(),
		fmt.Sprintf("CHEAT_CONFIG_PATH=%s", configFile),
	)

	// Define test cases
	testCases := []struct {
		name string
		args []string
	}{
		{"SimpleSearch", []string{"-s", "echo"}},
		{"RegexSearch", []string{"-r", "-s", "^#.*example"}},
		{"ColorizedSearch", []string{"-c", "-s", "grep"}},
		{"ComplexRegex", []string{"-r", "-s", "(git|hg|svn)\\s+(add|commit|push)"}},
		{"AllCheatpaths", []string{"-a", "-s", "list"}},
	}

	// Warm up - run once to ensure everything is loaded
	warmupCmd := exec.Command(cheatBin, "-l")
	warmupCmd.Env = env
	warmupCmd.Run()

	// Run benchmarks
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Reset timer to exclude setup
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				cmd := exec.Command(cheatBin, tc.args...)
				cmd.Env = env

				var stdout, stderr bytes.Buffer
				cmd.Stdout = &stdout
				cmd.Stderr = &stderr

				err := cmd.Run()
				if err != nil {
					b.Fatalf("Command failed: %v\nStderr: %s", err, stderr.String())
				}

				if stdout.Len() == 0 {
					b.Fatal("No output from search")
				}
			}
		})
	}
}

// BenchmarkListCommand benchmarks the list command for comparison
func BenchmarkListCommand(b *testing.B) {
	// Build the cheat binary in .tmp (using absolute path)
	rootDir, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		b.Fatalf("Failed to get root dir: %v", err)
	}
	tmpDir := filepath.Join(rootDir, ".tmp", "bench-test")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		b.Fatalf("Failed to create temp dir: %v", err)
	}

	cheatBin := filepath.Join(tmpDir, "cheat-bench")

	// Clean up the binary when done
	b.Cleanup(func() {
		os.Remove(cheatBin)
	})

	cmd := exec.Command("go", "build", "-o", cheatBin, "./cmd/cheat")
	cmd.Dir = rootDir
	if output, err := cmd.CombinedOutput(); err != nil {
		b.Fatalf("Failed to build cheat: %v\nOutput: %s", err, output)
	}

	// Set up test environment (simplified - reuse if possible)
	configDir := filepath.Join(tmpDir, "config")
	cheatsheetDir := filepath.Join(configDir, "cheatsheets", "community")

	// Check if we need to clone
	if _, err := os.Stat(cheatsheetDir); os.IsNotExist(err) {
		_, err := git.PlainClone(cheatsheetDir, false, &git.CloneOptions{
			URL:           "https://github.com/cheat/cheatsheets.git",
			Depth:         1,
			SingleBranch:  true,
			ReferenceName: plumbing.ReferenceName("refs/heads/master"),
			Progress:      nil,
		})
		if err != nil {
			b.Fatalf("Failed to clone cheatsheets: %v", err)
		}
	}

	// Create config
	configFile := filepath.Join(configDir, "conf.yml")
	configContent := fmt.Sprintf(`---
cheatpaths:
  - name: community
    path: %s
    tags: [ community ]
    readonly: true
`, cheatsheetDir)

	os.MkdirAll(configDir, 0755)
	os.WriteFile(configFile, []byte(configContent), 0644)

	env := append(os.Environ(),
		fmt.Sprintf("CHEAT_CONFIG_PATH=%s", configFile),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cmd := exec.Command(cheatBin, "-l")
		cmd.Env = env

		var stdout bytes.Buffer
		cmd.Stdout = &stdout

		if err := cmd.Run(); err != nil {
			b.Fatalf("Command failed: %v", err)
		}
	}
}
