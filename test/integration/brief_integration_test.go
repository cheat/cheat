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

// TestBriefFlagIntegration exercises the -b/--brief flag end-to-end.
func TestBriefFlagIntegration(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("integration test uses Unix-specific env vars")
	}

	// Build the cheat binary once for all sub-tests.
	binPath := filepath.Join(t.TempDir(), "cheat_test")
	build := exec.Command("go", "build", "-o", binPath, "./cmd/cheat")
	build.Dir = repoRoot(t)
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build cheat: %v\nOutput: %s", err, output)
	}

	// Set up a temp environment with some cheatsheets.
	root := t.TempDir()
	sheetsDir := filepath.Join(root, "sheets")
	os.MkdirAll(sheetsDir, 0755)

	os.WriteFile(
		filepath.Join(sheetsDir, "tar"),
		[]byte("---\nsyntax: bash\ntags: [ compression ]\n---\ntar xf archive.tar\n"),
		0644,
	)
	os.WriteFile(
		filepath.Join(sheetsDir, "curl"),
		[]byte("---\nsyntax: bash\ntags: [ networking, http ]\n---\ncurl https://example.com\n"),
		0644,
	)

	confPath := filepath.Join(root, "conf.yml")
	conf := fmt.Sprintf("---\neditor: vi\ncolorize: false\ncheatpaths:\n  - name: test\n    path: %s\n    readonly: true\n", sheetsDir)
	os.WriteFile(confPath, []byte(conf), 0644)

	env := []string{
		"CHEAT_CONFIG_PATH=" + confPath,
		"HOME=" + root,
		"PATH=" + os.Getenv("PATH"),
		"EDITOR=vi",
	}

	run := func(t *testing.T, args ...string) string {
		t.Helper()
		cmd := exec.Command(binPath, args...)
		cmd.Dir = root
		cmd.Env = env
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat %v failed: %v\nOutput: %s", args, err, output)
		}
		return string(output)
	}

	t.Run("brief output omits file path column", func(t *testing.T) {
		output := run(t, "-b")
		lines := strings.Split(strings.TrimSpace(output), "\n")

		// Header should have title and tags but not file
		if !strings.Contains(lines[0], "title:") {
			t.Errorf("expected title: in header, got: %s", lines[0])
		}
		if !strings.Contains(lines[0], "tags:") {
			t.Errorf("expected tags: in header, got: %s", lines[0])
		}
		if strings.Contains(lines[0], "file:") {
			t.Errorf("brief output should not contain file: column, got: %s", lines[0])
		}

		// Data lines should not contain the sheets directory path
		for _, line := range lines[1:] {
			if strings.Contains(line, sheetsDir) {
				t.Errorf("brief output should not contain file paths, got: %s", line)
			}
		}
	})

	t.Run("list output still includes file path column", func(t *testing.T) {
		output := run(t, "-l")
		lines := strings.Split(strings.TrimSpace(output), "\n")

		if !strings.Contains(lines[0], "file:") {
			t.Errorf("list output should contain file: column, got: %s", lines[0])
		}
	})

	t.Run("brief with filter works", func(t *testing.T) {
		output := run(t, "-b", "tar")
		if !strings.Contains(output, "tar") {
			t.Errorf("expected tar in output, got: %s", output)
		}
		if strings.Contains(output, "curl") {
			t.Errorf("filter should exclude curl, got: %s", output)
		}
	})

	t.Run("combined -lb works identically to -b", func(t *testing.T) {
		briefOnly := run(t, "-b", "tar")
		combined := run(t, "-lb", "tar")
		if briefOnly != combined {
			t.Errorf("-b and -lb should produce identical output\n-b:\n%s\n-lb:\n%s", briefOnly, combined)
		}
	})

	t.Run("brief with tag filter works", func(t *testing.T) {
		output := run(t, "-b", "-t", "networking")
		if !strings.Contains(output, "curl") {
			t.Errorf("expected curl in tag-filtered output, got: %s", output)
		}
		if strings.Contains(output, "tar") {
			// tar is tagged "compression", not "networking"
			t.Errorf("tag filter should exclude tar, got: %s", output)
		}
		if strings.Contains(output, "file:") {
			t.Errorf("brief output should not contain file: column, got: %s", output)
		}
	})
}
