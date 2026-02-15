package main

// configs returns the default configuration template
func configs() string {
	return `---
# The editor to use with 'cheat -e <sheet>'. Defaults to $EDITOR or $VISUAL.
editor: EDITOR_PATH

# Should 'cheat' always colorize output?
colorize: false

# Which 'chroma' colorscheme should be applied to the output?
# Options are available here:
#   https://github.com/alecthomas/chroma/tree/master/styles
style: monokai

# Which 'chroma' "formatter" should be applied?
# One of: "terminal", "terminal256", "terminal16m"
formatter: terminal256

# Through which pager should output be piped?
# 'less -FRX' is recommended on Unix systems
# 'more' is recommended on Windows
pager: PAGER_PATH

# The paths at which cheatsheets are available. Tags associated with a cheatpath
# are automatically attached to all cheatsheets residing on that path.
#
# Whenever cheatsheets share the same title (like 'tar'), the most local
# cheatsheets (those which come later in this file) take precedence over the
# less local sheets. This allows you to create your own "overides" for
# "upstream" cheatsheets.
#
# But what if you want to view the "upstream" cheatsheets instead of your own?
# Cheatsheets may be filtered by 'tags' in combination with the '--tag' flag.
# 
# Example: 'cheat tar --tag=community' will display the 'tar' cheatsheet that
# is tagged as 'community' rather than your own.
#
# Paths that come earlier are considered to be the most "global", and paths
# that come later are considered to be the most "local". The most "local" paths
# take precedence.
#
# See: https://github.com/cheat/cheat/blob/master/doc/cheat.1.md#cheatpaths
cheatpaths:

  # Cheatsheets that are tagged "personal" are stored here by default:
  - name: personal
    path: PERSONAL_PATH
    tags: [ personal ]
    readonly: false

  # Cheatsheets that are tagged "work" are stored here by default:
  - name: work
    path: WORK_PATH
    tags: [ work ]
    readonly: false

  # Community cheatsheets (https://github.com/cheat/cheatsheets):
  # To install: git clone https://github.com/cheat/cheatsheets COMMUNITY_PATH
  - name: community
    path: COMMUNITY_PATH
    tags: [ community ]
    readonly: true

  # You can also use glob patterns to automatically load cheatsheets from all
  # directories that match.
  #
  # Example: overload cheatsheets for projects under ~/src/github.com/example/*/
  #- name: example-projects
  #  path: ~/src/github.com/example/**/.cheat
  #  tags: [ example ]
  #  readonly: true`
}
