package config

import (
	"os"
	"testing"
)

// TestInit asserts that configs are properly initialized
func TestInit(t *testing.T) {

	// initialize a temporary config file
	confFile, err := os.CreateTemp("", "cheat-test")
	if err != nil {
		t.Errorf("failed to create temp file: %v", err)
	}

	// clean up the temp file
	defer os.Remove(confFile.Name())

	// initialize the config file
	conf := "mock config data"
	if err = Init(confFile.Name(), conf); err != nil {
		t.Errorf("failed to init config file: %v", err)
	}

	// read back the config file contents
	bytes, err := os.ReadFile(confFile.Name())
	if err != nil {
		t.Errorf("failed to read config file: %v", err)
	}

	// assert that the contents were written correctly
	got := string(bytes)
	if got != conf {
		t.Errorf("failed to write configs: want: %s, got: %s", conf, got)
	}
}
