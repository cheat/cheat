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

		The next attempt at solving this was to search for a `.git` literal
		string in the cheatsheet file path. If a match was not found, we would
		continue to walk the directory, as before.

		If a match *was* found, we determined whether `.git` referred to a file
		or directory, and would only stop walking the path in the latter case.

		This, however, caused crashes if a cheatpath contained a `.gitignore`
		file. (Presumably, a crash would likewise occur on the presence of
		`.gitattributes`, `.gitmodules`, etc.)

		See: https://github.com/cheat/cheat/issues/699

		Accounting for all of the above (hopefully?), the current solution is
		not to search for `.git`, but `.git/` (including the directory
		separator), and then only ceasing to walk the directory on a match.

		To summarize, this code must account for the following possibilities:

		1. A cheatpath is not a repository
		2. A cheatpath is a repository
		3. A cheatpath is a repository, and contains a `.git*` file
		4. A cheatpath is a submodule

		Care must be taken to support the above on both Unix and Windows
		systems, which have different directory separators and line-endings.

		There is a lot of nuance to all of this, and it would be worthwhile to
		do two things to stop writing bugs here:

		1. Build integration tests around all of this
		2. Discard string-matching solutions entirely, and use `go-git` instead

		NB: A reasonable smoke-test for ensuring that skipping is being applied
		correctly is to run the following command:

		    make && strace ./dist/cheat -l | wc -l

		That check should be run twice: once normally, and once after
		commenting out the "skip" check in `sheets.Load`.

		The specific line counts don't matter; what matters is that the number
		of syscalls should be significantly lower with the skip check enabled.
	*/

	// determine if the literal string `.git` appears within `path`
	pos := strings.Index(path, fmt.Sprintf(".git%s", string(os.PathSeparator)))

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
