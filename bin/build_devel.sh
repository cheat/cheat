#!/bin/bash

# locate the cheat project root
BINDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APPDIR=$(readlink -f "$BINDIR/..")

# compile the executable
cd "$APPDIR/cmd/cheat"
go clean && go generate && go build
mv "$APPDIR/cmd/cheat/cheat" "$APPDIR/dist/cheat"

# display a build checksum
md5sum "$APPDIR/dist/cheat"
