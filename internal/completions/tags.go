package completions

import (
	"github.com/spf13/cobra"

	"github.com/cheat/cheat/internal/sheets"
)

// Tags provides completion for the --tag flag.
func Tags(
	_ *cobra.Command,
	_ []string,
	_ string,
) ([]string, cobra.ShellCompDirective) {

	conf, err := loadConfig()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	cheatsheets, err := sheets.Load(conf.Cheatpaths)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	return sheets.Tags(cheatsheets), cobra.ShellCompDirectiveNoFileComp
}
