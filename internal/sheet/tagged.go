package sheet

import "slices"

// Tagged returns true if a sheet was tagged with `needle`
func (s *Sheet) Tagged(needle string) bool {
	return slices.Contains(s.Tags, needle)
}
