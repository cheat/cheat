package display

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cheat/cheat/internal/config"
)

// Write writes output either directly to stdout, or through a pager,
// depending upon configuration.
func Write(out string, conf config.Config) {
	// if no pager was configured, print the output to stdout and exit
	if conf.Pager == "" {
		fmt.Print(out)
		os.Exit(0)
	}

	// otherwise, pipe output through the pager
	parts := strings.Split(conf.Pager, " ")
	pager := parts[0]
	args := parts[1:]

	// run the pager
	cmd := exec.Command(pager, args...)
	cmd.Stdin = strings.NewReader(out)
	cmd.Stdout = os.Stdout

	// handle errors
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write to pager: %v\n", err)
		os.Exit(1)
	}
}
