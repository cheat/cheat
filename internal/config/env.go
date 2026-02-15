package config

import (
	"os"
	"runtime"
	"strings"
)

// EnvVars reads environment variables into a map of strings.
func EnvVars() map[string]string {
	envvars := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if runtime.GOOS == "windows" {
			pair[0] = strings.ToUpper(pair[0])
		}
		envvars[pair[0]] = pair[1]
	}
	return envvars
}
