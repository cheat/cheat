package config

import (
	"testing"
)

// TestColor asserts that colorization rules are properly respected
func TestColor(t *testing.T) {

	// mock a config
	conf := Config{}

	opts := map[string]interface{}{"--colorize": false}
	if conf.Color(opts) {
		t.Errorf("failed to respect --colorize (false)")
	}

	opts = map[string]interface{}{"--colorize": true}
	if !conf.Color(opts) {
		t.Errorf("failed to respect --colorize (true)")
	}
}
