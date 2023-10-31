NAME := go-todo-api
MIGRATIONNAME := migration
PKG := github.com/wizact/$(NAME)
SHELL := /bin/bash
GO := go
BUILDTAGS :=
PREFIX?=$(shell pwd)
OUTDIR := ${PREFIX}/out

OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

# Directories that we need created to build binaries.
BUILD_DIRS := $(OUTDIR)/server/$(OS)_$(ARCH)             \
              $(OUTDIR)/db-migration/$(OS)_$(ARCH)

# VERSION := $(shell cat VERSION.txt)
REGISTRY := "docker.pkg.github.com/wizact/go-todo-api/"

GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)

ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif
ifeq ($(GITCOMMIT),)
    GITCOMMIT := ${GITHUB_SHA}
endif

CTIMEVAR=-X $(PKG)/version.GITCOMMIT=$(GITCOMMIT) -X $(PKG)/version.VERSION=$(VERSION)
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"

.PHONY: gen-db-resource
gen-db-resource:
	cd ./db/ && rm -f ./resourcefile.go && \
	./migration.sh

.PHONY: build-migrate-db
build-migrate-db: | $(BUILD_DIRS)
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


$(BUILD_DIRS):
	mkdir -p $@

.PHONY: clean-bin
clean-bin:
	echo $(OUTDIR)
	rm -rf $(OUTDIR)

# dbver=$$(echo './cmd/db-migration/2_alter_table.up.sql' | sed 's/^\([0-9]*\)_.*\.sql$$/\1/') && \
# echo $$dbver
# 's/^\([0-9]*\)_.*\.sql$$/\1/')
# 
# cd ./cmd/db-migration/ && $(GO) build \
# -tags "static_build"
# -o $(OUTDIR)/db-migration/$(MIGRATIONNAME) .
# 		
# echo $$dbver

