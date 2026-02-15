package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGitDir(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "cheat-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test directory structure
	testDirs := []string{
		filepath.Join(tempDir, ".git"),
		filepath.Join(tempDir, ".git", "objects"),
		filepath.Join(tempDir, ".git", "refs"),
		filepath.Join(tempDir, "regular"),
		filepath.Join(tempDir, "regular", ".git"),
		filepath.Join(tempDir, "submodule"),
	}

	for _, dir := range testDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			t.Fatalf("failed to create dir %s: %v", dir, err)
		}
	}

	// Create test files
	testFiles := map[string]string{
		filepath.Join(tempDir, ".gitignore"):           "*.tmp\n",
		filepath.Join(tempDir, ".gitattributes"):       "* text=auto\n",
		filepath.Join(tempDir, "submodule", ".git"):    "gitdir: ../.git/modules/submodule\n",
		filepath.Join(tempDir, "regular", "sheet.txt"): "content\n",
	}

	for file, content := range testFiles {
		if err := os.WriteFile(file, []byte(content), 0644); err != nil {
			t.Fatalf("failed to create file %s: %v", file, err)
		}
	}

	tests := []struct {
		name    string
		path    string
		want    bool
		wantErr bool
	}{
		{
			name: "not in git directory",
			path: filepath.Join(tempDir, "regular", "sheet.txt"),
			want: false,
		},
		{
			name: "in .git directory",
			path: filepath.Join(tempDir, ".git", "objects", "file"),
			want: true,
		},
		{
			name: "in .git/refs directory",
			path: filepath.Join(tempDir, ".git", "refs", "heads", "main"),
			want: true,
		},
		{
			name: ".gitignore file",
			path: filepath.Join(tempDir, ".gitignore"),
			want: false,
		},
		{
			name: ".gitattributes file",
			path: filepath.Join(tempDir, ".gitattributes"),
			want: false,
		},
		{
			name: "submodule with .git file",
			path: filepath.Join(tempDir, "submodule", "sheet.txt"),
			want: false,
		},
		{
			name: "path with .git in middle",
			path: filepath.Join(tempDir, "regular", ".git", "sheet.txt"),
			want: true,
		},
		{
			name: "nonexistent path without .git",
			path: filepath.Join(tempDir, "nonexistent", "file"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GitDir(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GitDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GitDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGitDirEdgeCases(t *testing.T) {
	// Test with paths that have .git but not as a directory separator
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "file ending with .git",
			path: "/tmp/myfile.git",
			want: false,
		},
		{
			name: "directory ending with .git",
			path: "/tmp/myrepo.git",
			want: false,
		},
		{
			name: ".github directory",
			path: "/tmp/.github/workflows",
			want: false,
		},
		{
			name: "legitimate.git-repo name",
			path: "/tmp/legitimate.git-repo/file",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GitDir(tt.path)
			if err != nil {
				// It's ok if the path doesn't exist for these edge case tests
				return
			}
			if got != tt.want {
				t.Errorf("GitDir(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestGitDirPathSeparator(t *testing.T) {
	// Test that the function correctly uses os.PathSeparator
	// This is important for cross-platform compatibility

	// Create a path with the wrong separator for the current OS
	var wrongSep string
	if os.PathSeparator == '/' {
		wrongSep = `\`
	} else {
		wrongSep = `/`
	}

	// Path with wrong separator should not be detected as git dir
	path := fmt.Sprintf("some%spath%s.git%sfile", wrongSep, wrongSep, wrongSep)
	isGit, err := GitDir(path)

	if err != nil {
		// Path doesn't exist, which is fine
		return
	}

	if isGit {
		t.Errorf("GitDir() incorrectly detected git dir with wrong path separator")
	}
}
