// Package completions provides dynamic shell completion functions and
// completion script generation for the cheat CLI.
package completions

import (
	"sort"

	"github.com/spf13/cobra"

	"github.com/cheat/cheat/internal/sheets"
)

// Cheatsheets provides completion for cheatsheet names.
func Cheatsheets(
	_ *cobra.Command,
	args []string,
	_ string,
) ([]string, cobra.ShellCompDirective) {

	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	conf, err := loadConfig()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	consolidated := sheets.Consolidate(cheatsheets)

	names := make([]string, 0, len(consolidated))
	for name := range consolidated {
		names = append(names, name)
	}
	sort.Strings(names)

	return names, cobra.ShellCompDirectiveNoFileComp
}
