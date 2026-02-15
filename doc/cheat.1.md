% CHEAT(1) | General Commands Manual

NAME
====

**cheat** — create and view command-line cheatsheets

SYNOPSIS
========

| **cheat** \[options] \[_CHEATSHEET_]

DESCRIPTION
===========
**cheat** allows you to create and view interactive cheatsheets on the
command-line. It was designed to help remind \*nix system administrators of
options for commands that they use frequently, but not frequently enough to
remember.

OPTIONS
=======

--init
: Print a config file to stdout.

--conf
: Display the config file path.

-a, --all
: Search among all cheatpaths.

-b, --brief
: List cheatsheets without file paths.

-c, --colorize
: Colorize output.

-d, --directories
: List cheatsheet directories.

-e, --edit=_CHEATSHEET_
: Open _CHEATSHEET_ for editing.

-l, --list
: List available cheatsheets.

-p, --path=_PATH_
: Filter only to sheets found on path _PATH_.

-r, --regex
: Treat search _PHRASE_ as a regular expression.

-s, --search=_PHRASE_
: Search cheatsheets for _PHRASE_.

-t, --tag=_TAG_
: Filter only to sheets tagged with _TAG_.

-T, --tags
: List all tags in use.

-v, --version
: Print the version number.

--rm=_CHEATSHEET_
: Remove (deletes) _CHEATSHEET_.


EXAMPLES
========

To view the foo cheatsheet:
: cheat _foo_

To edit (or create) the foo cheatsheet:
: cheat -e _foo_

To edit (or create) the foo/bar cheatsheet on the 'work' cheatpath:
: cheat -p _work_ -e _foo/bar_

To view all cheatsheet directories:
: cheat -d

To list all available cheatsheets:
: cheat -l

To briefly list all cheatsheets whose titles match 'apt':
: cheat -b _apt_

To list all tags in use:
: cheat -T

To list available cheatsheets that are tagged as 'personal':
: cheat -l -t _personal_

To search for 'ssh' among all cheatsheets, and colorize matches:
: cheat -c -s _ssh_

To search (by regex) for cheatsheets that contain an IP address:
: cheat -c -r -s _'(?:[0-9]{1,3}\.){3}[0-9]{1,3}'_

To remove (delete) the foo/bar cheatsheet:
: cheat --rm _foo/bar_

To view the configuration file path:
: cheat --conf


FILES
=====

Configuration
-------------
**cheat** is configured via a YAML file that is conventionally named
_conf.yaml_.  **cheat** will search for _conf.yaml_ in varying locations,
depending upon your platform:

### Linux, OSX, and other Unixes ###

1. **CHEAT_CONFIG_PATH**
2. **XDG_CONFIG_HOME**/cheat/conf.yaml
3. **$HOME**/.config/cheat/conf.yml
4. **$HOME**/.cheat/conf.yml

### Windows ###

1. **CHEAT_CONFIG_PATH**
2. **APPDATA**/cheat/conf.yml
3. **PROGRAMDATA**/cheat/conf.yml

**cheat** will search in the order specified above. The first _conf.yaml_
encountered will be respected.

If **cheat** cannot locate a config file, it will ask if you'd like to generate
one automatically. Alternatively, you may also generate a config file manually
by running **cheat --init** and saving its output to the appropriate location
for your platform.


Cheatpaths
----------
**cheat** reads its cheatsheets from "cheatpaths", which are the directories in
which cheatsheets are stored. Cheatpaths may be configured in _conf.yaml_, and
viewed via **cheat -d**.

For detailed instructions on how to configure cheatpaths, please refer to the
comments in conf.yml.


Autocompletion
--------------
Autocompletion scripts for **bash**, **zsh**, and **fish** are available for
download:

- <https://github.com/cheat/cheat/blob/master/scripts/cheat.bash>
- <https://github.com/cheat/cheat/blob/master/scripts/cheat.fish>
- <https://github.com/cheat/cheat/blob/master/scripts/cheat.zsh>

The **bash** and **zsh** scripts provide optional integration with **fzf**, if
the latter is available on your **PATH**.

The installation process will vary per system and shell configuration, and thus
will not be discussed here.


ENVIRONMENT
===========

**CHEAT_CONFIG_PATH**

: The path at which the config file is available. If **CHEAT_CONFIG_PATH** is
set, all other config paths will be ignored.

**CHEAT_USE_FZF**

: If set, autocompletion scripts will attempt to integrate with **fzf**.

RETURN VALUES
=============

0. Successful termination

1. Application error

2. Cheatsheet(s) not found


BUGS
====

See GitHub issues: <https://github.com/cheat/cheat/issues>


AUTHOR
======

Christopher Allen Lane <chris@chris-allen-lane.com>


SEE ALSO
========

**fzf(1)**

