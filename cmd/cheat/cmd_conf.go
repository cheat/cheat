package main

import (
	"fmt"

	"github.com/cheat/cheat/internal/config"
)

func cmdConf(opts map[string]interface{}, conf config.Config) {
	fmt.Println(conf.Path)
}
