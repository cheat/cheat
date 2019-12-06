#!/bin/bash

# locate the cheat project root
BINDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APPDIR=$(readlink -f "$BINDIR/..")

# update the vendored dependencies
go mod vendor && go mod tidy

# build embeds
cd "$APPDIR/cmd/cheat"
go clean && go generate

# compile AMD64 for Linux, OSX, and Windows
env GOOS=darwin  GOARCH=amd64 go build -mod vendor -o \
  "$APPDIR/dist/cheat-darwin-amd64"  "$APPDIR/cmd/cheat"

env GOOS=linux   GOARCH=amd64 go build -mod vendor -o \
  "$APPDIR/dist/cheat-linux-amd64"   "$APPDIR/cmd/cheat"

env GOOS=windows GOARCH=amd64 go build -mod vendor -o \
  "$APPDIR/dist/cheat-win-amd64.exe" "$APPDIR/cmd/cheat"
