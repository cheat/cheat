// Package cheatpath implements functions pertaining to cheatsheet file path
// management.
package cheatpath

import "fmt"

// Cheatpath encapsulates cheatsheet path information
type Cheatpath struct {
	Name     string   `yaml:"name"`
	Path     string   `yaml:"path"`
	ReadOnly bool     `yaml:"readonly"`
	Tags     []string `yaml:"tags"`
}

// Validate ensures that the Cheatpath is valid
func (c Cheatpath) Validate() error {
	// Check that name is not empty
	if c.Name == "" {
		return fmt.Errorf("cheatpath name cannot be empty")
	}

	// Check that path is not empty
	if c.Path == "" {
		return fmt.Errorf("cheatpath path cannot be empty")
	}

	return nil
}
