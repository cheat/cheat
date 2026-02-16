package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/repo"
)

// cmdUpdate updates git-backed cheatpaths.
func cmdUpdate(_ *cobra.Command, _ []string, conf config.Config) {

	hasError := false

	for _, path := range conf.Cheatpaths {
		err := repo.Pull(path.Path)

		switch {
		case err == nil:
			fmt.Printf("%s: ok\n", path.Name)

		case errors.Is(err, git.ErrRepositoryNotExists):
			fmt.Printf("%s: skipped (not a git repository)\n", path.Name)

		case errors.Is(err, repo.ErrDirtyWorktree):
			fmt.Printf("%s: skipped (dirty worktree)\n", path.Name)

		default:
			fmt.Fprintf(os.Stderr, "%s: error (%v)\n", path.Name, err)
			hasError = true
		}
	}

	if hasError {
		os.Exit(1)
	}
}
