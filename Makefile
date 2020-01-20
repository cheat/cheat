# paths
makefile := $(realpath $(lastword $(MAKEFILE_LIST)))
root_dir := $(shell dirname $(makefile))
cmd_dir  := $(root_dir)/cmd/cheat
dist_dir := $(root_dir)/dist

# executables
CAT    := cat
COLUMN := column
CTAGS  := ctags
GO     := go
GREP   := grep
LINT   := revive
MKDIR  := mkdir -p
RM     := rm
SCC    := scc
SED    := sed
SORT   := sort
ZIP    := zip -m

# build flags
BUILD_FLAGS  := -ldflags="-s -w" -mod vendor
GOBIN        :=

# NB: this is a  kludge to specify the desired build targets. This information
# would "naturally" be best structured as an array of structs, but lacking that
# capability, we're condensing that information into strings which we will
# later split.
#
# Format: <architecture>/<os>/<arm-version>/<executable-name>
.PHONY: $(RELEASES)
RELEASES :=                               \
	amd64/darwin/0/cheat-darwin-amd64       \
	amd64/linux/0/cheat-linux-amd64         \
	amd64/windows/0/cheat-windows-amd64.exe \
	arm/linux/5/cheat-linux-arm5            \
	arm/linux/6/cheat-linux-arm6            \
	arm/linux/7/cheat-linux-arm7

# macros to unpack the above
parts = $(subst /, ,$@)
arch  = $(word 1, $(parts))
os    = $(word 2, $(parts))
arm   = $(word 3, $(parts))
bin   = $(word 4, $(parts))


## build: builds an executable for your architecture
.PHONY: build
build: clean generate
	$(GO) build $(BUILD_FLAGS) -o $(dist_dir)/cheat $(cmd_dir)

## build-release: builds release executables
.PHONY: build-release
build-release: $(RELEASES)

.PHONY: generate
generate:
	$(GO) generate $(cmd_dir)

.PHONY: $(RELEASES)
$(RELEASES): clean generate check
ifeq ($(arch),arm)
	GOARCH=$(arch) GOOS=$(os) GOARM=$(arm) $(GO) build $(BUILD_FLAGS) -o $(dist_dir)/$(bin) $(cmd_dir) && \
	$(ZIP) $(dist_dir)/$(bin).zip $(dist_dir)/$(bin)
else
	GOARCH=$(arch) GOOS=$(os) $(GO) build $(BUILD_FLAGS) -o $(dist_dir)/$(bin) $(cmd_dir) && \
	$(ZIP) $(dist_dir)/$(bin).zip $(dist_dir)/$(bin)
endif

## install: builds and installs cheat on your PATH
.PHONY: install
install:
	$(GO) install $(BUILD_FLAGS) $(GOBIN) $(cmd_dir) 

$(dist_dir):
	$(MKDIR) $(dist_dir)

## clean: removes compiled executables
.PHONY: clean
clean: $(dist_dir)
	$(RM) -f $(dist_dir)/*

## distclean: removes the tags file
.PHONY: distclean
distclean:
	$(RM) $(root_dir)/tags

## setup: installs revive (linter) and scc (sloc tool)
.PHONY: setup
setup:
	GO111MODULE=off $(GO) get -u github.com/boyter/scc github.com/mgechev/revive

## sloc: counts "semantic lines of code"
.PHONY: sloc
sloc:
	$(SCC) --exclude-dir=vendor

## tags: builds a tags file
.PHONY: tags
tags:
	$(CTAGS) -R $(root_dir) --exclude=$(root_dir)/vendor

## vendor: downloads, tidies, and verifies dependencies
.PHONY: vendor
vendor: lint # kludge: revive appears to complain if the vendor directory disappears while a lint is running
	$(GO) mod vendor && $(GO) mod tidy && $(GO) mod verify

## fmt: runs go fmt
.PHONY: fmt
fmt:
	$(GO) fmt $(root_dir)/...

## lint: lints go source files
.PHONY: lint
lint:
	$(LINT) -exclude $(root_dir)/vendor/... $(root_dir)/... && \
	$(GO) vet $(root_dir)/...

## test: runs unit-tests
.PHONY: test
test: 
	$(GO) test $(root_dir)/...

## check: formats, lints, vendors, and run unit-tests
.PHONY: check
check: fmt lint vendor test

## help: displays this help text
.PHONY: help
help:
	@$(CAT) $(makefile) | \
	$(SORT)             | \
	$(GREP) "^##"       | \
	$(SED) 's/## //g'   | \
	$(COLUMN) -t -s ':'
