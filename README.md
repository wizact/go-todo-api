# Go TODO API

Go TODO API is a reference implementation of a modular monolith using Clean Architecture, hexagonal architecture, and Domain-Driven Design (DDD). It keeps independently owned domain modules in one deployable application while enforcing dependency boundaries in code.

The user module demonstrates the core patterns:

- domain entities, value objects, and aggregates own business state and invariants;
- application and domain services coordinate use cases;
- input and output ports define the contracts at architectural boundaries;
- HTTP, SQLite, NATS, and email integrations remain outer adapters;
- one concrete repository owns each aggregate write boundary; and
- read projections may be separate, but aggregate and projection writes that represent one business operation are atomic.

The project uses Go 1.27 with the `go1.27.1` toolchain.

## Architecture and design

- [Application architecture](./docs/applicationarchitecture.md)
- [Application structure](./docs/applicationstructure.md)
- [Registration verification flow](./docs/registrationverification.md)

## Scaffold a module

Preview and create an architecture shell with the local Go toolchain:

```sh
make preview-module MODULE_NAME=todo
make gen-module MODULE_NAME=todo
go test ./internal/todo/...
```

The scaffold creates package roles and an editable constructor with explicit
`Dependencies`, private configuration, ordered functional options, validation,
and service assembly insertion points. Define domain models, ports, workflows,
adapters, and HTTP contracts for the new module. Select concrete adapters and
mount routes at application bootstrap; generation does not register the module.
Existing modules are never overwritten, including after developer edits.

Direct alternatives are `go run ./cmd/modulegen -name todo -dry-run` and
`MODULE_NAME=todo go generate -run modulegen .`. Generation is explicit and is
not part of build or test. See [module scaffolding](./docs/modulescaffolding.md)
for the generated layout, custom templates, safety guarantees, and acceptance
checks.

## Development

- [Development environment](./docs/dev.md)
- [Database migrations](./docs/databasemigration.md)
- [Environment variables](./docs/environmentvariables.md)
