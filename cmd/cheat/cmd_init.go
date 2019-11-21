package main

import (
	"fmt"
	"os"
	"runtime"
	"text/template"

	"github.com/cheat/cheat/internal/config"
)

// cmdInit displays an example config file.
func cmdInit() {

	prefFolderPath, err := config.PreferredFolderPath(runtime.GOOS)
	if err != nil {
		fmt.Fprintln(os.Stderr, "could not locate config folder path: ", err)
		os.Exit(1)
	}

	if prefFolderPath != "" {
		prefFolderPath += string(os.PathSeparator)
	}

	type ConfigValues struct {
		CheatsheetsBasePath string
	}

	values := ConfigValues{prefFolderPath}
	t := template.Must(template.New("configs").Parse(configs()))
	err = t.Execute(os.Stdout, values)
	if err != nil {
		fmt.Fprintln(os.Stderr, "executing template:: ", err)
		os.Exit(1)
	}
}
