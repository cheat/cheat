package config

import (
	"fmt"
	"io/ioutil"
	"os"

	cp "github.com/cheat/cheat/internal/cheatpath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

// Config encapsulates configuration parameters
type Config struct {
	Colorize   bool           `yaml:colorize`
	Editor     string         `yaml:editor`
	Cheatpaths []cp.Cheatpath `yaml:cheatpaths`
	Style      string         `yaml:style`
	Formatter  string         `yaml:formatter`
}

// New returns a new Config struct
func New(opts map[string]interface{}, confPath string) (Config, error) {

	// read the config file
	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		return Config{}, fmt.Errorf("could not read config file: %v", err)
	}

	// initialize a config object
	conf := Config{}

	// unmarshal the yaml
	err = yaml.UnmarshalStrict(buf, &conf)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshal yaml: %v", err)
	}

	// expand ~ in config paths
	for i, cheatpath := range conf.Cheatpaths {

		expanded, err := homedir.Expand(cheatpath.Path)
		if err != nil {
			return Config{}, fmt.Errorf("failed to expand ~: %v", err)
		}

		conf.Cheatpaths[i].Path = expanded
	}

	// if an editor was not provided in the configs, look to envvars
	if conf.Editor == "" {
		if os.Getenv("VISUAL") != "" {
			conf.Editor = os.Getenv("VISUAL")
		} else if os.Getenv("EDITOR") != "" {
			conf.Editor = os.Getenv("EDITOR")
		} else {
			return Config{}, fmt.Errorf("no editor set")
		}
	}

	// if a chroma style was not provided, set a default
	if conf.Style == "" {
		conf.Style = "bw"
	}

	// if a chroma formatter was not provided, set a default
	if conf.Formatter == "" {
		conf.Formatter = "terminal16m"
	}

	return conf, nil
}
