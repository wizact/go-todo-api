.DEFAULT_GOAL := help
NAME := go-todo-api
MIGRATIONNAME := migration
PKG := github.com/wizact/$(NAME)
SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
GO := go
GO_VERSION := 1.19
BUILD_IMAGE := golang:$(GO_VERSION)-alpine
BUILDTAGS :=
PREFIX?=$(shell pwd)
# set to 1 for debugging
DBG ?=

GOFLAGS ?=
# Because we store the module cache locally.
GOFLAGS := $(GOFLAGS) -modcacherw

# Do not echo recipes.
MAKEFLAGS += -s

OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

# Directories that we need created to build binaries.
BUILD_DIRS := bin/$(OS)_$(ARCH)					  \
			  .go/bin/$(OS)_$(ARCH)               \
              .go/bin/$(OS)_$(ARCH)/$(OS)_$(ARCH) \
			  .go/cache                           \
              .go/pkg
			  

REGISTRY := "docker.pkg.github.com/wizact/go-todo-api/"
 
VERSION ?= $(shell git describe --tags --always --dirty)
CTIMEVAR=-X $(PKG)/version.VERSION=$(VERSION)
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"

# The binaries to build (just the basenames)
BINS ?= server db-migration

BIN_EXTENSION :=
ifeq ($(OS), windows)
  BIN_EXTENSION := .exe
endif

OUTBINS = $(foreach bin,$(BINS),bin/$(OS)_$(ARCH)/$(bin)$(BIN_EXTENSION))

build: # @HELP builds the binaries in a container
build: $(OUTBINS)
	echo

$(foreach outbin,$(OUTBINS),$(eval  \
    $(outbin): .go/$(outbin).stamp  \
))

# This is the target definition for all outbins.
$(OUTBINS):
	true

# Each stampfile target can reference an $(OUTBIN) variable.
$(foreach outbin,$(OUTBINS),$(eval $(strip   \
    .go/$(outbin).stamp: OUTBIN = $(outbin)  \
)))

# This is the target definition for all stampfiles.
# This will build the binary under ./.go and update the real binary iff needed.
STAMPS = $(foreach outbin,$(OUTBINS),.go/$(outbin).stamp)
.PHONY: $(STAMPS)
$(STAMPS): go-build
	echo -ne "binary: $(OUTBIN)  "
	if ! cmp -s .go/$(OUTBIN) $(OUTBIN); then  \
		mv .go/$(OUTBIN) $(OUTBIN);            \
		date >$@;                              \
		echo;                                  \
	else                                       \
		echo "(cached)";                       \
	fi

go-build:| $(BUILD_DIRS)
	echo "# building for $(OS)/$(ARCH)"
	docker run                                                  \
		-ti                                                     \
		--rm                                                    \
		-u $$(id -u):$$(id -g)                                  \
		-v $$(pwd):/src                                         \
		-w /src                                                 \
		-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
		-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
		-v $$(pwd)/.go/cache:/.cache                            \
		--env GOCACHE="/.cache/gocache"                         \
		--env GOMODCACHE="/.cache/gomodcache"                   \
		--env ARCH="$(ARCH)"                                    \
		--env OS="$(OS)"                                        \
		--env VERSION="$(VERSION)"                              \
		--env GOFLAGS="$(GOFLAGS)"                              \
		--env DEBUG="$(DBG)"                                    \
		$(BUILD_IMAGE)                                          \
		./build/build.sh ./...

$(BUILD_DIRS):
	mkdir -p $@

shell: # @HELP launches a shell in the containerized build environment
shell: | $(BUILD_DIRS)
	echo "# launching a shell in the containerized build environment"
	docker run                                                  \
		-ti                                                     \
		--rm                                                    \
		-u $$(id -u):$$(id -g)                                  \
		-v $$(pwd):/src                                         \
		-w /src                                                 \
		-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin                \
		-v $$(pwd)/.go/bin/$(OS)_$(ARCH):/go/bin/$(OS)_$(ARCH)  \
		-v $$(pwd)/.go/cache:/.cache                            \
		--env GOCACHE="/.cache/gocache"                         \
		--env GOMODCACHE="/.cache/gomodcache"                   \
		--env ARCH="$(ARCH)"                                    \
		--env OS="$(OS)"                                        \
		--env VERSION="$(VERSION)"                              \
		--env GOFLAGS="$(GOFLAGS)"                              \
		$(BUILD_IMAGE)                                          \
		/bin/sh $(CMD)

.PHONY: gen-db-resource
gen-db-resource: # @HELP creates a resourcefile and embeds migration scripts in the go file
gen-db-resource:
	rm -f ${PREFIX}/db/resourcefile.go && \
	${PREFIX}/build/migration.sh

.PHONY: build-db-migration
build-db-migration: # @HELP gets the latest db version and creates a binary migration file
build-db-migration: | $(BUILD_DIRS) gen-db-resource
	latestup=$$(echo `ls -r ./db/migrations/**.up.sql | head -1`) && \
	if [ -z "$$latestup" ]; then \
		echo 'cannot find any up script' $$latestup;  \
	else \
		echo 'last up command is:' $$latestup; \
	fi && \
	dbver=$$(echo $$latestup | sed -n '/^\.[\\\/]db[\\\/]migrations[\\\/]/ s/[^0-9]*\([0-9]\+\)_*.*up.sql$$/\1/p') && \
	if [ $$dbver -gt 0 ]; then \
		echo 'lastest db version will be:' $$dbver; \
	else \
		echo 'error finding the latest up command version'; \
		exit 1; \
	fi

.PHONY: clean-bins
clean-bins: # @HELP clear all the files in the out bin folder
clean-bins:
	echo $@
	rm -rf bin/
	rm -rf .go/

.PHONY: help
help: # @HELP prints this message
help:
	echo "VARIABLES:"
	echo "  OS = $(OS)"
	echo "  ARCH = $(ARCH)"
	echo "  GO_VERSION = $(GO_VERSION)"
	echo "  VERSION = $(VERSION)"
	echo "  GOLDFlags = $(GO_LDFLAGS_STATIC)"
	echo "  BUILD_IMAGE = $(BUILD_IMAGE)"
	echo
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
		| awk '                                   \
			BEGIN {FS = ": *# *@HELP"};           \
			{ printf "  %-30s %s\n", $$1, $$2 };  \
		'