// Package config manages application configuration and settings.
//
// The config package provides functionality to:
//   - Load configuration from YAML files
//   - Validate configuration values
//   - Manage platform-specific configuration paths
//   - Handle editor and pager settings
//   - Configure colorization and formatting options
//
// # Configuration Structure
//
// The main configuration file (conf.yml) contains:
//   - Editor preferences
//   - Pager settings
//   - Colorization options
//   - Cheatpath definitions
//   - Formatting preferences
//
// Example configuration:
//
//	---
//	editor: vim
//	colorize: true
//	style: monokai
//	formatter: terminal256
//	pager: less -FRX
//	cheatpaths:
//	  - name: personal
//	    path: ~/cheat
//	    tags: []
//	    readonly: false
//	  - name: community
//	    path: ~/cheat/.cheat
//	    tags: [community]
//	    readonly: true
//
// # Platform-Specific Paths
//
// The package automatically detects configuration paths based on the operating system:
//   - Linux/Unix: $XDG_CONFIG_HOME/cheat/conf.yml or ~/.config/cheat/conf.yml
//   - macOS: ~/Library/Application Support/cheat/conf.yml
//   - Windows: %APPDATA%\cheat\conf.yml
//
// # Environment Variables
//
// The following environment variables are respected:
//   - CHEAT_CONFIG_PATH: Override the configuration file location
//   - CHEAT_USE_FZF: Enable fzf integration when set to "true"
//   - EDITOR: Default editor if not specified in config
//   - VISUAL: Fallback editor if EDITOR is not set
//   - PAGER: Default pager if not specified in config
package config
