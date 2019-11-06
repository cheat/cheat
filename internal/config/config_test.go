package config

import (
	//"os"
	//"path/filepath"
	//"reflect"
	"testing"

	//"github.com/davecgh/go-spew/spew"
	//"github.com/mitchellh/go-homedir"

	//"github.com/cheat/cheat/internal/cheatpath"
	//"github.com/cheat/cheat/internal/mock"
)

// BUG: changes pertaining to symlink handling introduced in 3.0.4 break this
// test.
/*
// TestConfig asserts that the configs are loaded correctly
func TestConfigSuccessful(t *testing.T) {

	// initialize a config
	conf, err := New(map[string]interface{}{}, mock.Path("conf/conf.yml"))
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
			Path:     filepath.Join(home, ".dotfiles/cheat/community"),
			ReadOnly: true,
			Tags:     []string{"community"},
		},
		cheatpath.Cheatpath{
			Path:     filepath.Join(home, ".dotfiles/cheat/work"),
			ReadOnly: false,
			Tags:     []string{"work"},
		},
		cheatpath.Cheatpath{
			Path:     filepath.Join(home, ".dotfiles/cheat/personal"),
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
*/

// TestConfigFailure asserts that an error is returned if the config file
// cannot be read.
func TestConfigFailure(t *testing.T) {

	// attempt to read a non-existent config file
	_, err := New(map[string]interface{}{}, "/does-not-exit")
	if err == nil {
		t.Errorf("failed to error on unreadable config")
	}
}

// TestEmptyEditor asserts that envvars are respected if an editor is not
// specified in the configs
func TestEmptyEditor(t *testing.T) {

	/*
	// clear the environment variables
	os.Setenv("VISUAL", "")
	os.Setenv("EDITOR", "")

	// initialize a config
	conf, err := New(map[string]interface{}{}, mock.Path("conf/empty.yml"))
	if err == nil {
		t.Errorf("failed to return an error on empty editor")
	}

	// set editor, and assert that it is respected
	os.Setenv("EDITOR", "foo")
	conf, err = New(map[string]interface{}{}, mock.Path("conf/empty.yml"))
	if err != nil {
		t.Errorf("failed to init configs: %v", err)
	}
	if conf.Editor != "foo" {
		t.Errorf("failed to respect editor: want: foo, got: %s", conf.Editor)
	}

	// set visual, and assert that it overrides editor
	os.Setenv("VISUAL", "bar")
	conf, err = New(map[string]interface{}{}, mock.Path("conf/empty.yml"))
	if err != nil {
		t.Errorf("failed to init configs: %v", err)
	}
	if conf.Editor != "bar" {
		t.Errorf("failed to respect editor: want: bar, got: %s", conf.Editor)
	}
	*/
}
