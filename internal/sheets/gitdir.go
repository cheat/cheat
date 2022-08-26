package sheets

import (
	"fmt"
	"os"
	"strings"
)

// isGitDir returns `true` if `path` is within a `.git` directory, or `false`
// otherwise
func isGitDir(path string) (bool, error) {

	/*
		A bit of context is called for here, because this functionality has
		previously caused a number of tricky, subtle bugs.

		Fundamentally, here we are simply trying to avoid walking over the
		contents of the `.git` directory. Doing so potentially makes
		hundreds/thousands of needless syscalls, and can noticeably harm
		performance on machines with slow disks.

		The earliest effort to solve this problem involved simply returning
		`fs.SkipDir` when the cheatsheet file path began with `.`, signifying a
		hidden directory. This, however, caused two problems:

		1. The `.cheat` directory was ignored
		2. Cheatsheets installed by `brew` (which were by default installed to
		`~/.config/cheat`) were ignored

		See: https://github.com/cheat/cheat/issues/690

		To remedy this, the exclusion criteria were narrowed, and the search
		for a literal `.` was replaced with a search for a literal `.git`.
		This, however, broke user installations that stored cheatsheets in
		`git` submodules, because such an installation would contain a `.git`
		file that pointed to the upstream repository.

		See: https://github.com/cheat/cheat/issues/694

		Accounting for all of the above, we are now searching for the presence
		of a `.git` literal string in the cheatsheet file path. If it is not
		found, we continue to walk the directory, as before.

		If it *is* found, we determine if `.git` refers to a file or directory,
		and only stop walking the path in the latter case.
	*/

	// determine if the literal string `.git` appears within `path`
	pos := strings.Index(path, ".git")

	// if it does not, we know for certain that we are not within a `.git`
	// directory.
	if pos == -1 {
		return false, nil
	}

	// If `path` does contain the string `.git`, we need to determine if we're
	// inside of a `.git` directory, or if `path` points to a cheatsheet that's
	// stored within a `git` submodule.
	//
	// See: https://github.com/cheat/cheat/issues/694

	// truncate `path` to the occurrence of `.git`
	f, err := os.Stat(path[:pos+4])
	if err != nil {
		return false, fmt.Errorf("failed to stat path %s: %v", path, err)
	}

	// return true or false depending on whether the truncated path is a
	// directory
	return f.Mode().IsDir(), nil
}
