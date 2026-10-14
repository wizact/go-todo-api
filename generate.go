// Package gotodoapi hosts repository-level generation directives.
package gotodoapi

// Set MODULE_NAME explicitly; generation never runs as part of build or test.
// Example: MODULE_NAME=todo go generate -run modulegen .
//go:generate go run ./cmd/modulegen -name $MODULE_NAME
