package sheet

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Copy copies a cheatsheet to a new location
func (s *Sheet) Copy(dest string) error {

	// NB: while the `infile` has already been loaded and parsed into a `sheet`
	// struct, we're going to read it again here. This is a bit wasteful, but
	// necessary if we want the "raw" file contents (including the front-matter).
	// This is because the frontmatter is parsed and then discarded when the file
	// is loaded via `sheets.Load`.
	infile, err := os.Open(s.Path)
	if err != nil {
		return fmt.Errorf("failed to open cheatsheet: %s, %v", s.Path, err)
	}
	defer infile.Close()

	// create any necessary subdirectories
	dirs := filepath.Dir(dest)
	if dirs != "." {
		if err := os.MkdirAll(dirs, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %s, %v", dirs, err)
		}
	}

	// create the outfile
	outfile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create outfile: %s, %v", dest, err)
	}
	defer outfile.Close()

	// copy file contents
	_, err = io.Copy(outfile, infile)
	if err != nil {
		// Clean up the partially written file on error
		os.Remove(dest)
		return fmt.Errorf(
			"failed to copy file: infile: %s, outfile: %s, err: %v",
			s.Path,
			dest,
			err,
		)
	}

	return nil
}
