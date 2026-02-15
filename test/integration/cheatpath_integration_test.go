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

// hasCwdCheatpath checks whether the --directories output contains a
// cheatpath named "cwd".  The output format is "name: path\n" per line
// (tabwriter-aligned), so we look for a line beginning with "cwd".
func hasCwdCheatpath(output string) bool {
	for _, line := range strings.Split(output, "\n") {
		if strings.HasPrefix(line, "cwd") {
			return true
		}
	}
	return false
}

// TestLocalCheatpathIntegration exercises the recursive .cheat directory
// discovery end-to-end: it builds the real cheat binary, sets up filesystem
// layouts, and verifies behaviour from the user's perspective.
func TestLocalCheatpathIntegration(t *testing.T) {
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

	// cheatEnv returns a minimal environment for the cheat binary.
	cheatEnv := func(confPath, home string) []string {
		return []string{
			"CHEAT_CONFIG_PATH=" + confPath,
			"HOME=" + home,
			"PATH=" + os.Getenv("PATH"),
			"EDITOR=vi",
		}
	}

	// writeConfig writes a minimal valid config file referencing sheetsDir.
	writeConfig := func(t *testing.T, dir, sheetsDir string) string {
		t.Helper()
		conf := fmt.Sprintf("---\neditor: vi\ncolorize: false\ncheatpaths:\n  - name: base\n    path: %s\n    readonly: true\n", sheetsDir)
		confPath := filepath.Join(dir, "conf.yml")
		if err := os.WriteFile(confPath, []byte(conf), 0644); err != nil {
			t.Fatalf("failed to write config: %v", err)
		}
		return confPath
	}

	t.Run("parent .cheat is discovered from subdirectory", func(t *testing.T) {
		root := t.TempDir()

		// Configured cheatpath (empty but must exist for validation)
		sheetsDir := filepath.Join(root, "sheets")
		os.MkdirAll(sheetsDir, 0755)

		// .cheat at root with a cheatsheet
		dotCheat := filepath.Join(root, ".cheat")
		os.Mkdir(dotCheat, 0755)
		os.WriteFile(
			filepath.Join(dotCheat, "localsheet"),
			[]byte("---\nsyntax: bash\n---\necho hello from local\n"),
			0644,
		)

		confPath := writeConfig(t, root, sheetsDir)

		// Work from a subdirectory
		workDir := filepath.Join(root, "src", "pkg")
		os.MkdirAll(workDir, 0755)
		env := cheatEnv(confPath, root)

		// --directories should list "cwd" cheatpath
		cmd := exec.Command(binPath, "--directories")
		cmd.Dir = workDir
		cmd.Env = env
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat --directories failed: %v\nOutput: %s", err, output)
		}
		if !hasCwdCheatpath(string(output)) {
			t.Errorf("expected 'cwd' cheatpath in --directories output:\n%s", output)
		}

		// Viewing the cheatsheet should show its content
		cmd2 := exec.Command(binPath, "localsheet")
		cmd2.Dir = workDir
		cmd2.Env = env
		output2, err := cmd2.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat localsheet failed: %v\nOutput: %s", err, output2)
		}
		if !strings.Contains(string(output2), "echo hello from local") {
			t.Errorf("expected cheatsheet content, got:\n%s", output2)
		}
	})

	t.Run("grandparent .cheat is discovered from deep subdirectory", func(t *testing.T) {
		root := t.TempDir()

		sheetsDir := filepath.Join(root, "sheets")
		os.MkdirAll(sheetsDir, 0755)

		dotCheat := filepath.Join(root, ".cheat")
		os.Mkdir(dotCheat, 0755)
		os.WriteFile(
			filepath.Join(dotCheat, "deepsheet"),
			[]byte("---\nsyntax: bash\n---\ndeep discovery works\n"),
			0644,
		)

		confPath := writeConfig(t, root, sheetsDir)

		deepDir := filepath.Join(root, "a", "b", "c", "d", "e")
		os.MkdirAll(deepDir, 0755)

		cmd := exec.Command(binPath, "deepsheet")
		cmd.Dir = deepDir
		cmd.Env = cheatEnv(confPath, root)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat deepsheet failed: %v\nOutput: %s", err, output)
		}
		if !strings.Contains(string(output), "deep discovery works") {
			t.Errorf("expected cheatsheet content, got:\n%s", output)
		}
	})

	t.Run("nearest .cheat wins over ancestor .cheat", func(t *testing.T) {
		root := t.TempDir()

		sheetsDir := filepath.Join(root, "sheets")
		os.MkdirAll(sheetsDir, 0755)

		// .cheat at root
		rootCheat := filepath.Join(root, ".cheat")
		os.Mkdir(rootCheat, 0755)
		os.WriteFile(
			filepath.Join(rootCheat, "shared"),
			[]byte("---\nsyntax: bash\n---\nfrom root\n"),
			0644,
		)

		// .cheat at project/ (nearer)
		projectDir := filepath.Join(root, "project")
		os.MkdirAll(projectDir, 0755)
		projectCheat := filepath.Join(projectDir, ".cheat")
		os.Mkdir(projectCheat, 0755)
		os.WriteFile(
			filepath.Join(projectCheat, "shared"),
			[]byte("---\nsyntax: bash\n---\nfrom project nearest\n"),
			0644,
		)

		confPath := writeConfig(t, root, sheetsDir)

		workDir := filepath.Join(projectDir, "src")
		os.MkdirAll(workDir, 0755)
		env := cheatEnv(confPath, root)

		// --directories should list the nearer cheatpath
		cmd := exec.Command(binPath, "--directories")
		cmd.Dir = workDir
		cmd.Env = env
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat --directories failed: %v\nOutput: %s", err, output)
		}
		if !strings.Contains(string(output), projectCheat) {
			t.Errorf("expected project .cheat path in output, got:\n%s", output)
		}

		// "shared" sheet should come from the nearer .cheat
		cmd2 := exec.Command(binPath, "shared")
		cmd2.Dir = workDir
		cmd2.Env = env
		output2, err := cmd2.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat shared failed: %v\nOutput: %s", err, output2)
		}
		if !strings.Contains(string(output2), "from project nearest") {
			t.Errorf("expected nearest .cheat content, got:\n%s", output2)
		}
	})

	t.Run("no .cheat directory means no cwd cheatpath", func(t *testing.T) {
		root := t.TempDir()

		sheetsDir := filepath.Join(root, "sheets")
		os.MkdirAll(sheetsDir, 0755)
		// Need at least one sheet for --directories to work without error
		os.WriteFile(filepath.Join(sheetsDir, "placeholder"),
			[]byte("---\nsyntax: bash\n---\nplaceholder\n"), 0644)

		confPath := writeConfig(t, root, sheetsDir)

		// No .cheat anywhere under root
		cmd := exec.Command(binPath, "--directories")
		cmd.Dir = root
		cmd.Env = cheatEnv(confPath, root)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat --directories failed: %v\nOutput: %s", err, output)
		}
		if hasCwdCheatpath(string(output)) {
			t.Errorf("'cwd' cheatpath should not appear when no .cheat exists:\n%s", output)
		}
	})

	t.Run(".cheat file (not directory) is ignored", func(t *testing.T) {
		root := t.TempDir()

		sheetsDir := filepath.Join(root, "sheets")
		os.MkdirAll(sheetsDir, 0755)
		os.WriteFile(filepath.Join(sheetsDir, "placeholder"),
			[]byte("---\nsyntax: bash\n---\nplaceholder\n"), 0644)

		// Create .cheat as a regular file
		os.WriteFile(filepath.Join(root, ".cheat"), []byte("not a dir"), 0644)

		confPath := writeConfig(t, root, sheetsDir)

		cmd := exec.Command(binPath, "--directories")
		cmd.Dir = root
		cmd.Env = cheatEnv(confPath, root)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat --directories failed: %v\nOutput: %s", err, output)
		}
		if hasCwdCheatpath(string(output)) {
			t.Errorf("'cwd' should not appear for a .cheat file:\n%s", output)
		}
	})
}
