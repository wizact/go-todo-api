# Module scaffolding engine

Phase 4 provides the generation engine and CLI. Bundled architecture templates
and the repository's `go generate` entry point are added in Phase 5. The engine
creates architecture and wiring shells; domain models, business rules, ports,
service graphs, and HTTP contracts remain module-specific.

## Current CLI

Run from the repository root with a trusted template directory:

```sh
go run ./cmd/modulegen -name todo -templates /path/to/templates -dry-run
go run ./cmd/modulegen -name todo -templates /path/to/templates
```

`-root` defaults to the current directory. The root must contain the Go module
and an existing `internal` directory. The CLI obtains the repository import path
using `go list -m`; it does not parse or modify `go.mod`.

This phase does not bundle a production scaffold. A template directory contains
`.tmpl` files whose relative paths become module-relative output paths. For
example, `module.go.tmpl` becomes `internal/todo/module.go`. Other files are
ignored. Symlink template files are rejected. Templates are trusted input.

Templates receive `ModuleName` and `ImportPath`. Rendering uses `text/template`
with missing keys rejected; Go output is formatted with `go/format`. No timestamp
or random value is introduced. Output paths are sorted for deterministic plans.

## Safety and boundaries

`BuildPlan` validates and renders everything without writing files. Names must
be lowercase Go package names without delimiters; keywords, `main`, and `test`
are rejected. Existing module directories, files, and symlinks are rejected.

`WritePlan` rechecks output paths and exclusively creates the module directory
and files. Go's `os.Root` confines access to the repository and module directory,
including symlink resolution. Traversal, absolute paths, duplicate paths, and
backslash paths are rejected. A failed write removes the new module directory;
cleanup errors are returned alongside the original error. Existing modules are
never overwritten or removed.

Generation is not an atomic transaction visible to other processes: files become
visible during creation. Do not concurrently edit a module while it is being
created. Bootstrap files, migrations, and `go.mod` are not modified.

## Validation

Tests cover deterministic rendering, formatting and template failures, invalid
names and paths, missing configuration, dry-run behavior, repeat generation,
existing content preservation, escaping symlinks, and partial-write cleanup.
