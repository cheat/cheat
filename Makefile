SHELL     = /bin/sh

# executables
GO       := go
LINTER   := revive
CTAGS    := ctags
GTAGS    := gtags

# paths
MAKEFILE := $(realpath $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(shell dirname $(MAKEFILE))
CMD_DIR  := $(ROOT_DIR)/cmd/cheat
DIST_DIR := $(ROOT_DIR)/dist

all: distclean check
	mkdir -p $(DIST_DIR) &&            \
	cd $(CMD_DIR)        &&            \
	$(GO) mod vendor && $(GO) mod tidy \
	$(GO) fmt $(ROOT_DIR)/...          \
	cd $(CMD_DIR) && $(GO) generate    \
	$(GO) build -mod vendor -o $(DIST_DIR)/cheat

check:
	$(GO) test $(CMD_DIR)/...

clean:
	rm -f $(ROOT_DIR)/GPATH;  \
	rm -f $(ROOT_DIR)/GRTAGS; \
	rm -f $(ROOT_DIR)/GTAGS;  \
	rm -f $(ROOT_DIR)/tags

distclean:
	rm -rf $(DIST_DIR)

dist:
	echo "TODO"

tags:
	$(CTAGS) -R . && $(GTAGS)

install: 
	echo "TODO"

lint:
	@$(LINTER) -exclude vendor/... ./...

.PHONY:         \
	all           \
	check         \
	clean         \
	dist          \
	distclean     \
	install       \
	lint
