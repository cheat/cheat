package installer

import (
	"path/filepath"
	"strings"

	"github.com/cheat/cheat/internal/config"
)

// cheatsheetDirs returns the community, personal, and work cheatsheet
// directory paths derived from a config file path.
func cheatsheetDirs(confpath string) (community, personal, work string) {
	confdir := filepath.Dir(confpath)
	community = filepath.Join(confdir, "cheatsheets", "community")
	personal = filepath.Join(confdir, "cheatsheets", "personal")
	work = filepath.Join(confdir, "cheatsheets", "work")
	return
}

// ExpandTemplate replaces placeholder tokens in the config template with
// platform-appropriate paths derived from confpath.
func ExpandTemplate(configs string, confpath string) string {
	community, personal, work := cheatsheetDirs(confpath)

	// substitute paths
	configs = strings.ReplaceAll(configs, "COMMUNITY_PATH", community)
	configs = strings.ReplaceAll(configs, "PERSONAL_PATH", personal)
	configs = strings.ReplaceAll(configs, "WORK_PATH", work)

	// locate and set a default pager
	configs = strings.ReplaceAll(configs, "PAGER_PATH", config.Pager())

	// locate and set a default editor
	if editor, err := config.Editor(); err == nil {
		configs = strings.ReplaceAll(configs, "EDITOR_PATH", editor)
	}

	return configs
}

// CommentCommunity comments out the community cheatpath block in the config
// template. This is used when the community cheatsheets directory won't exist
// (either because the user declined to download them, or because the config
// is being output as an example).
func CommentCommunity(configs string, confpath string) string {
	community, _, _ := cheatsheetDirs(confpath)

	return strings.ReplaceAll(configs,
		"  - name: community\n"+
			"    path: "+community+"\n"+
			"    tags: [ community ]\n"+
			"    readonly: true",
		"  #- name: community\n"+
			"  #  path: "+community+"\n"+
			"  #  tags: [ community ]\n"+
			"  #  readonly: true",
	)
}
