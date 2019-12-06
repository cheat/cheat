package cheatpath

// Cheatpath encapsulates cheatsheet path information
type Cheatpath struct {
	Name     string   `yaml:name`
	Path     string   `yaml:path`
	ReadOnly bool     `yaml:readonly`
	Tags     []string `yaml:tags`
}
