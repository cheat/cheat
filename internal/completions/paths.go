package completions

import (
	"github.com/spf13/cobra"
)

// Paths provides completion for the --path flag.
func Paths(
	_ *cobra.Command,
	_ []string,
	_ string,
) ([]string, cobra.ShellCompDirective) {

	conf, err := loadConfig()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	names := make([]string, 0, len(conf.Cheatpaths))
	for _, cp := range conf.Cheatpaths {
		names = append(names, cp.Name)
	}

	return names, cobra.ShellCompDirectiveNoFileComp
}
