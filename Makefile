# paths
makefile := $(realpath $(lastword $(MAKEFILE_LIST)))
cmd_dir  := ./cmd/cheat
dist_dir := ./dist

# parallel jobs for build-release (can be overridden)
JOBS ?= 8

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
	$(dist_dir)/cheat-darwin-arm64 \
	$(dist_dir)/cheat-linux-386    \
	$(dist_dir)/cheat-linux-amd64  \
	$(dist_dir)/cheat-linux-arm5   \
	$(dist_dir)/cheat-linux-arm6   \
	$(dist_dir)/cheat-linux-arm64  \
	$(dist_dir)/cheat-linux-arm7   \
	$(dist_dir)/cheat-netbsd-amd64  \
	$(dist_dir)/cheat-openbsd-amd64  \
	$(dist_dir)/cheat-solaris-amd64  \
	$(dist_dir)/cheat-windows-amd64.exe

## build: build an executable for your architecture
.PHONY: build
build: | clean $(dist_dir) fmt lint vet vendor man
	$(GO) build $(BUILD_FLAGS) -o $(dist_dir)/cheat $(cmd_dir)

## build-release: build release executables
# Runs prepare once, then builds all binaries in parallel
# Override jobs with: make build-release JOBS=16
.PHONY: build-release
build-release: prepare
	$(MAKE) -j$(JOBS) $(releases)

# cheat-darwin-amd64
$(dist_dir)/cheat-darwin-amd64:
	GOARCH=amd64 GOOS=darwin \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-darwin-arm64
$(dist_dir)/cheat-darwin-arm64:
	GOARCH=arm64 GOOS=darwin \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-386
$(dist_dir)/cheat-linux-386:
	GOARCH=386 GOOS=linux \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-amd64
$(dist_dir)/cheat-linux-amd64:
	GOARCH=amd64 GOOS=linux \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-arm5
$(dist_dir)/cheat-linux-arm5:
	GOARCH=arm GOOS=linux GOARM=5 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-arm6
$(dist_dir)/cheat-linux-arm6:
	GOARCH=arm GOOS=linux GOARM=6 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-linux-arm7
$(dist_dir)/cheat-linux-arm7:
	GOARCH=arm GOOS=linux GOARM=7 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz
	
# cheat-linux-arm64
$(dist_dir)/cheat-linux-arm64:
	GOARCH=arm64 GOOS=linux \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-netbsd-amd64
$(dist_dir)/cheat-netbsd-amd64:
	GOARCH=amd64 GOOS=netbsd \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-openbsd-amd64
$(dist_dir)/cheat-openbsd-amd64:
	GOARCH=amd64 GOOS=openbsd \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-plan9-amd64
$(dist_dir)/cheat-plan9-amd64:
	GOARCH=amd64 GOOS=plan9 \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-solaris-amd64
$(dist_dir)/cheat-solaris-amd64:
	GOARCH=amd64 GOOS=solaris \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(GZIP) $@ && chmod -x $@.gz

# cheat-windows-amd64
$(dist_dir)/cheat-windows-amd64.exe:
	GOARCH=amd64 GOOS=windows \
	$(GO) build $(BUILD_FLAGS) -o $@ $(cmd_dir) && $(ZIP) $@.zip $@ -j

# ./dist
$(dist_dir):
	$(MKDIR) $(dist_dir)

# .tmp
.tmp:
	$(MKDIR) .tmp

## install: build and install cheat on your PATH
.PHONY: install
install: build
	$(GO) install $(BUILD_FLAGS) $(GOBIN) $(cmd_dir) 

## clean: remove compiled executables
.PHONY: clean
clean:
	$(RM) -f $(dist_dir)/*
	$(RM) -rf .tmp

## distclean: remove the tags file
.PHONY: distclean
distclean:
	$(RM) -f tags
	@$(DOCKER) image rm -f $(docker_image)

## setup: install revive (linter) and scc (sloc tool)
.PHONY: setup
setup:
	$(GO) install github.com/boyter/scc@latest
	$(GO) install github.com/mgechev/revive@latest

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
.PHONY: vendor-update
vendor-update:
	$(GO) get -t -u ./... && $(GO) mod vendor && $(GO) mod tidy && $(GO) mod verify

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

## test-integration: run integration tests (requires network)
.PHONY: test-integration
test-integration:
	$(GO) test -tags=integration -count=1 ./...

## test-all: run all tests (unit and integration)
.PHONY: test-all
test-all: test test-integration

## test-fuzz: run quick fuzz tests for security-critical functions
.PHONY: test-fuzz
test-fuzz:
	@./build/fuzz.sh 15s

## test-fuzz-long: run extended fuzz tests (10 minutes each)
.PHONY: test-fuzz-long
test-fuzz-long:
	@./build/fuzz.sh 10m

## coverage: generate a test coverage report
.PHONY: coverage
coverage: .tmp
	$(GO) test ./... -coverprofile=.tmp/cheat-coverage.out && \
	$(GO) tool cover -html=.tmp/cheat-coverage.out -o .tmp/cheat-coverage.html && \
	echo "Coverage report generated: .tmp/cheat-coverage.html" && \
	(sensible-browser .tmp/cheat-coverage.html 2>/dev/null || \
	 xdg-open .tmp/cheat-coverage.html 2>/dev/null || \
	 open .tmp/cheat-coverage.html 2>/dev/null || \
	 echo "Please open .tmp/cheat-coverage.html in your browser")

## coverage-text: show test coverage by function in terminal
.PHONY: coverage-text
coverage-text: .tmp
	$(GO) test ./... -coverprofile=.tmp/cheat-coverage.out && \
	$(GO) tool cover -func=.tmp/cheat-coverage.out | $(SORT) -k3 -n

## benchmark: run performance benchmarks
.PHONY: benchmark
benchmark: .tmp
	$(GO) test -tags=integration -bench=. -benchtime=10s -benchmem ./cmd/cheat | tee .tmp/benchmark-latest.txt && \
	$(RM) -f cheat.test

## benchmark-cpu: run benchmarks with CPU profiling
.PHONY: benchmark-cpu
benchmark-cpu: .tmp
	$(GO) test -tags=integration -bench=. -benchtime=10s -cpuprofile=.tmp/cpu.prof ./cmd/cheat && \
	$(RM) -f cheat.test && \
	echo "CPU profile saved to .tmp/cpu.prof" && \
	echo "View with: go tool pprof -http=:8080 .tmp/cpu.prof"

## benchmark-mem: run benchmarks with memory profiling
.PHONY: benchmark-mem
benchmark-mem: .tmp
	$(GO) test -tags=integration -bench=. -benchtime=10s -benchmem -memprofile=.tmp/mem.prof ./cmd/cheat && \
	$(RM) -f cheat.test && \
	echo "Memory profile saved to .tmp/mem.prof" && \
	echo "View with: go tool pprof -http=:8080 .tmp/mem.prof"

## check: format, lint, vet, vendor, and run unit-tests
.PHONY: check
check: | vendor fmt lint vet test

.PHONY: prepare
prepare: | clean $(dist_dir) vendor fmt lint vet test

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
