package main

import (
	"fmt"

	"github.com/cheat/cheat/internal/config"
)

func cmdConf(_ map[string]interface{}, conf config.Config) {
	fmt.Println(conf.Path)
}
