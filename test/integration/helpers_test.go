package integration

import (
	"path/filepath"
	"runtime"
	"testing"
)

// repoRoot returns the absolute path to the repository root.
// It derives this from the known location of this source file
// (test/integration/) relative to the repo root.
func repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to determine repo root via runtime.Caller")
	}
	// file is <repo>/test/integration/helpers_test.go → go up two dirs
	return filepath.Dir(filepath.Dir(filepath.Dir(file)))
}

// repoRootBench is the same as repoRoot but for use in benchmarks.
func repoRootBench(b *testing.B) string {
	b.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		b.Fatal("failed to determine repo root via runtime.Caller")
	}
	return filepath.Dir(filepath.Dir(filepath.Dir(file)))
}
