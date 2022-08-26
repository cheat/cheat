package main

// Code generated .* DO NOT EDIT.

import (
	"strings"
)

func configs() string {
	return strings.TrimSpace(`---
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

# Cheatpaths are paths at which cheatsheets are available on your local
# filesystem.
#
# It is useful to sort cheatsheets into different cheatpaths for organizational
# purposes. For example, you might want one cheatpath for community
# cheatsheets, one for personal cheatsheets, one for cheatsheets pertaining to
# your day job, one for code snippets, etc.
#
# Cheatpaths are scoped, such that more "local" cheatpaths take priority over
# more "global" cheatpaths. (The most global cheatpath is listed first in this
# file; the most local is listed last.) For example, if there is a 'tar'
# cheatsheet on both global and local paths, you'll be presented with the local
# one by default. ('cheat -p' can be used to view cheatsheets from alternative
# cheatpaths.)
#
# Cheatpaths can also be tagged as "read only". This instructs cheat not to
# automatically create cheatsheets on a read-only cheatpath. Instead, when you
# would like to edit a read-only cheatsheet using 'cheat -e', cheat will
# perform a copy-on-write of that cheatsheet from a read-only cheatpath to a
# writeable cheatpath.
#
# This is very useful when you would like to maintain, for example, a
# "pristine" repository of community cheatsheets on one cheatpath, and an
# editable personal reponsity of cheatsheets on another cheatpath.
#
# Cheatpaths can be also configured to automatically apply tags to cheatsheets
# on certain paths, which can be useful for querying purposes.
# Example: 'cheat -t work jenkins'.
#
# Community cheatsheets must be installed separately, though you may have
# downloaded them automatically when installing 'cheat'. If not, you may
# download them here:
#
# https://github.com/cheat/cheatsheets
cheatpaths:
  # Cheatpath properties mean the following:
  #   'name': the name of the cheatpath (view with 'cheat -d', filter with 'cheat -p')
  #   'path': the filesystem path of the cheatsheet directory (view with 'cheat -d')
  #   'tags': tags that should be automatically applied to sheets on this path
  #   'readonly': shall user-created ('cheat -e') cheatsheets be saved here?
  - name: community
    path: COMMUNITY_PATH
    tags: [ community ]
    readonly: true

  # If you have personalized cheatsheets, list them last. They will take
  # precedence over the more global cheatsheets.
  - name: personal
    path: PERSONAL_PATH
    tags: [ personal ]
    readonly: false

  # While it requires no configuration here, it's also worth noting that
  # cheat will automatically append directories named '.cheat' within the
  # current working directory to the 'cheatpath'. This can be very useful if
  # you'd like to closely associate cheatsheets with, for example, a directory
  # containing source code.
  #
  # Such "directory-scoped" cheatsheets will be treated as the most "local"
  # cheatsheets, and will override less "local" cheatsheets. Similarly,
  # directory-scoped cheatsheets will always be editable ('readonly: false').
`)
}
