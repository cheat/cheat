package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Init initializes a config file
func Init(confpath string, configs string) error {

	// assert that the config directory exists
	if err := os.MkdirAll(filepath.Dir(confpath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// write the config file
	if err := ioutil.WriteFile(confpath, []byte(configs), 0644); err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}

	return nil
}
