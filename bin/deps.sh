#!/bin/bash

# This script installs all Go dependencies required for
# building `cheat` locally.

go get -u github.com/alecthomas/chroma
go get -u github.com/davecgh/go-spew/spew
go get -u github.com/docopt/docopt-go
go get -u github.com/mgutz/ansi
go get -u github.com/mitchellh/go-homedir
go get -u github.com/tj/front
