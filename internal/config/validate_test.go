package config

import (
	"testing"

	"github.com/cheat/cheat/internal/cheatpath"
)

// TestValidateCorrect asserts that valid configs are validated successfully
func TestValidateCorrect(t *testing.T) {

	// mock a config
	conf := Config{
		Colorize:  true,
		Editor:    "vim",
		Formatter: "terminal16m",
		Cheatpaths: []cheatpath.Cheatpath{
			cheatpath.Cheatpath{
				Name:     "foo",
				Path:     "/foo",
				ReadOnly: false,
				Tags:     []string{},
			},
		},
	}

	// assert that no errors are returned
	if err := conf.Validate(); err != nil {
		t.Errorf("failed to validate valid config: %v", err)
	}
}

// TestInvalidateMissingEditor asserts that configs with unspecified editors
// are invalidated
func TestInvalidateMissingEditor(t *testing.T) {

	// mock a config
	conf := Config{
		Colorize:  true,
		Formatter: "terminal16m",
		Cheatpaths: []cheatpath.Cheatpath{
			cheatpath.Cheatpath{
				Name:     "foo",
				Path:     "/foo",
				ReadOnly: false,
				Tags:     []string{},
			},
		},
	}

	// assert that no errors are returned
	if err := conf.Validate(); err == nil {
		t.Errorf("failed to invalidate config with unspecified editor")
	}
}

// TestInvalidateMissingCheatpaths asserts that configs without cheatpaths are
// invalidated
func TestInvalidateMissingCheatpaths(t *testing.T) {

	// mock a config
	conf := Config{
		Colorize:  true,
		Editor:    "vim",
		Formatter: "terminal16m",
	}

	// assert that no errors are returned
	if err := conf.Validate(); err == nil {
		t.Errorf("failed to invalidate config without cheatpaths")
	}
}

// TestMissingInvalidFormatters asserts that configs which contain invalid
// formatters are invalidated
func TestMissingInvalidFormatters(t *testing.T) {

	// mock a config
	conf := Config{
		Colorize: true,
		Editor:   "vim",
	}

	// assert that no errors are returned
	if err := conf.Validate(); err == nil {
		t.Errorf("failed to invalidate config without formatter")
	}
}

// TestInvalidateDuplicateCheatpathNames asserts that configs which contain
// cheatpaths with duplcated names are invalidated
func TestInvalidateDuplicateCheatpathNames(t *testing.T) {

	// mock a config
	conf := Config{
		Colorize:  true,
		Editor:    "vim",
		Formatter: "terminal16m",
		Cheatpaths: []cheatpath.Cheatpath{
			cheatpath.Cheatpath{
				Name:     "foo",
				Path:     "/foo",
				ReadOnly: false,
				Tags:     []string{},
			},
			cheatpath.Cheatpath{
				Name:     "foo",
				Path:     "/bar",
				ReadOnly: false,
				Tags:     []string{},
			},
		},
	}

	// assert that no errors are returned
	if err := conf.Validate(); err == nil {
		t.Errorf("failed to invalidate config with cheatpaths with duplicate names")
	}
}

// TestInvalidateDuplicateCheatpathPaths asserts that configs which contain
// cheatpaths with duplcated paths are invalidated
func TestInvalidateDuplicateCheatpathPaths(t *testing.T) {

	// mock a config
	conf := Config{
		Colorize:  true,
		Editor:    "vim",
		Formatter: "terminal16m",
		Cheatpaths: []cheatpath.Cheatpath{
			cheatpath.Cheatpath{
				Name:     "foo",
				Path:     "/foo",
				ReadOnly: false,
				Tags:     []string{},
			},
			cheatpath.Cheatpath{
				Name:     "bar",
				Path:     "/foo",
				ReadOnly: false,
				Tags:     []string{},
			},
		},
	}

	// assert that no errors are returned
	if err := conf.Validate(); err == nil {
		t.Errorf("failed to invalidate config with cheatpaths with duplicate paths")
	}
}
