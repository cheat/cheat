package repo

import (
	"fmt"
	"os"
	"strings"
)

// gitSep is the `.git` path component surrounded by path separators.
// Used to match `.git` as a complete path component, not as a suffix
// of a directory name (e.g., `personal.git`).
var gitSep = string(os.PathSeparator) + ".git" + string(os.PathSeparator)

// GitDir returns `true` if we are iterating over a directory contained within
// a repositories `.git` directory.
func GitDir(path string) (bool, error) {

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

		The next attempt at solving this was to search for a `.git` literal
		string in the cheatsheet file path. If a match was not found, we would
		continue to walk the directory, as before.

		If a match *was* found, we determined whether `.git` referred to a file
		or directory, and would only stop walking the path in the latter case.

		This, however, caused crashes if a cheatpath contained a `.gitignore`
		file. (Presumably, a crash would likewise occur on the presence of
		`.gitattributes`, `.gitmodules`, etc.)

		See: https://github.com/cheat/cheat/issues/699

		Accounting for all of the above, the next solution was to search not
		for `.git`, but `.git/` (including the directory separator), and then
		only ceasing to walk the directory on a match.

		This, however, also had a bug: searching for `.git/` also matched
		directory names that *ended with* `.git`, like `personal.git/`. This
		caused cheatsheets stored under such paths to be silently skipped.

		See: https://github.com/cheat/cheat/issues/711

		The current (and hopefully final) solution requires the path separator
		on *both* sides of `.git`, i.e., searching for `/.git/`. This ensures
		that `.git` is matched only as a complete path component, not as a
		suffix of a directory name.

		To summarize, this code must account for the following possibilities:

		1. A cheatpath is not a repository
		2. A cheatpath is a repository
		3. A cheatpath is a repository, and contains a `.git*` file
		4. A cheatpath is a submodule
		5. A cheatpath is a hidden directory
		6. A cheatpath is inside a directory whose name ends with `.git`

		Care must be taken to support the above on both Unix and Windows
		systems, which have different directory separators and line-endings.

		NB: `filepath.Walk` always passes absolute paths to the walk function,
		so `.git` will never appear as the first path component. This is what
		makes the "separator on both sides" approach safe.

		A reasonable smoke-test for ensuring that skipping is being applied
		correctly is to run the following command:

		    make && strace ./dist/cheat -l | wc -l

		That check should be run twice: once normally, and once after
		commenting out the "skip" check in `sheets.Load`.

		The specific line counts don't matter; what matters is that the number
		of syscalls should be significantly lower with the skip check enabled.
	*/

	// determine if `.git` appears as a complete path component
	pos := strings.Index(path, gitSep)

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
	f, err := os.Stat(path[:pos+5])
	if err != nil {
		return false, fmt.Errorf("failed to stat path %s: %v", path, err)
	}

	// return true or false depending on whether the truncated path is a
	// directory
	return f.Mode().IsDir(), nil
}
