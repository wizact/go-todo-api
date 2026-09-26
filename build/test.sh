#!/bin/sh
set -eu

echo "Testing..."

sh ./build/check-go-version.sh
go test ./...
