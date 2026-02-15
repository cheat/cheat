package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cheat/cheat/internal/config"
)

func cmdConf(_ *cobra.Command, _ []string, conf config.Config) {
	fmt.Println(conf.Path)
}
