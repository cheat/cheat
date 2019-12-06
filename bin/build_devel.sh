#!/bin/bash

# locate the cheat project root
BINDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APPDIR=$(readlink -f "$BINDIR/..")

# update the vendored dependencies
go mod vendor && go mod tidy

# compile the executable
cd "$APPDIR/cmd/cheat"
go clean && go generate && go build -mod vendor
mv "$APPDIR/cmd/cheat/cheat" "$APPDIR/dist/cheat"

# display a build checksum
md5sum "$APPDIR/dist/cheat"
