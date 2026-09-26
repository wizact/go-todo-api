#!/bin/sh

set -eu

expected_go_series="1.27"

module_go_version="$(go list -m -f '{{.GoVersion}}')"
case "${module_go_version}" in
    "${expected_go_series}"|"${expected_go_series}."*) ;;
    *)
        echo "go.mod requires Go ${module_go_version}; want Go ${expected_go_series}.x" >&2
        exit 1
        ;;
esac

toolchain_go_version="$(go env GOVERSION)"
case "${toolchain_go_version}" in
    "go${expected_go_series}"|"go${expected_go_series}."*) ;;
    *)
        echo "active toolchain is ${toolchain_go_version}; want Go ${expected_go_series}.x" >&2
        exit 1
        ;;
esac
