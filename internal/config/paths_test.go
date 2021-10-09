package config

import (
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

// TestValidatePathsNix asserts that the proper config paths are returned on
// *nix platforms
func TestValidatePathsNix(t *testing.T) {

	// mock the user's home directory
	home := "/home/foo"

	// mock some envvars
	envvars := map[string]string{
		"XDG_CONFIG_HOME": "/home/bar",
	}

	// specify the platforms to test
	oses := []string{
		"android",
		"darwin",
		"freebsd",
		"linux",
	}

	// test each *nix os
	for _, os := range oses {
		// get the paths for the platform
		paths, err := Paths(os, home, envvars)
		if err != nil {
			t.Errorf("paths returned an error: %v", err)
		}

		// specify the expected output
		want := []string{
			"/home/bar/cheat/conf.yml",
			"/home/foo/.config/cheat/conf.yml",
			"/home/foo/.cheat/conf.yml",
			"/etc/cheat/conf.yml",
		}

		// assert that output matches expectations
		if !reflect.DeepEqual(paths, want) {
			t.Errorf(
				"failed to return expected paths: want:\n%s, got:\n%s",
				spew.Sdump(want),
				spew.Sdump(paths),
			)
		}
	}
}

// TestValidatePathsNixNoXDG asserts that the proper config paths are returned
// on *nix platforms when `XDG_CONFIG_HOME is not set
func TestValidatePathsNixNoXDG(t *testing.T) {

	// mock the user's home directory
	home := "/home/foo"

	// mock some envvars
	envvars := map[string]string{}

	// specify the platforms to test
	oses := []string{
		"darwin",
		"freebsd",
		"linux",
	}

	// test each *nix os
	for _, os := range oses {
		// get the paths for the platform
		paths, err := Paths(os, home, envvars)
		if err != nil {
			t.Errorf("paths returned an error: %v", err)
		}

		// specify the expected output
		want := []string{
			"/home/foo/.config/cheat/conf.yml",
			"/home/foo/.cheat/conf.yml",
			"/etc/cheat/conf.yml",
		}

		// assert that output matches expectations
		if !reflect.DeepEqual(paths, want) {
			t.Errorf(
				"failed to return expected paths: want:\n%s, got:\n%s",
				spew.Sdump(want),
				spew.Sdump(paths),
			)
		}
	}
}

// TestValidatePathsWindows asserts that the proper config paths are returned
// on Windows platforms
func TestValidatePathsWindows(t *testing.T) {

	// mock the user's home directory
	home := "not-used-on-windows"

	// mock some envvars
	envvars := map[string]string{
		"APPDATA":     "/apps",
		"PROGRAMDATA": "/programs",
	}

	// get the paths for the platform
	paths, err := Paths("windows", home, envvars)
	if err != nil {
		t.Errorf("paths returned an error: %v", err)
	}

	// specify the expected output
	want := []string{
		"/apps/cheat/conf.yml",
		"/programs/cheat/conf.yml",
	}

	// assert that output matches expectations
	if !reflect.DeepEqual(paths, want) {
		t.Errorf(
			"failed to return expected paths: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(paths),
		)
	}
}

// TestValidatePathsUnsupported asserts that an error is returned on
// unsupported platforms
func TestValidatePathsUnsupported(t *testing.T) {
	_, err := Paths("unsupported", "", map[string]string{})
	if err == nil {
		t.Errorf("failed to return error on unsupported platform")
	}
}

// TestValidatePathsCheatConfigPath asserts that the proper config path is
// returned when `CHEAT_CONFIG_PATH` is explicitly specified.
func TestValidatePathsCheatConfigPath(t *testing.T) {

	// mock the user's home directory
	home := "/home/foo"

	// mock some envvars
	envvars := map[string]string{
		"XDG_CONFIG_HOME":   "/home/bar",
		"CHEAT_CONFIG_PATH": "/home/baz/conf.yml",
	}

	// get the paths for the platform
	paths, err := Paths("linux", home, envvars)
	if err != nil {
		t.Errorf("paths returned an error: %v", err)
	}

	// specify the expected output
	want := []string{
		"/home/baz/conf.yml",
	}

	// assert that output matches expectations
	if !reflect.DeepEqual(paths, want) {
		t.Errorf(
			"failed to return expected paths: want:\n%s, got:\n%s",
			spew.Sdump(want),
			spew.Sdump(paths),
		)
	}
}
