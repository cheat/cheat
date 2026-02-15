package completions

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// Generate writes a shell completion script to the given writer.
func Generate(cmd *cobra.Command, shell string, w io.Writer) error {
	switch shell {
	case "bash":
		return cmd.Root().GenBashCompletionV2(w, true)
	case "zsh":
		return cmd.Root().GenZshCompletion(w)
	case "fish":
		return cmd.Root().GenFishCompletion(w, true)
	case "powershell":
		return cmd.Root().GenPowerShellCompletionWithDesc(w)
	default:
		return fmt.Errorf("unsupported shell: %s (valid: bash, zsh, fish, powershell)", shell)
	}
}
