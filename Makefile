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
BUILDTAGS :=
# set to 1 for debugging
DBG ?=

ALL_PLATFORMS ?= linux/amd64

GOFLAGS ?=
# Because we store the module cache locally.
GOFLAGS := $(GOFLAGS) -modcacherw

# Do not echo recipes.
MAKEFLAGS += -s

OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))



# Directories that we need created to build binaries.
BUILD_DIRS := out/$(OS)_$(ARCH)					  \
			  .go/bin/$(OS)_$(ARCH)               \
			  .go/cache                           \
              .go/pkg
			  

REGISTRY := "docker.pkg.github.com/wizact/go-todo-api/"
 
VERSION ?= $(shell git describe --tags --always --dirty)

# The binaries to build (just the basenames)
BINS ?= server db-migration

BIN_EXTENSION :=
ifeq ($(OS), windows)
  BIN_EXTENSION := .exe
endif

OUTBINS = $(foreach bin,$(BINS),bin/$(OS)_$(ARCH)/$(bin)$(BIN_EXTENSION))

XOUTBINS = $(foreach bin,$(BINS),$(bin)$(OS)_$(ARCH)$(BIN_EXTENSION))

db-migration_cgo = 1
server_cgo = 1

$(foreach bin,$(BINS),$(eval $(strip   \
    xbuild-$(bin): OUTBIN = $(bin)_$(OS)_$(ARCH)$(BIN_EXTENSION) \
)))

xbuild-%:| $(BUILD_DIRS) # @HELP TODO
	echo $@
	echo "building $(firstword $(subst _, ,$*)) for $(OS)/$(ARCH)"
	docker run                                                  \
		-ti                                                     \
		--rm                                                    \
		-v $$(pwd):/src                                         \
		-w /src                                                 \
		-v $$(pwd)/.go/cache:/.cache                            \
		--env GOCACHE="/.cache/gocache"                         \
		--env GOMODCACHE="/.cache/gomodcache"                   \
		--env ARCH="$(ARCH)"                                    \
		--env OS="$(OS)"                                        \
		--env BUILDTAGS="$(BUILDTAGS)"							\
		--env CGO=$($(firstword $(subst _, ,$*))_cgo)			\
		--env VERSION="$(VERSION)"                              \
		--env GOFLAGS="$(GOFLAGS)"                              \
		--env DEBUG="$(DBG)"                                    \
		--env OUTDIR=".go/bin/$(OS)_$(ARCH)"					\
		--env NAME=$(firstword $(subst _, ,$*))					\
		$(BUILD_IMAGE)                                          \
		./build/xbuild.sh

xbuild-directory:
	mkdir -p $$(pwd)/.go/bin/$(OS)_$(ARCH)

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
	echo -ne "binary: $(OUTBIN) "
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

build-%:
	$(MAKE) build                         \
	    --no-print-directory              \
	    GOOS=$(firstword $(subst _, ,$*)) \
	    GOARCH=$(lastword $(subst _, ,$*))

all-build: # @HELP builds binaries for all platforms
all-build: $(addprefix build-, $(subst /,_, $(ALL_PLATFORMS)))

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

.PHONY: clean-bins
clean-bins: # @HELP clear all the files in the out bin folder
clean-bins:
	echo $@
	rm -rf bin/

.PHONY: help
help: # @HELP prints this message
help:
	echo "VARIABLES:"
	echo "  OS = $(OS)"
	echo "  ARCH = $(ARCH)"
	echo "  GO_VERSION = $(GO_VERSION)"
	echo "  VERSION = $(VERSION)"
	echo "  BUILD_IMAGE = $(BUILD_IMAGE)"
	echo "  GOFLAGS = $(GOFLAGS)"
	echo "  REGISTRY = $(REGISTRY)"
	echo
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
		| awk '                                   \
			BEGIN {FS = ": *# *@HELP"};           \
			{ printf "  %-30s %s\n", $$1, $$2 };  \
		'