#!/bin/sh
set -eu

echo "Testing..."

sh ./build/checkgoversion.sh
go test ./...
