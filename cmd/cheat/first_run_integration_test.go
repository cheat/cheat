package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestFirstRunIntegration exercises the end-to-end first-run experience:
// no config exists, the binary creates one, and subsequent runs succeed.
// This is the regression test for issues #721, #771, and #730.
func TestFirstRunIntegration(t *testing.T) {
	// Build the cheat binary
	binName := "cheat_test"
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	binPath := filepath.Join(t.TempDir(), binName)
	build := exec.Command("go", "build", "-o", binPath, ".")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build cheat: %v\nOutput: %s", err, output)
	}

	t.Run("init comments out community", func(t *testing.T) {
		testHome := t.TempDir()
		env := firstRunEnv(testHome)

		cmd := exec.Command(binPath, "--init")
		cmd.Env = env
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("--init failed: %v\nOutput: %s", err, output)
		}
		outStr := string(output)

		// No placeholder strings should survive (regression for #721)
		assertNoPlaceholders(t, outStr)

		// Community cheatpath should be commented out
		assertCommunityCommentedOut(t, outStr)

		// Personal and work cheatpaths should be active (uncommented)
		assertCheatpathActive(t, outStr, "personal")
		assertCheatpathActive(t, outStr, "work")

		// Should include clone instructions
		if !strings.Contains(outStr, "git clone") {
			t.Error("expected git clone instructions in --init output")
		}

		// Save the config and verify it loads without errors.
		// --init only outputs config, it doesn't create directories,
		// so we need to create the cheatpath dirs the config references.
		confpath := filepath.Join(testHome, "conf.yml")
		if err := os.WriteFile(confpath, output, 0644); err != nil {
			t.Fatalf("failed to write config: %v", err)
		}

		// Determine the confdir that --init used (same logic as cmd_init.go)
		initConfpaths := firstRunConfpaths(testHome)
		initConfdir := filepath.Dir(initConfpaths[0])
		for _, name := range []string{"personal", "work"} {
			dir := filepath.Join(initConfdir, "cheatsheets", name)
			if err := os.MkdirAll(dir, 0755); err != nil {
				t.Fatalf("failed to create %s dir: %v", name, err)
			}
		}

		cmd2 := exec.Command(binPath, "--directories")
		cmd2.Env = append(append([]string{}, env...), "CHEAT_CONFIG_PATH="+confpath)
		output2, err := cmd2.CombinedOutput()
		if err != nil {
			t.Fatalf("config from --init failed to load: %v\nOutput: %s", err, output2)
		}
	})

	t.Run("decline config creation", func(t *testing.T) {
		testHome := t.TempDir()
		env := firstRunEnv(testHome)

		cmd := exec.Command(binPath)
		cmd.Env = env
		cmd.Stdin = strings.NewReader("n\n")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("cheat exited with error: %v\nOutput: %s", err, output)
		}

		// Verify no config was created
		if firstRunConfigExists(testHome) {
			t.Error("config file was created despite user declining")
		}
	})

	t.Run("accept config decline community", func(t *testing.T) {
		testHome := t.TempDir()
		env := firstRunEnv(testHome)

		// First run: yes to create config, no to community cheatsheets
		cmd := exec.Command(binPath)
		cmd.Env = env
		cmd.Stdin = strings.NewReader("y\nn\n")
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("first run failed: %v\nOutput: %s", err, output)
		}
		outStr := string(output)

		// Parse the config path from output
		confpath := parseCreatedConfPath(t, outStr)
		if confpath == "" {
			t.Fatalf("could not find config path in output:\n%s", outStr)
		}

		// Verify config file exists
		if _, err := os.Stat(confpath); os.IsNotExist(err) {
			t.Fatalf("config file not found at %s", confpath)
		}

		// Verify config file contents
		content, err := os.ReadFile(confpath)
		if err != nil {
			t.Fatalf("failed to read config: %v", err)
		}
		contentStr := string(content)

		// No placeholder strings should survive (regression for #721)
		assertNoPlaceholders(t, contentStr)

		// Community cheatpath should be commented out
		assertCommunityCommentedOut(t, contentStr)

		// Personal and work cheatpaths should be active (uncommented)
		assertCheatpathActive(t, contentStr, "personal")
		assertCheatpathActive(t, contentStr, "work")

		// Verify personal and work directories were created
		confdir := filepath.Dir(confpath)
		for _, name := range []string{"personal", "work"} {
			dir := filepath.Join(confdir, "cheatsheets", name)
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				t.Errorf("expected %s directory at %s", name, dir)
			}
		}

		// Community directory should NOT exist
		communityDir := filepath.Join(confdir, "cheatsheets", "community")
		if _, err := os.Stat(communityDir); err == nil {
			t.Error("community directory should not exist when declined")
		}

		// --- Second run: verify the config loads successfully ---
		// This is the core regression test for #721/#771/#730:
		// previously, the second run would fail because config.New()
		// hard-errored on the missing community cheatpath directory.
		// Use --directories (not --list, which exits 2 when no sheets exist).
		cmd2 := exec.Command(binPath, "--directories")
		cmd2.Env = append(append([]string{}, env...), "CHEAT_CONFIG_PATH="+confpath)
		output2, err := cmd2.CombinedOutput()
		if err != nil {
			t.Fatalf(
				"second run failed (regression for #721/#771/#730): %v\nOutput: %s",
				err, output2,
			)
		}

		// Verify the output lists the expected cheatpaths
		outStr2 := string(output2)
		if !strings.Contains(outStr2, "personal") {
			t.Errorf("expected 'personal' cheatpath in --directories output:\n%s", outStr2)
		}
		if !strings.Contains(outStr2, "work") {
			t.Errorf("expected 'work' cheatpath in --directories output:\n%s", outStr2)
		}
	})
}

// firstRunEnv returns a minimal environment for a clean first-run test.
func firstRunEnv(home string) []string {
	env := []string{
		"PATH=" + os.Getenv("PATH"),
	}

	switch runtime.GOOS {
	case "windows":
		env = append(env,
			"APPDATA="+filepath.Join(home, "AppData", "Roaming"),
			"USERPROFILE="+home,
			"SystemRoot="+os.Getenv("SystemRoot"),
		)
	default:
		env = append(env,
			"HOME="+home,
			"EDITOR=vi",
		)
	}

	return env
}

// parseCreatedConfPath extracts the config file path from the installer's
// "Created config file: <path>" output. The message may appear mid-line
// (after prompt text), so we search for the substring anywhere in the output.
func parseCreatedConfPath(t *testing.T, output string) string {
	t.Helper()
	const marker = "Created config file: "
	idx := strings.Index(output, marker)
	if idx < 0 {
		return ""
	}
	rest := output[idx+len(marker):]
	// the path ends at the next newline
	if nl := strings.IndexByte(rest, '\n'); nl >= 0 {
		rest = rest[:nl]
	}
	return strings.TrimSpace(rest)
}

// firstRunConfpaths returns the config file paths that cheat would check
// for the given home directory, matching the logic in config.Paths().
func firstRunConfpaths(home string) []string {
	switch runtime.GOOS {
	case "windows":
		return []string{
			filepath.Join(home, "AppData", "Roaming", "cheat", "conf.yml"),
		}
	default:
		return []string{
			filepath.Join(home, ".config", "cheat", "conf.yml"),
		}
	}
}

// assertNoPlaceholders verifies that no template placeholder strings survived
// in the config output. This is the regression check for #721 (literal
// PAGER_PATH appearing in the config).
func assertNoPlaceholders(t *testing.T, content string) {
	t.Helper()
	placeholders := []string{
		"PAGER_PATH",
		"COMMUNITY_PATH",
		"PERSONAL_PATH",
		"WORK_PATH",
	}
	for _, p := range placeholders {
		if strings.Contains(content, p) {
			t.Errorf("placeholder %q was not replaced in config", p)
		}
	}
	// EDITOR_PATH is special: it survives if no editor is found.
	// In our test env EDITOR=vi is set, so it should be replaced.
	if strings.Contains(content, "editor: EDITOR_PATH") {
		t.Error("placeholder EDITOR_PATH was not replaced in config")
	}
}

// assertCommunityCommentedOut verifies that the community cheatpath entry
// is commented out (not active) in the config.
func assertCommunityCommentedOut(t *testing.T, content string) {
	t.Helper()
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "- name: community" {
			t.Error("community cheatpath should be commented out")
			return
		}
	}
	if !strings.Contains(content, "#- name: community") {
		t.Error("expected commented-out community cheatpath")
	}
}

// assertCheatpathActive verifies that a named cheatpath is present and
// uncommented in the config.
func assertCheatpathActive(t *testing.T, content string, name string) {
	t.Helper()
	marker := "- name: " + name
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == marker {
			return
		}
	}
	t.Errorf("expected active (uncommented) cheatpath %q", name)
}

// firstRunConfigExists checks whether a cheat config file exists under the
// given home directory at any of the standard locations.
func firstRunConfigExists(home string) bool {
	candidates := []string{
		filepath.Join(home, ".config", "cheat", "conf.yml"),
		filepath.Join(home, ".cheat", "conf.yml"),
		filepath.Join(home, "AppData", "Roaming", "cheat", "conf.yml"),
	}
	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return true
		}
	}
	return false
}
