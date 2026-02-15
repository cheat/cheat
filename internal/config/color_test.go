package config

import (
	"testing"
)

// TestColor asserts that colorization rules are properly respected
func TestColor(t *testing.T) {

	// mock a config
	conf := Config{}

	if conf.Color(false) {
		t.Errorf("failed to respect forceColorize (false)")
	}

	if !conf.Color(true) {
		t.Errorf("failed to respect forceColorize (true)")
	}
}
