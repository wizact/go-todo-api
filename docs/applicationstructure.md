# Application structure

The repository is organized first by executable or domain module, then by architectural role. The user module is the reference implementation.

```text
cmd/
  server/                         process entry point
  dbmigration/                    migration entry point
db/
  migrations/                     ordered SQLite migrations
internal/
  api/
    handlers/                     process-level health check route
  infra/
    db/                           database connection infrastructure
    httpmiddleware/               shared HTTP response and authorization behavior
  user/
    adapters/
      controller/                 HTTP-to-application mapping
      http/                       user routes, HTTP parsing, and route tests
      repository/                 SQLite and memory adapters
    application/
      registration/               application-owned verification state
      service/                    registration workflow orchestration
    domain/
      aggregate/                  user aggregate root
      factories/                  aggregate construction
      model/                      entities and value objects
      usecase/                    user account use cases
    ports/
      service/                    application-service contracts
      usecase/                    driving use-case contracts
      event/                      event-publisher contracts
      repository/                 persistence contracts
    module.go                     module composition root
pkg/
  communication/                  email and event-listener module
  eventlibrary/                   NATS event adapter
```

## Domain

The domain contains business concepts that remain meaningful without HTTP, SQLite, NATS, or email.

### Models

`domain/model` contains entities and value objects. Constructors validate enterprise-wide invariants, and rehydration functions restore already validated persistent state.

### Aggregates

`domain/aggregate` contains aggregate roots. Changes to aggregate-owned state happen through aggregate behavior. For example, `User.VerifyRegistration` activates the user and marks the email as verified.

The aggregate does not own registration tokens, token hashing, expiration, or delivery. Those are application workflow concerns.

### Domain use cases

`domain/usecase` implements use cases that coordinate domain behavior. `UserAccount` depends on repository and event ports rather than concrete adapters.

## Application

Application code coordinates workflows that are required by this application but are not intrinsic state of the domain model.

`application/registration.Verification` represents a short-lived credential using a user ID, bcrypt digest, and expiry. `application/service.Registration` creates and verifies that credential while asking the user aggregate to perform the resulting state transition.

## Ports

Ports are interfaces at architectural boundaries:

- `ports/usecase` describes use cases invoked by controllers;
- `ports/service` exposes application services to listeners or controllers;
- `ports/repository` describes persistence required by inner services; and
- `ports/event` describes event publication required by domain use cases.

Ports are grouped by role directly under `ports`; there are no input/output directory tiers. The shared event library follows the same layout under `pkg/eventlibrary/ports/event`.

Ports may reference types from the domain or application layer. Concrete adapters depend on and implement ports; inner layers do not import adapter packages.

Keep ports small and consumer-oriented. Multiple narrow port interfaces may be implemented by one concrete adapter when the same aggregate write boundary owns their operations.

## Adapters

User HTTP handlers live in `internal/user/adapters/http`, alongside their route tests and request fixtures. They parse HTTP requests and register routes against Gorilla Mux. Controllers validate and translate those requests before invoking a use-case or application-service port. Neither adapter contains business rules.

`internal/api` mounts module routes and owns server startup and health checks. Shared HTTP response and authorization behavior lives in `internal/infra/httpmiddleware`, so module adapters do not depend on the application bootstrap package. The current authorization placeholder is unchanged by this separation.

Repositories map between domain/application types and storage models. Both SQLite and memory adapters implement equivalent behavior. The SQLite adapter owns the transaction that keeps aggregate state, projections, and registration credential consumption consistent.

The user repository's missing-user errors are part of its port contract. Both adapters return errors matching `domain.ErrUserIDNotFound` for a missing ID during lookup or update, and `domain.ErrUserEmailNotFound` for a missing email during lookup. Consumers use `errors.Is` rather than comparing message strings. The memory adapter also preserves matching its existing adapter-specific errors. A missing email is an ordinary first-registration lookup result, not a persistence failure.

## API and infrastructure

`internal/api` and `cmd` are the outermost application layers. They assemble HTTP routes, middleware, and process lifecycle behavior. `internal/infra` and reusable `pkg` packages integrate databases, messaging, and communication providers.

External input is validated at these boundaries. Errors returned to clients are sanitized while wrapped errors retain internal diagnostic context.

## Adding a module

Use the smallest structure that preserves the dependency rule:

1. Define the domain model and aggregate boundary.
2. Define consumer-owned ports for external interactions.
3. Implement use cases or application workflows against those ports.
4. Add thin adapters for HTTP, persistence, messaging, or other drivers.
5. Select concrete adapters at application bootstrap; assemble module services from explicit dependencies.

Avoid creating layers, interfaces, or projections that the module does not yet need.

### Module dependencies and configuration

The user module exposes `NewModuleWithDependencies(Dependencies, ...Option)`.
Required collaborators are explicit: one repository implements both user and
registration persistence, and an event publisher publishes registration events.
Using the same repository preserves atomic registration completion.

Optional settings use module-local functional options. `WithPublicBaseURL` changes
registration email URLs. Defaults apply first, options apply in order (last wins),
and the final configuration is validated before service assembly. Nil options,
missing dependencies, and invalid URLs return errors. Empty URLs do not restore
defaults. Options configure values only; they perform no I/O.

`internal/api/bootstrap.go` selects SQLite and NATS adapters and returns factory
errors. The injected constructor reads no environment and starts no workers.
`NewModule` remains a deprecated compatibility wrapper with panic-based errors;
its adapter factories are isolated in `internal/user/legacy.go`.

Future scaffolds can reproduce dependency/configuration/constructor conventions.
Each module still defines its own ports, options, services, and domain behavior.

### Scaffolded conventions versus module-owned behavior

Use `make preview-module MODULE_NAME=todo` and `make gen-module MODULE_NAME=todo`
for the reference architecture shell. See [module scaffolding](./modulescaffolding.md)
for usage and verification.

| Scaffolded convention | Developer-owned implementation |
| --- | --- |
| Domain, application, port, and adapter package roles | Entities, invariants, consistency boundaries, and workflows |
| Explicit dependency struct and service container | Dependency fields, service graph, and accessors |
| Defaults, ordered options, and validation sequence | Settings, defaults, option setters, and validation rules |
| HTTP and repository adapter locations | Routes, boundary contracts, storage, and event publication |
| Module documentation and editable package shells | Bootstrap integration, migrations, and business tests |

The generated constructor initially assembles an empty shell. It is not a
ready-to-use service or a generic CRUD module. Add only the capabilities the new
module requires; remove unused package shells. Shared middleware remains outside
individual modules, while domain-specific adapters remain inside their module.
