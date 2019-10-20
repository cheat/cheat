package sheet

// Tagged returns true if a sheet was tagged with `needle`
func (s *Sheet) Tagged(needle string) bool {

	// if any of the tags match `needle`, return `true`
	for _, tag := range s.Tags {
		if tag == needle {
			return true
		}
	}

	// otherwise, return `false`
	return false
}
