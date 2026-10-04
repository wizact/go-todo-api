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

- [Application architecture](./docs/application-architecture.md)
- [Application structure](./docs/application-structure.md)
- [Registration verification flow](./docs/registrationverification.md)

## Development

- [Development environment](./docs/dev.md)
- [Database migrations](./docs/database-migration.md)
- [Environment variables](./docs/environment-variables.md)
