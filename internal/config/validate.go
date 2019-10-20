package config

import (
	"fmt"
)

// Validate returns an error if the config is invalid
func (c *Config) Validate() error {

	// assert that an editor was specified
	if c.Editor == "" {
		return fmt.Errorf("config error: editor unspecified")
	}

	// assert that at least one cheatpath was specified
	if len(c.Cheatpaths) == 0 {
		return fmt.Errorf("config error: no cheatpaths specified")
	}

	// assert that each path and name is unique
	names := make(map[string]bool)
	paths := make(map[string]bool)

	// assert that each cheatpath is valid
	for _, cheatpath := range c.Cheatpaths {

		// assert that the cheatpath is valid
		if err := cheatpath.Validate(); err != nil {
			return fmt.Errorf("config error: %v", err)
		}

		// assert that the name is unique
		if _, ok := names[cheatpath.Name]; ok {
			return fmt.Errorf(
				"config error: cheatpath name is not unique: %s",
				cheatpath.Name,
			)
		}
		names[cheatpath.Name] = true

		// assert that the path is unique
		if _, ok := paths[cheatpath.Path]; ok {
			return fmt.Errorf(
				"config error: cheatpath path is not unique: %s",
				cheatpath.Path,
			)
		}
		paths[cheatpath.Path] = true
	}

	// TODO: assert valid styles?

	// assert that the formatter is valid
	formatters := map[string]bool{
		"terminal":    true,
		"terminal256": true,
		"terminal16m": true,
	}
	if _, ok := formatters[c.Formatter]; !ok {
		return fmt.Errorf("config error: formatter is invalid: %s", c.Formatter)
	}

	return nil
}
