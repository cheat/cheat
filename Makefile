# paths
makefile := $(realpath $(lastword $(MAKEFILE_LIST)))
cmd_dir  := ./cmd/cheat
dist_dir := ./dist

# executables
CAT    := cat
COLUMN := column
CTAGS  := ctags
DOCKER := docker
GO     := go
GREP   := grep
GZIP   := gzip --best
LINT   := revive
MAN    := man
MKDIR  := mkdir -p
PANDOC := pandoc
RM     := rm
SCC    := scc
SED    := sed
SORT   := sort
ZIP    := zip -m

docker_image := cheat-devel:latest

# build flags
BUILD_FLAGS  := -ldflags="-s -w" -mod vendor -trimpath
GOBIN        :=
TMPDIR       := /tmp

# release binaries
releases :=                        \
	$(dist_dir)/cheat-darwin-amd64 \
	$(dist_dir)/cheat-linux-386    \
	$(dist_dir)/cheat-linux-amd64  \
	$(dist_dir)/cheat-linux-arm5   \
	$(dist_dir)/cheat-linux-arm6   \
	$(dist_dir)/cheat-linux-arm7   \
	$(dist_dir)/cheat-windows-amd64.exe

## build: build an executable for your architecture
.PHONY: build
build: $(dist_dir) clean fmt lint vet vendor generate man
	$(GO) build $(BUILD_FLAGS) -o $(dist_dir)/cheat $(cmd_dir)

## build-release: build release executables
.PHONY: build-release
build-release: $(releases)

# cheat-darwin-amd64
$(dist_dir)/cheat-darwin-amd64: prepare
	GOARCH=amd64 GOOS=darwin \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-386
$(dist_dir)/cheat-linux-386: prepare
	GOARCH=386 GOOS=linux \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-amd64
$(dist_dir)/cheat-linux-amd64: prepare
	GOARCH=amd64 GOOS=linux \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-arm5
$(dist_dir)/cheat-linux-arm5: prepare
	GOARCH=arm GOOS=linux GOARM=5 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-arm6
$(dist_dir)/cheat-linux-arm6: prepare
	GOARCH=arm GOOS=linux GOARM=6 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-arm7
$(dist_dir)/cheat-linux-arm7: prepare
	GOARCH=arm GOOS=linux GOARM=7 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-windows-amd64
$(dist_dir)/cheat-windows-amd64.exe: prepare
	GOARCH=amd64 GOOS=windows \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(ZIP) $@.zip $@

# ./dist
$(dist_dir):
	$(MKDIR) $(dist_dir)

.PHONY: generate
generate:
	$(GO) generate $(cmd_dir)

## install: build and install cheat on your PATH
.PHONY: install
install: build
	$(GO) install $(BUILD_FLAGS) $(GOBIN) $(cmd_dir) 

## clean: remove compiled executables
.PHONY: clean
clean: $(dist_dir)
	$(RM) -f $(dist_dir)/*

## distclean: remove the tags file
.PHONY: distclean
distclean:
	$(RM) -f tags
	@$(DOCKER) image rm -f $(docker_image)

## setup: install revive (linter) and scc (sloc tool)
.PHONY: setup
setup:
	GO111MODULE=off $(GO) get -u github.com/boyter/scc github.com/mgechev/revive

## sloc: count "semantic lines of code"
.PHONY: sloc
sloc:
	$(SCC) --exclude-dir=vendor

## tags: build a tags file
.PHONY: tags
tags:
	$(CTAGS) -R --exclude=vendor --languages=go

## man: build a man page
# NB: pandoc may not be installed, so we're ignoring this error on failure
.PHONY: man
man:
	-$(PANDOC) -s -t man doc/cheat.1.md -o doc/cheat.1

## vendor: download, tidy, and verify dependencies
.PHONY: vendor
vendor:
	$(GO) mod vendor && $(GO) mod tidy && $(GO) mod verify

## vendor-update: update vendored dependencies
vendor-update:
	$(GO) get -t -u ./... && $(GO) mod vendor

## fmt: run go fmt
.PHONY: fmt
fmt:
	$(GO) fmt ./...

## lint: lint go source files
.PHONY: lint
lint: vendor
	$(LINT) -exclude vendor/... ./...

## vet: vet go source files
.PHONY: vet
vet:
	$(GO) vet ./...

## test: run unit-tests
.PHONY: test
test: 
	$(GO) test ./...

## coverage: generate a test coverage report
.PHONY: coverage
coverage:
	$(GO) test ./... -coverprofile=$(TMPDIR)/cheat-coverage.out && \
	$(GO) tool cover -html=$(TMPDIR)/cheat-coverage.out

## check: format, lint, vet, vendor, and run unit-tests
.PHONY: check
check: | vendor fmt lint vet test

.PHONY: prepare
prepare: | $(dist_dir) clean generate vendor fmt lint vet test

## docker-setup: create a docker image for use during development
.PHONY: docker-setup
docker-setup:
	$(DOCKER) build  -t $(docker_image) -f Dockerfile .

## docker-sh: shell into the docker development container
.PHONY: docker-sh
docker-sh:
	$(DOCKER) run -v $(shell pwd):/app -ti $(docker_image) /bin/ash

## help: display this help text
.PHONY: help
help:
	@$(CAT) $(makefile) | \
	$(SORT)             | \
	$(GREP) "^##"       | \
	$(SED) 's/## //g'   | \
	$(COLUMN) -t -s ':'
