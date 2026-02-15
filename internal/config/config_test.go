package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/go-homedir"

	"github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/mock"
)

// TestConfig asserts that the configs are loaded correctly
func TestConfigSuccessful(t *testing.T) {

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
	conf, err := New(map[string]interface{}{}, mock.Path("conf/conf.yml"), false)
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
	want := []cheatpath.Cheatpath{
		cheatpath.Cheatpath{
			Path:     filepath.Join(home, ".dotfiles", "cheat", "community"),
			ReadOnly: true,
			Tags:     []string{"community"},
		},
		cheatpath.Cheatpath{
			Path:     filepath.Join(home, ".dotfiles", "cheat", "work"),
			ReadOnly: false,
			Tags:     []string{"work"},
		},
		cheatpath.Cheatpath{
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
	_, err := New(map[string]interface{}{}, "/does-not-exit", false)
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
	conf, err := New(map[string]interface{}{}, mock.Path("conf/conf.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "vim" {
		t.Errorf("expected config file editor: want: vim, got: %s", conf.Editor)
	}

	// $EDITOR should override the config file value
	os.Setenv("EDITOR", "nano")
	conf, err = New(map[string]interface{}{}, mock.Path("conf/conf.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "nano" {
		t.Errorf("$EDITOR should override config: want: nano, got: %s", conf.Editor)
	}

	// $VISUAL should override both $EDITOR and the config file value
	os.Setenv("VISUAL", "emacs")
	conf, err = New(map[string]interface{}{}, mock.Path("conf/conf.yml"), false)
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
	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "foo" {
		t.Errorf("failed to respect $EDITOR: want: foo, got: %s", conf.Editor)
	}

	// set $VISUAL and assert it takes precedence over $EDITOR
	os.Setenv("VISUAL", "bar")
	conf, err = New(map[string]interface{}{}, mock.Path("conf/empty.yml"), false)
	if err != nil {
		t.Fatalf("failed to init configs: %v", err)
	}
	if conf.Editor != "bar" {
		t.Errorf("failed to respect $VISUAL: want: bar, got: %s", conf.Editor)
	}
}
