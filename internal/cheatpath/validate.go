package cheatpath

import (
	"fmt"
)

// Validate ensures that the Path is valid
func (c Path) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("cheatpath name cannot be empty")
	}
	if c.Path == "" {
		return fmt.Errorf("cheatpath path cannot be empty")
	}
	return nil
}
