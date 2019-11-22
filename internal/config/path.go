package config

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/mitchellh/go-homedir"
)

const (
	configFilePathEnvVar = "CHEAT_CONFIG_PATH"
	configFileName       = "conf.yml"
	configFolder         = "cheat"
)

// PreferredFolderPath returns the default cheat folder path
func PreferredFolderPath() (string, error) {

	switch runtime.GOOS {

	case "darwin":

		// macOS default folder path
		return path.Join("~/Library/Application Support", configFolder), nil

	case "linux":

		// Linux default folder path
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome != "" {
			return path.Join(xdgConfigHome, configFolder), nil
		}
		return path.Join("~/.config", configFolder), nil

	case "windows":

		// Windows default folder path
		folderPath := path.Join(os.Getenv("APPDATA"), configFolder)
		return strings.ReplaceAll(folderPath, "/", "\\"), nil

	default:

		// Unsupported platforms
		return "", fmt.Errorf("unsupported os: %s", runtime.GOOS)

	}
}

// PreferredConfigPath returns the default config file path
func PreferredConfigPath() (string, error) {

	// if CHEAT_CONFIG_PATH is set, return it
	envFilePath := os.Getenv(configFilePathEnvVar)
	if envFilePath != "" {
		return envFilePath, nil
	}

	configFolder, err := PreferredFolderPath()
	if err != nil {
		return "", err
	}

	configPath := path.Join(configFolder, configFileName)
	if runtime.GOOS == "windows" {
		configPath = strings.ReplaceAll(configPath, "/", "\\")
	}

	return configPath, nil
}

// Path returns the config file path
func Path() (string, error) {

	// if CHEAT_CONFIG_PATH is set, return it
	envFilePath, err := getExpandedEnv(configFilePathEnvVar)
	if err != nil {
		return "", err
	}

	if envFilePath != "" {
		return envFilePath, nil
	}

	var paths []string

	switch runtime.GOOS {

	case "darwin":

		homePath := os.Getenv("HOME")

		// macOS config paths
		paths = []string{
			path.Join(homePath, "Library/Application Support", configFolder, configFileName),
			path.Join(os.Getenv("XDG_CONFIG_HOME"), configFolder, configFileName),
			path.Join(homePath, ".config", configFolder, configFileName),
			path.Join(homePath, "."+configFolder, configFileName),
			path.Join("/Library/Application Support", configFolder, configFileName),
			path.Join("/etc", configFolder, configFileName),
		}

	case "linux":

		homePath := os.Getenv("HOME")

		// Linux config paths
		paths = []string{
			path.Join(os.Getenv("XDG_CONFIG_HOME"), configFolder, configFileName),
			path.Join(homePath, ".config", configFolder, configFileName),
			path.Join(homePath, "."+configFolder, configFileName),
			path.Join("/etc", configFolder, configFileName),
		}

	case "windows":

		// Windows config paths
		paths = []string{
			path.Join(os.Getenv("APPDATA"), configFolder, configFileName),
			path.Join(os.Getenv("PROGRAMDATA"), configFolder, configFileName),
		}

		for i, p := range paths {
			paths[i] = strings.ReplaceAll(p, "/", "\\")
		}

	default:

		// Unsupported platforms
		return "", fmt.Errorf("unsupported os: %s", runtime.GOOS)

	}

	// check if the config file exists on any paths
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	// we can't find the config file if we make it this far
	return "", fmt.Errorf("could not locate config file")
}

// getExpandedEnv returns an expanded environment variable if set
func getExpandedEnv(envVar string) (string, error) {

	value := os.Getenv(envVar)
	if value != "" {

		// expand environment variable
		expanded, err := homedir.Expand(value)
		if err != nil {
			return "", fmt.Errorf("failed to expand %s: %v", envVar, err)
		}

		return expanded, nil
	}

	return "", nil
}
