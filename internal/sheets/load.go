package sheets

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	cp "github.com/cheat/cheat/internal/cheatpath"
	"github.com/cheat/cheat/internal/sheet"
)

// Load produces a map of cheatsheet titles to filesystem paths
func Load(cheatpaths []cp.Cheatpath) ([]map[string]sheet.Sheet, error) {

	// create a slice of maps of sheets. This structure will store all sheets
	// that are associated with each cheatpath.
	sheets := make([]map[string]sheet.Sheet, len(cheatpaths))

	// iterate over each cheatpath
	for _, cheatpath := range cheatpaths {

		// vivify the map of cheatsheets on this specific cheatpath
		pathsheets := make(map[string]sheet.Sheet)

		// recursively iterate over the cheatpath, and load each cheatsheet
		// encountered along the way
		err := filepath.Walk(
			cheatpath.Path, func(
				path string,
				info os.FileInfo,
				err error) error {

				// fail if an error occurred while walking the directory
				if err != nil {
					return fmt.Errorf("failed to walk path: %v", err)
				}

				// don't register directories as cheatsheets
				if info.IsDir() {
					return nil
				}

				// calculate the cheatsheet's "title" (the phrase with which it may be
				// accessed. Eg: `cheat tar` - `tar` is the title)
				title := strings.TrimPrefix(
					strings.TrimPrefix(path, cheatpath.Path),
					string(os.PathSeparator),
				)

				// Don't walk the `.git` directory. Doing so creates
				// hundreds/thousands of needless syscalls and could
				// potentially harm performance on machines with slow disks.
				skip, err := isGitDir(path)
				if err != nil {
					return fmt.Errorf("failed to identify .git directory: %v", err)
				}
				if skip {
					return fs.SkipDir
				}

				// parse the cheatsheet file into a `sheet` struct
				s, err := sheet.New(
					title,
					cheatpath.Name,
					path,
					cheatpath.Tags,
					cheatpath.ReadOnly,
				)
				if err != nil {
					return fmt.Errorf(
						"failed to load sheet: %s, path: %s, err: %v",
						title,
						path,
						err,
					)
				}

				// register the cheatsheet on its cheatpath, keyed by its title
				pathsheets[title] = s
				return nil
			})
		if err != nil {
			return sheets, fmt.Errorf("failed to load cheatsheets: %v", err)
		}

		// store the sheets on this cheatpath alongside the other cheatsheets on
		// other cheatpaths
		sheets = append(sheets, pathsheets)
	}

	// return the cheatsheets, grouped by cheatpath
	return sheets, nil
}
