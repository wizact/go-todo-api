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

# build configuration in the format of <bin>_<config> = <value>
db-migration_cgo = 1
server_cgo = 1

ALL_PLATFORMS ?= linux/amd64

GOFLAGS ?=
# Because we store the module cache locally.
GOFLAGS := $(GOFLAGS) -modcacherw

# Do not echo recipes.
MAKEFLAGS += -s

# OS/ARCH can be provided using the command line, e.g. OS=linux ARCH=arm64
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

# Directories that we need created to build binaries. the final artefacts will live in out/<OS>_<ARCH>
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

# ARTEFACTS = $(foreach bin,$(BINS),$(bin)$(OS)_$(ARCH)$(BIN_EXTENSION))

$(foreach bin,$(BINS),$(eval $(strip   \
    build-$(bin): OUTBIN = $(bin)_$(OS)_$(ARCH)$(BIN_EXTENSION) \
)))

build-%: # @HELP run the build command for each bins (BINS). usage: make build-<bin> OS=linux ARCH=arm64
build-%:| $(BUILD_DIRS) 
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
		--env OUTNAME=$(OUTBIN)									\
		--env NAME=$(firstword $(subst _, ,$*))					\
		$(BUILD_IMAGE)                                          \
		./build/build.sh


$(foreach bin,$(BINS),$(eval $(strip   \
    artefact-$(bin): OUTBIN = $(bin)_$(OS)_$(ARCH)$(BIN_EXTENSION) \
)))

artefact-%: # @HELP copies the artefact from .go/<OS>_<ARCH>/<bin> to out/<OS>_<ARCH>/<bin> if they are newer
artefact-%: | $(BUILD_DIRS)
	if ! cmp -s .go/bin/$(OS)_$(ARCH)/$(OUTBIN) ./out/$(OS)_$(ARCH)/$(OUTBIN); then  	\
		mv .go/bin/$(OS)_$(ARCH)/$(OUTBIN) out/$(OS)_$(ARCH)/$(OUTBIN);            		\
		date >out/$(OS)_$(ARCH)/$@.stamp;                              					\
		echo;                                  											\
	else                                       											\
		echo "(cached)";                       											\
	fi

$(BUILD_DIRS):
	mkdir -p $@

all-artefact: # @HELP copies the all artefacts from .go/<OS>_<ARCH>/ to out/<OS>_<ARCH>/ if they are newer
all-artefact: $(addprefix artefact-, $(BINS))

all-build: # @HELP builds binaries for bins defined in BINS var. usage: make all-build OS=linux ARCH=arm64
all-build: $(addprefix build-, $(BINS))

shell: # @HELP launches a shell in the containerized build environment
shell: | $(BUILD_DIRS)
	echo "# launching a shell in the containerized build environment"
	docker run                                                  \
		-ti                                                     \
		--rm                                                    \
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

.PHONY: run-server
run-server: # @HELP runs the http server on the localhost port of 9000
run-server:
	go run ./cmd/server/*.go start-server -address=localhost -port=9000 -tls=false

.PHONY: run-db-migration
run-db-migration: # @HELP generates the new db resources and run the migration cmd
run-db-migration: gen-db-resource
	TODOAPI_DBPATH=${PREFIX}/db/todo.db && \
		go run ./cmd/db-migration/*.go

.PHONY: gen-db-resource
gen-db-resource: # @HELP creates a resourcefile and embeds migration scripts in the go file
gen-db-resource:
	rm -f ${PREFIX}/db/resourcefile.go && \
	${PREFIX}/build/migration.sh

.PHONY: clean-bins
clean: # @HELP clear all the files in the out and .go folder
clean:
	echo $@
	rm -rf .go/*
	rm -rf out/*

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