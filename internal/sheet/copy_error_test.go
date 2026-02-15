package sheet

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// TestCopyErrors tests error cases for the Copy method
func TestCopyErrors(t *testing.T) {
	tests := []struct {
		name  string
		setup func() (*Sheet, string, func())
	}{
		{
			name: "source file does not exist",
			setup: func() (*Sheet, string, func()) {
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
		},
		{
			name: "destination directory creation fails",
			setup: func() (*Sheet, string, func()) {
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

				blockerFile := filepath.Join(os.TempDir(), "copy-blocker-file")
				if err := os.WriteFile(blockerFile, []byte("blocker"), 0644); err != nil {
					t.Fatalf("failed to create blocker file: %v", err)
				}

				dest := filepath.Join(blockerFile, "subdir", "dest.txt")

				cleanup := func() {
					os.Remove(src.Name())
					os.Remove(blockerFile)
				}
				return sheet, dest, cleanup
			},
		},
		{
			name: "destination file creation fails",
			setup: func() (*Sheet, string, func()) {
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sheet, dest, cleanup := tt.setup()
			defer cleanup()

			err := sheet.Copy(dest)
			if err == nil {
				t.Error("Copy() expected error, got nil")
			}
		})
	}
}

// TestCopyUnreadableSource verifies that Copy returns an error when the source
// file cannot be opened (e.g., permission denied).
func TestCopyUnreadableSource(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("chmod does not restrict reads on Windows")
	}

	src, err := os.CreateTemp("", "copy-test-unreadable-*")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(src.Name())

	if _, err := src.WriteString("test content"); err != nil {
		t.Fatalf("failed to write content: %v", err)
	}
	src.Close()

	sheet := &Sheet{
		Title:     "test",
		Path:      src.Name(),
		CheatPath: "test",
	}

	dest := filepath.Join(os.TempDir(), "copy-unreadable-test.txt")
	defer os.Remove(dest)

	if err := os.Chmod(src.Name(), 0000); err != nil {
		t.Skip("Cannot change file permissions on this platform")
	}
	defer os.Chmod(src.Name(), 0644)

	err = sheet.Copy(dest)
	if err == nil {
		t.Error("expected Copy to fail with permission error")
	}

	// Destination should not exist since the error occurs before it is created
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		t.Error("destination file should not exist after open failure")
	}
}
