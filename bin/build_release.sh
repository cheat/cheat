#!/bin/bash

# locate the cheat project root
BINDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APPDIR=$(readlink -f "$BINDIR/..")

# update the vendored dependencies
go mod vendor && go mod tidy

# build embeds
cd "$APPDIR/cmd/cheat"
go clean && go generate

# amd64/darwin
env GOOS=darwin  GOARCH=amd64 go build -mod vendor -o \
  "$APPDIR/dist/cheat-darwin-amd64"  "$APPDIR/cmd/cheat"

# amd64/linux
env GOOS=linux   GOARCH=amd64 go build -mod vendor -o \
  "$APPDIR/dist/cheat-linux-amd64"   "$APPDIR/cmd/cheat"

# amd64/windows
env GOOS=windows GOARCH=amd64 go build -mod vendor -o \
  "$APPDIR/dist/cheat-win-amd64.exe" "$APPDIR/cmd/cheat"

# arm7/linux
env GOOS=linux GOARCH=arm GOARM=7 go build -mod vendor -o \
  "$APPDIR/dist/cheat-linux-arm7" "$APPDIR/cmd/cheat"

# arm6/linux
env GOOS=linux GOARCH=arm GOARM=6 go build -mod vendor -o \
  "$APPDIR/dist/cheat-linux-arm6" "$APPDIR/cmd/cheat"

# arm5/linux
env GOOS=linux GOARCH=arm GOARM=5 go build -mod vendor -o \
  "$APPDIR/dist/cheat-linux-arm5" "$APPDIR/cmd/cheat"
