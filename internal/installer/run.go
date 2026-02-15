package installer

import (
	"fmt"
	"os"

	"github.com/cheat/cheat/internal/config"
	"github.com/cheat/cheat/internal/repo"
)

// Run runs the installer
func Run(configs string, confpath string) error {

	// expand template placeholders with platform-appropriate paths
	configs = ExpandTemplate(configs, confpath)

	// determine cheatsheet directory paths
	community, personal, work := cheatsheetDirs(confpath)

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
		fmt.Printf("Cloning community cheatsheets to %s.\n", community)
		if err := repo.Clone(community); err != nil {
			return fmt.Errorf("failed to clone cheatsheets: %v", err)
		}
	} else {
		configs = CommentCommunity(configs, confpath)
	}

	// always create personal and work directories
	for _, dir := range []string{personal, work} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// the config file does not exist, so we'll try to create one
	if err = config.Init(confpath, configs); err != nil {
		return fmt.Errorf("failed to create config file: %v", err)
	}

	return nil
}
