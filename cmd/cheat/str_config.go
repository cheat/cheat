package main

// Code generated .* DO NOT EDIT.

import (
	"strings"
)

func configs() string {
	return strings.TrimSpace(`---
# The editor to use with 'cheat -e <sheet>'. Defaults to $EDITOR or $VISUAL.
editor: vim

# Should 'cheat' always colorize output?
colorize: true

# Which 'chroma' colorscheme should be applied to the output?
# Options are available here:
#   https://github.com/alecthomas/chroma/tree/master/styles
style: monokai

# Which 'chroma' "formatter" should be applied?
# One of: "terminal", "terminal256", "terminal16m"
formatter: terminal16m

# The paths at which cheatsheets are available. Tags associated with a cheatpath
# are automatically attached to all cheatsheets residing on that path.
#
# Whenever cheatsheets share the same title (like 'tar'), the most local
# cheatsheets (those which come later in this file) take precedent over the
# less local sheets. This allows you to create your own "overides" for
# "upstream" cheatsheets.
#
# But what if you want to view the "upstream" cheatsheets instead of your own?
# Cheatsheets may be filtered via 'cheat -t <tag>' in combination with other
# commands. So, if you want to view the 'tar' cheatsheet that is tagged as
# 'community' rather than your own, you can use: cheat tar -t community
cheatpaths:

  # Paths that come earlier are considered to be the most "global", and will
  # thus be overridden by more local cheatsheets. That being the case, you
  # should probably list community cheatsheets first.
  #
  # Note that the paths and tags listed below are just examples. You may freely
  # change them to suit your needs.
  - name: community
    path: ~/.dotfiles/cheat/community
    tags: [ community ]
    readonly: true

  # Maybe your company or department maintains a repository of cheatsheets as
  # well. It's probably sensible to list those second.
  - name: work
    path: ~/.dotfiles/cheat/work
    tags: [ work ]
    readonly: false

  # If you have personalized cheatsheets, list them last. They will take
  # precedence over the more global cheatsheets.
  - name: personal
    path: ~/.dotfiles/cheat/personal
    tags: [ personal ]
    readonly: false
`)
}
