# Module scaffolding engine

The generator provides bundled architecture templates, a CLI, and an explicit
`go generate` entry point. The engine
creates architecture and wiring shells; domain models, business rules, ports,
service graphs, and HTTP contracts remain module-specific.

## Usage

Run from the repository root:

```sh
# Makefile entry points:
make preview-module MODULE_NAME=todo
make gen-module MODULE_NAME=todo

# Direct CLI:
go run ./cmd/modulegen -name todo -dry-run
go run ./cmd/modulegen -name todo

# Equivalent creation through the repository generation directive:
MODULE_NAME=todo go generate -run modulegen .
```

`-root` defaults to the current directory. The root must contain the Go module
and an existing `internal` directory. The CLI obtains the repository import path
using `go list -m`; it does not parse or modify `go.mod`.

Bundled templates are embedded in the generator and do not depend on the
working directory. They create an editable module constructor, role-specific
package documentation, and a module README. No domain entities, CRUD, endpoints,
repository methods, event types, or deployment adapters are invented.

Use `-templates /path/to/templates` to replace the bundle with trusted custom
templates. A template directory contains `.tmpl` files whose relative paths
become module-relative output paths. For
example, `module.go.tmpl` becomes `internal/todo/module.go`. Other files are
ignored. Files ending in `.bundle.tmpl` use Go named template definitions: each
definition name is its module-relative output path. Bundle top-level content is
not an output file. Duplicate output paths are rejected. Symlink template files
are rejected. Templates are trusted input.

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

### Repeatable acceptance checks

```sh
go test ./cmd/modulegen -run TestAcceptance -v
go test ./...
go vet ./...
go test -race ./cmd/modulegen ./internal/scaffolding
```

Acceptance tests build the real CLI, preview twice, generate, compile every
new package, edit a generated file, and verify repeat generation preserves it.
They exercise `todo` and `inventory2` under different Go module paths and compare
all pre-existing files, including application bootstrap, the user module,
`go.mod`, and a migration. Separate isolated fixtures exercise the repository's
actual `go generate` directive and both Makefile targets. These checks require
Go and Make; they need no database, NATS server, Docker, or provider credentials.

Fixtures live only in temporary directories and copy generator sources through
an explicit allowlist. Running these tests does not scaffold into the checkout.
Command subprocesses use bounded timeouts and disable workspace discovery for
fixture isolation. Existing unit tests cover malformed templates, unsafe paths,
invalid names, symlink containment, and partial-write cleanup.

## Generated structure and ownership

```text
internal/todo/
  module.go
  README.md
  domain/{model,aggregate,usecase}/doc.go
  application/service/doc.go
  ports/usecase/doc.go
  ports/{repository,event}/doc.go
  ports/service/doc.go
  adapters/controller/doc.go
  adapters/controller/model/doc.go
  adapters/http/doc.go
  adapters/repository/doc.go
```

`module.go` supplies explicit `Dependencies`, private configuration, an `Option`
type, and defaults → ordered options → validation → assembly. Dependency fields,
optional settings and setters, validation rules, service graphs, and accessors
remain developer-owned insertion points. The initial validation methods accept
empty structs because no collaborators or settings have been invented.

Generated files are editable source, not files to regenerate after domain work.
Remove unused package shells when a module does not need those capabilities.
The generator never changes application bootstrap, migrations, or other modules.

`MODULE_NAME` is required for the generation directive. To preview file paths,
use the CLI's `-dry-run`; `go generate -n` prints the command without running it.
Running `go generate ./...` also selects this directive and therefore requires
`MODULE_NAME`; use `-run modulegen .` to avoid unrelated generation directives.
