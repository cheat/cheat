#!/bin/bash

# TODO: this script has been made obsolete by the Makefile, yet downstream
# package managers plausibly rely on it for compiling locally. Remove this file
# after downstream maintainers have had time to modify their packages to simply
# invoke `make` in the project root.

# locate the cheat project root
BINDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APPDIR=$(readlink -f "$BINDIR/..")

# compile the executable
cd $APPDIR

make
