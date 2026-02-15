package installer

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "cheat-installer-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Save original stdin/stdout
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	tests := []struct {
		name          string
		configs       string
		confpath      string
		userInput     string
		wantErr       bool
		wantInErr     string
		checkFiles    []string
		dontWantFiles []string
	}{
		{
			name: "user declines community cheatsheets",
			configs: `---
editor: EDITOR_PATH
pager: PAGER_PATH
cheatpaths:
  - name: community
    path: COMMUNITY_PATH
    tags: [ community ]
    readonly: true
  - name: personal
    path: PERSONAL_PATH
    tags: [ personal ]
    readonly: false
`,
			confpath:      filepath.Join(tempDir, "conf1", "conf.yml"),
			userInput:     "n\n",
			wantErr:       false,
			checkFiles:    []string{"conf1/conf.yml"},
			dontWantFiles: []string{"conf1/cheatsheets/community", "conf1/cheatsheets/personal"},
		},
		{
			name: "user accepts but clone fails",
			configs: `---
cheatpaths:
  - name: community
    path: COMMUNITY_PATH
`,
			confpath:  filepath.Join(tempDir, "conf2", "conf.yml"),
			userInput: "y\n",
			wantErr:   true,
			wantInErr: "failed to clone cheatsheets",
		},
		{
			name:    "invalid config path",
			configs: "test",
			// /dev/null/... is truly uncreatable on Unix;
			// NUL\... is uncreatable on Windows
			confpath: func() string {
				if runtime.GOOS == "windows" {
					return `NUL\impossible\conf.yml`
				}
				return "/dev/null/impossible/conf.yml"
			}(),
			userInput: "n\n",
			wantErr:   true,
			wantInErr: "failed to create",
		},
	}

	// Pre-create a .git dir inside the community path so go-git's PlainClone
	// returns ErrRepositoryAlreadyExists (otherwise, on CI runners with
	// network access, the real clone succeeds and the test fails)
	fakeGitDir := filepath.Join(tempDir, "conf2", "cheatsheets", "community", ".git")
	if err := os.MkdirAll(fakeGitDir, 0755); err != nil {
		t.Fatalf("failed to create fake .git dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(fakeGitDir, "HEAD"), []byte("ref: refs/heads/main\n"), 0644); err != nil {
		t.Fatalf("failed to write fake HEAD: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create stdin pipe
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Create stdout pipe to suppress output
			_, wOut, _ := os.Pipe()
			os.Stdout = wOut

			// Write user input
			go func() {
				defer w.Close()
				io.WriteString(w, tt.userInput)
			}()

			// Run the installer
			err := Run(tt.configs, tt.confpath)

			// Close pipes
			wOut.Close()

			// Check error
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.wantInErr != "" && !strings.Contains(err.Error(), tt.wantInErr) {
				t.Errorf("Run() error = %v, want error containing %q", err, tt.wantInErr)
			}

			// Check created files
			for _, file := range tt.checkFiles {
				path := filepath.Join(tempDir, file)
				if _, err := os.Stat(path); os.IsNotExist(err) {
					t.Errorf("expected file %s to exist, but it doesn't", path)
				}
			}

			// Check files that shouldn't exist
			for _, file := range tt.dontWantFiles {
				path := filepath.Join(tempDir, file)
				if _, err := os.Stat(path); err == nil {
					t.Errorf("expected file %s to not exist, but it does", path)
				}
			}
		})
	}
}

func TestRunPromptError(t *testing.T) {
	// Save original stdin
	oldStdin := os.Stdin
	defer func() {
		os.Stdin = oldStdin
	}()

	// Close stdin to cause prompt error
	r, w, _ := os.Pipe()
	os.Stdin = r
	r.Close()
	w.Close()

	tempDir, _ := os.MkdirTemp("", "cheat-installer-prompt-test-*")
	defer os.RemoveAll(tempDir)

	err := Run("test", filepath.Join(tempDir, "conf.yml"))
	if err == nil {
		t.Error("expected error when prompt fails, got nil")
	}
	if !strings.Contains(err.Error(), "failed to prompt") {
		t.Errorf("expected 'failed to prompt' error, got: %v", err)
	}
}

func TestRunStringReplacements(t *testing.T) {
	// Test that path replacements work correctly
	configs := `---
editor: EDITOR_PATH
pager: PAGER_PATH
cheatpaths:
  - name: community
    path: COMMUNITY_PATH
  - name: personal
    path: PERSONAL_PATH
`

	// Create temp directory
	tempDir, err := os.MkdirTemp("", "cheat-installer-replace-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	confpath := filepath.Join(tempDir, "conf.yml")
	confdir := filepath.Dir(confpath)

	// Expected paths
	expectedCommunity := filepath.Join(confdir, "cheatsheets", "community")
	expectedPersonal := filepath.Join(confdir, "cheatsheets", "personal")

	// Save original stdin/stdout
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// Create stdin pipe with "n" answer
	r, w, _ := os.Pipe()
	os.Stdin = r
	go func() {
		defer w.Close()
		io.WriteString(w, "n\n")
	}()

	// Suppress stdout
	_, wOut, _ := os.Pipe()
	os.Stdout = wOut
	defer wOut.Close()

	// Run installer
	err = Run(configs, confpath)
	if err != nil {
		t.Fatalf("Run() failed: %v", err)
	}

	// Read the created config file
	content, err := os.ReadFile(confpath)
	if err != nil {
		t.Fatalf("failed to read config file: %v", err)
	}

	// Check replacements
	contentStr := string(content)
	if strings.Contains(contentStr, "COMMUNITY_PATH") {
		t.Error("COMMUNITY_PATH was not replaced")
	}
	if strings.Contains(contentStr, "PERSONAL_PATH") {
		t.Error("PERSONAL_PATH was not replaced")
	}
	if strings.Contains(contentStr, "EDITOR_PATH") && !strings.Contains(contentStr, fmt.Sprintf("editor: %s", "")) {
		t.Error("EDITOR_PATH was not replaced")
	}
	if strings.Contains(contentStr, "PAGER_PATH") && !strings.Contains(contentStr, fmt.Sprintf("pager: %s", "")) {
		t.Error("PAGER_PATH was not replaced")
	}

	// Verify correct paths were used
	if !strings.Contains(contentStr, expectedCommunity) {
		t.Errorf("expected community path %q in config", expectedCommunity)
	}
	if !strings.Contains(contentStr, expectedPersonal) {
		t.Errorf("expected personal path %q in config", expectedPersonal)
	}
}
