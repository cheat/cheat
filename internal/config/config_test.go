package config

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/go-homedir"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/mocks"
)

// TestFindLocalCheatpathInCurrentDir tests that .cheat in the given dir is found
func TestFindLocalCheatpathInCurrentDir(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cheatDir := filepath.Join(tempDir, ".cheat")
	if err := os.Mkdir(cheatDir, 0755); err != nil {
		t.Fatalf("failed to create .cheat dir: %v", err)
	}

	result := findLocalCheatpath(tempDir)
	if result != cheatDir {
		t.Errorf("expected %s, got %s", cheatDir, result)
	}
}

// TestFindLocalCheatpathInParent tests walking up to a parent directory
func TestFindLocalCheatpathInParent(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cheatDir := filepath.Join(tempDir, ".cheat")
	if err := os.Mkdir(cheatDir, 0755); err != nil {
		t.Fatalf("failed to create .cheat dir: %v", err)
	}

	subDir := filepath.Join(tempDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create sub dir: %v", err)
	}

	result := findLocalCheatpath(subDir)
	if result != cheatDir {
		t.Errorf("expected %s, got %s", cheatDir, result)
	}
}

// TestFindLocalCheatpathInGrandparent tests walking up multiple levels
func TestFindLocalCheatpathInGrandparent(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	cheatDir := filepath.Join(tempDir, ".cheat")
	if err := os.Mkdir(cheatDir, 0755); err != nil {
		t.Fatalf("failed to create .cheat dir: %v", err)
	}

	deepDir := filepath.Join(tempDir, "a", "b", "c")
	if err := os.MkdirAll(deepDir, 0755); err != nil {
		t.Fatalf("failed to create deep dir: %v", err)
	}

	result := findLocalCheatpath(deepDir)
	if result != cheatDir {
		t.Errorf("expected %s, got %s", cheatDir, result)
	}
}

// TestFindLocalCheatpathNearestWins tests that the closest .cheat is returned
func TestFindLocalCheatpathNearestWins(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create .cheat at root level
	if err := os.Mkdir(filepath.Join(tempDir, ".cheat"), 0755); err != nil {
		t.Fatalf("failed to create root .cheat dir: %v", err)
	}

	// Create sub/.cheat (the nearer one)
	subDir := filepath.Join(tempDir, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create sub dir: %v", err)
	}
	nearCheatDir := filepath.Join(subDir, ".cheat")
	if err := os.Mkdir(nearCheatDir, 0755); err != nil {
		t.Fatalf("failed to create sub .cheat dir: %v", err)
	}

	// Search from sub/deep/
	deepDir := filepath.Join(subDir, "deep")
	if err := os.Mkdir(deepDir, 0755); err != nil {
		t.Fatalf("failed to create deep dir: %v", err)
	}

	result := findLocalCheatpath(deepDir)
	if result != nearCheatDir {
		t.Errorf("expected nearest %s, got %s", nearCheatDir, result)
	}
}

// TestFindLocalCheatpathNotFound tests that empty string is returned when no .cheat exists
func TestFindLocalCheatpathNotFound(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	result := findLocalCheatpath(tempDir)
	if result != "" {
		t.Errorf("expected empty string, got %s", result)
	}
}

// TestFindLocalCheatpathSkipsFile tests that a file named .cheat is not matched
func TestFindLocalCheatpathSkipsFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "cheat-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create .cheat as a file, not a directory
	cheatFile := filepath.Join(tempDir, ".cheat")
	if err := os.WriteFile(cheatFile, []byte("not a directory"), 0644); err != nil {
		t.Fatalf("failed to create .cheat file: %v", err)
	}

	result := findLocalCheatpath(tempDir)
	if result != "" {
		t.Errorf("expected empty string for .cheat file, got %s", result)
	}
}

// TestFindLocalCheatpathSymlink tests that a .cheat symlink to a directory is found
func TestFindLocalCheatpathSymlink(t *testing.T) {
	tempDir := t.TempDir()

	// Create the real directory
	realDir := filepath.Join(tempDir, "real-cheat")
	if err := os.Mkdir(realDir, 0755); err != nil {
		t.Fatalf("failed to create real dir: %v", err)
	}

	// Symlink .cheat -> real-cheat
	cheatLink := filepath.Join(tempDir, ".cheat")
	if err := os.Symlink(realDir, cheatLink); err != nil {
		t.Fatalf("failed to create symlink: %v", err)
	}

	result := findLocalCheatpath(tempDir)
	if result != cheatLink {
		t.Errorf("expected %s, got %s", cheatLink, result)
	}
}

// TestFindLocalCheatpathSymlinkInAncestor tests discovery through a symlinked
// ancestor directory. When the cwd is reached via a symlink, filepath.Dir
// walks the symlinked path (not the real path), so .cheat must be findable
// through that chain.
func TestFindLocalCheatpathSymlinkInAncestor(t *testing.T) {
	tempDir := t.TempDir()

	// Create real/project/.cheat
	realProject := filepath.Join(tempDir, "real", "project")
	if err := os.MkdirAll(realProject, 0755); err != nil {
		t.Fatalf("failed to create real project dir: %v", err)
	}
	if err := os.Mkdir(filepath.Join(realProject, ".cheat"), 0755); err != nil {
		t.Fatalf("failed to create .cheat dir: %v", err)
	}

	// Create symlink: linked -> real/project
	linkedProject := filepath.Join(tempDir, "linked")
	if err := os.Symlink(realProject, linkedProject); err != nil {
		t.Fatalf("failed to create symlink: %v", err)
	}

	// Create sub inside the symlinked path
	subDir := filepath.Join(linkedProject, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create sub dir: %v", err)
	}

	// Search from linked/sub — should find linked/.cheat
	// (os.Stat follows symlinks, so linked/.cheat resolves to real/project/.cheat)
	result := findLocalCheatpath(subDir)
	expected := filepath.Join(linkedProject, ".cheat")
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

// TestFindLocalCheatpathPermissionDenied tests that unreadable ancestor
// directories are skipped and the walk continues upward.
func TestFindLocalCheatpathPermissionDenied(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Unix permissions do not apply on Windows")
	}
	if os.Getuid() == 0 {
		t.Skip("test requires non-root user")
	}

	tempDir := t.TempDir()

	// Resolve symlinks (macOS /var -> /private/var)
	tempDir, err := filepath.EvalSymlinks(tempDir)
	if err != nil {
		t.Fatalf("failed to resolve symlinks: %v", err)
	}

	// Create tempDir/.cheat (the target we want found)
	cheatDir := filepath.Join(tempDir, ".cheat")
	if err := os.Mkdir(cheatDir, 0755); err != nil {
		t.Fatalf("failed to create .cheat dir: %v", err)
	}

	// Create tempDir/restricted/ with its own .cheat and sub/
	restricted := filepath.Join(tempDir, "restricted")
	if err := os.Mkdir(restricted, 0755); err != nil {
		t.Fatalf("failed to create restricted dir: %v", err)
	}
	if err := os.Mkdir(filepath.Join(restricted, ".cheat"), 0755); err != nil {
		t.Fatalf("failed to create restricted .cheat dir: %v", err)
	}
	subDir := filepath.Join(restricted, "sub")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create sub dir: %v", err)
	}

	// Make restricted/ unreadable — blocks stat of children
	if err := os.Chmod(restricted, 0000); err != nil {
		t.Fatalf("failed to chmod: %v", err)
	}
	t.Cleanup(func() { os.Chmod(restricted, 0755) })

	// Walk from restricted/sub: stat("restricted/sub/.cheat") fails (EACCES),
	// stat("restricted/.cheat") fails (EACCES), walk continues to tempDir/.cheat
	result := findLocalCheatpath(subDir)
	if result != cheatDir {
		t.Errorf("expected %s (walked past restricted dir), got %s", cheatDir, result)
	}
}

// TestConfig asserts that the configs are loaded correctly
func TestConfigSuccessful(t *testing.T) {

	// Chdir into a temp directory so no ancestor .cheat directory can
	// leak into the cheatpaths (findLocalCheatpath walks the full
	// ancestor chain).
	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get cwd: %v", err)
	}
	defer os.Chdir(oldCwd)
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatalf("failed to chdir: %v", err)
	}

	// clear env vars so they don't override the config file value
	oldVisual := os.Getenv("VISUAL")
	oldEditor := os.Getenv("EDITOR")
	os.Unsetenv("VISUAL")
	os.Unsetenv("EDITOR")
	defer func() {
		os.Setenv("VISUAL", oldVisual)
		os.Setenv("EDITOR", oldEditor)
	}()

	// initialize a config
	conf, err := New(mocks.Path("conf/conf.yml"), false)
	if err != nil {
		t.Errorf("failed to parse config file: %v", err)
	}

	// assert that the expected values were returned
	if conf.Editor != "vim" {
		t.Errorf("failed to set editor: want: vim, got: %s", conf.Editor)
	}
	if !conf.Colorize {
		t.Errorf("failed to set colorize: want: true, got: %t", conf.Colorize)
	}

	// get the user's home directory (with ~ expanded)
	home, err := homedir.Dir()
	if err != nil {
		t.Errorf("failed to get homedir: %v", err)
	}

	// assert that the cheatpaths are correct
	want := []cheatpath.Path{
		cheatpath.Path{
			Path:     filepath.Join(home, ".dotfiles", "cheat", "community"),
			ReadOnly: true,
			Tags:     []string{"community"},
		},
		cheatpath.Path{
			Path:     filepath.Join(home, ".dotfiles", "cheat", "work"),
			ReadOnly: false,
			Tags:     []string{"work"},
		},
		cheatpath.Path{
			Path:     filepath.Join(home, ".dotfiles", "cheat", "personal"),
			ReadOnly: false,
			Tags:     []string{"personal"},
		},
	}

	if !reflect.DeepEqual(conf.Cheatpaths, want) {
		t.Errorf(
			"failed to return expected results: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(conf.Cheatpaths),
		)
	}
}

// TestConfigFailure asserts that an error is returned if the config file
// cannot be read.
func TestConfigFailure(t *testing.T) {

	// attempt to read a non-existent config file
	_, err := New("/does-not-exit", false)
	if err == nil {
		t.Errorf("failed to error on unreadable config")
	}
}

// TestEditorEnvOverride asserts that $VISUAL and $EDITOR override the
// config file value at runtime (regression test for #589)
func TestEditorEnvOverride(t *testing.T) {
	// save and clear the environment variables
	oldVisual := os.Getenv("VISUAL")
	oldEditor := os.Getenv("EDITOR")
	defer func() {
		os.Setenv("VISUAL", oldVisual)
		os.Setenv("EDITOR", oldEditor)
	}()

	// with no env vars, the config file value should be used
	os.Unsetenv("VISUAL")
	os.Unsetenv("EDITOR")
	conf, err := New(mocks.Path("conf/conf.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "vim" {
		t.Errorf("expected config file editor: want: vim, got: %s", conf.Editor)
	}

	// $EDITOR should override the config file value
	os.Setenv("EDITOR", "nano")
	conf, err = New(mocks.Path("conf/conf.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "nano" {
		t.Errorf("$EDITOR should override config: want: nano, got: %s", conf.Editor)
	}

	// $VISUAL should override both $EDITOR and the config file value
	os.Setenv("VISUAL", "emacs")
	conf, err = New(mocks.Path("conf/conf.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "emacs" {
		t.Errorf("$VISUAL should override all: want: emacs, got: %s", conf.Editor)
	}
}

// TestEditorEnvFallback asserts that env vars are used as fallback when
// no editor is specified in the config file
func TestEditorEnvFallback(t *testing.T) {
	// save and clear the environment variables
	oldVisual := os.Getenv("VISUAL")
	oldEditor := os.Getenv("EDITOR")
	defer func() {
		os.Setenv("VISUAL", oldVisual)
		os.Setenv("EDITOR", oldEditor)
	}()

	// set $EDITOR and assert it's used when config has no editor
	os.Unsetenv("VISUAL")
	os.Setenv("EDITOR", "foo")
	conf, err := New(mocks.Path("conf/empty.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "foo" {
		t.Errorf("failed to respect $EDITOR: want: foo, got: %s", conf.Editor)
	}

	// set $VISUAL and assert it takes precedence over $EDITOR
	os.Setenv("VISUAL", "bar")
	conf, err = New(mocks.Path("conf/empty.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "bar" {
		t.Errorf("failed to respect $VISUAL: want: bar, got: %s", conf.Editor)
	}
}
