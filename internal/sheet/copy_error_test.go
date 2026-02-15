package sheet

import (
	"os"
	"path/filepath"
	"testing"
)

// TestCopyErrors tests error cases for the Copy method
func TestCopyErrors(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() (*Sheet, string, func())
		wantErr bool
		errMsg  string
	}{
		{
			name: "source file does not exist",
			setup: func() (*Sheet, string, func()) {
				// Create a sheet with non-existent path
				sheet := &Sheet{
					Title:     "test",
					Path:      "/non/existent/file.txt",
					CheatPath: "test",
				}
				dest := filepath.Join(os.TempDir(), "copy-test-dest.txt")
				cleanup := func() {
					os.Remove(dest)
				}
				return sheet, dest, cleanup
			},
			wantErr: true,
			errMsg:  "failed to open cheatsheet",
		},
		{
			name: "destination directory creation fails",
			setup: func() (*Sheet, string, func()) {
				// Create a source file
				src, err := os.CreateTemp("", "copy-test-src-*")
				if err != nil {
					t.Fatalf("failed to create temp file: %v", err)
				}
				src.WriteString("test content")
				src.Close()

				sheet := &Sheet{
					Title:     "test",
					Path:      src.Name(),
					CheatPath: "test",
				}

				// Create a file where we want a directory
				blockerFile := filepath.Join(os.TempDir(), "copy-blocker-file")
				if err := os.WriteFile(blockerFile, []byte("blocker"), 0644); err != nil {
					t.Fatalf("failed to create blocker file: %v", err)
				}

				// Try to create dest under the blocker file (will fail)
				dest := filepath.Join(blockerFile, "subdir", "dest.txt")

				cleanup := func() {
					os.Remove(src.Name())
					os.Remove(blockerFile)
				}
				return sheet, dest, cleanup
			},
			wantErr: true,
			errMsg:  "failed to create directory",
		},
		{
			name: "destination file creation fails",
			setup: func() (*Sheet, string, func()) {
				// Create a source file
				src, err := os.CreateTemp("", "copy-test-src-*")
				if err != nil {
					t.Fatalf("failed to create temp file: %v", err)
				}
				src.WriteString("test content")
				src.Close()

				sheet := &Sheet{
					Title:     "test",
					Path:      src.Name(),
					CheatPath: "test",
				}

				// Create a directory where we want the file
				destDir := filepath.Join(os.TempDir(), "copy-test-dir")
				if err := os.Mkdir(destDir, 0755); err != nil && !os.IsExist(err) {
					t.Fatalf("failed to create dest dir: %v", err)
				}

				cleanup := func() {
					os.Remove(src.Name())
					os.RemoveAll(destDir)
				}
				return sheet, destDir, cleanup
			},
			wantErr: true,
			errMsg:  "failed to create outfile",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sheet, dest, cleanup := tt.setup()
			defer cleanup()

			err := sheet.Copy(dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.errMsg != "" {
				if !contains(err.Error(), tt.errMsg) {
					t.Errorf("Copy() error = %v, want error containing %q", err, tt.errMsg)
				}
			}
		})
	}
}

// TestCopyIOError tests the io.Copy error case
func TestCopyIOError(t *testing.T) {
	// This is difficult to test without mocking io.Copy
	// The error case would occur if the source file is modified
	// or removed after opening but before copying
	t.Skip("Skipping io.Copy error test - requires file system race condition")
}

// TestCopyCleanupOnError verifies that partially written files are cleaned up on error
func TestCopyCleanupOnError(t *testing.T) {
	// Create a source file that we'll make unreadable after opening
	src, err := os.CreateTemp("", "copy-test-cleanup-*")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(src.Name())

	// Write some content
	content := "test content for cleanup"
	if _, err := src.WriteString(content); err != nil {
		t.Fatalf("failed to write content: %v", err)
	}
	src.Close()

	sheet := &Sheet{
		Title:     "test",
		Path:      src.Name(),
		CheatPath: "test",
	}

	// Destination path
	dest := filepath.Join(os.TempDir(), "copy-cleanup-test.txt")
	defer os.Remove(dest) // Clean up if test fails

	// Make the source file unreadable (simulating a read error during copy)
	// This is platform-specific, but should work on Unix-like systems
	if err := os.Chmod(src.Name(), 0000); err != nil {
		t.Skip("Cannot change file permissions on this platform")
	}
	defer os.Chmod(src.Name(), 0644) // Restore permissions for cleanup

	// Attempt to copy - this should fail during io.Copy
	err = sheet.Copy(dest)
	if err == nil {
		t.Error("Expected Copy to fail with permission error")
	}

	// Verify the destination file was cleaned up
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Error("Destination file should have been removed after copy failure")
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsHelper(s, substr))
}

func containsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
