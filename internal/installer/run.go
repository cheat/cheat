package installer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/repo"
)

// Run runs the installer
func Run(configs string, confpath string) error {

	// determine the appropriate paths for config data and (optional) community
	// cheatsheets based on the user's platform
	confdir := filepath.Dir(confpath)

	// create paths for community and personal cheatsheets
	community := filepath.Join(confdir, "cheatsheets", "community")
	personal := filepath.Join(confdir, "cheatsheets", "personal")

	// set default cheatpaths
	configs = strings.Replace(configs, "COMMUNITY_PATH", community, -1)
	configs = strings.Replace(configs, "PERSONAL_PATH", personal, -1)

	// locate and set a default pager
	configs = strings.Replace(configs, "PAGER_PATH", config.Pager(), -1)

	// locate and set a default editor
	if editor, err := config.Editor(); err == nil {
		configs = strings.Replace(configs, "EDITOR_PATH", editor, -1)
	}

	// prompt the user to download the community cheatsheets
	yes, err := Prompt(
		"Would you like to download the community cheatsheets? [Y/n]",
		true,
	)
	if err != nil {
		return fmt.Errorf("failed to prompt: %v", err)
	}

	// clone the community cheatsheets if so instructed
	if yes {
		// clone the community cheatsheets
		fmt.Printf("Cloning community cheatsheets to %s.\n", community)
		if err := repo.Clone(community); err != nil {
			return fmt.Errorf("failed to clone cheatsheets: %v", err)
		}

		// also create a directory for personal cheatsheets
		fmt.Printf("Cloning personal cheatsheets to %s.\n", personal)
		if err := os.MkdirAll(personal, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// the config file does not exist, so we'll try to create one
	if err = config.Init(confpath, configs); err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}

	return nil
}
