#!/bin/sh

set -o errexit
set -o nounset
set -o pipefail

if [ -z "${OS:-}" ]; then
    echo "OS must be set"
    exit 1
fi
if [ -z "${ARCH:-}" ]; then
    echo "ARCH must be set"
    exit 1
fi
if [ -z "${VERSION:-}" ]; then
    echo "VERSION must be set"
    exit 1
fi

export CGO_ENABLED=0
export GOARCH="${ARCH}"
export GOOS="${OS}"
export GO111MODULE=on

if [[ "${DEBUG:-}" == 1 ]]; then
    # Debugging - disable optimizations and inlining
    gogcflags="all=-N -l"
    goasmflags=""
    goldflags=""
else
    # Not debugging - trim paths, disable symbols and DWARF.
    goasmflags="all=-trimpath=$(pwd)"
    gogcflags="all=-trimpath=$(pwd)"
    goldflags="-s -w"
fi

always_ldflags="-X $(go list -m)/pkg/version.Version=${VERSION}"
go install                                                        \
    -installsuffix "static"                                       \
    -gcflags="${gogcflags}"                                       \
    -asmflags="${goasmflags}"                                     \
    -ldflags="${always_ldflags} ${goldflags} -extldflags -static" \
    "$@"