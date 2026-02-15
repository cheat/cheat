// Package cheatpath implements functions pertaining to cheatsheet file path
// management.
package cheatpath

// Path encapsulates cheatsheet path information
type Path struct {
	Name     string   `yaml:"name"`
	Path     string   `yaml:"path"`
	ReadOnly bool     `yaml:"readonly"`
	Tags     []string `yaml:"tags"`
}
