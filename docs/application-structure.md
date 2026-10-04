# Application structure

The repository is organized first by executable or domain module, then by architectural role. The user module is the reference implementation.

```text
cmd/
  server/                         process entry point
  db-migration/                   migration entry point
db/
  migrations/                     ordered SQLite migrations
internal/
  api/
    handlers/                     HTTP routing and boundary validation
    middleware/                   HTTP cross-cutting behavior
  infra/
    db/                           database connection infrastructure
  user/
    adapters/
      controllers/                HTTP-to-application mapping
      repositories/               SQLite and memory adapters
    application/
      registration/               application-owned verification state
      services/                   registration workflow orchestration
    domain/
      aggregates/                 user aggregate root
      factories/                  aggregate construction
      models/                     entities and value objects
      services/                   user account use cases
    ports/
      applications/               application-service contracts
      input/use_cases/             driving use-case contracts
      output/events/              event-publisher contracts
      output/repositories/        persistence contracts
    module.go                     module composition root
pkg/
  communication/                  email and event-listener module
  event-library/                  NATS event adapter
```

## Domain

The domain contains business concepts that remain meaningful without HTTP, SQLite, NATS, or email.

### Models

`domain/models` contains entities and value objects. Constructors validate enterprise-wide invariants, and rehydration functions restore already validated persistent state.

### Aggregates

`domain/aggregates` contains aggregate roots. Changes to aggregate-owned state happen through aggregate behavior. For example, `User.VerifyRegistration` activates the user and marks the email as verified.

The aggregate does not own registration tokens, token hashing, expiration, or delivery. Those are application workflow concerns.

### Domain services

`domain/services` implements use cases that coordinate domain behavior. `UserAccountService` depends on repository and event ports rather than concrete adapters.

## Application

Application code coordinates workflows that are required by this application but are not intrinsic state of the domain model.

`application/registration.Verification` represents a short-lived credential using a user ID, bcrypt digest, and expiry. `application/services.Registration` creates and verifies that credential while asking the user aggregate to perform the resulting state transition.

## Ports

Ports are interfaces at architectural boundaries:

- input ports describe use cases invoked by controllers;
- application ports expose application services to listeners or controllers;
- repository ports describe persistence required by inner services; and
- event ports describe event publication required by the domain service.

Ports may reference types from the domain or application layer. Concrete adapters depend on and implement ports; inner layers do not import adapter packages.

Keep ports small and consumer-oriented. Multiple narrow port interfaces may be implemented by one concrete adapter when the same aggregate write boundary owns their operations.

## Adapters

Controllers validate and translate external requests before invoking an input or application port. They do not contain business rules.

Repositories map between domain/application types and storage models. Both SQLite and memory adapters implement equivalent behavior. The SQLite adapter owns the transaction that keeps aggregate state, projections, and registration credential consumption consistent.

## API and infrastructure

`internal/api` and `cmd` are the outermost application layers. They assemble HTTP routes, middleware, and process lifecycle behavior. `internal/infra` and reusable `pkg` packages integrate databases, messaging, and communication providers.

External input is validated at these boundaries. Errors returned to clients are sanitized while wrapped errors retain internal diagnostic context.

## Adding a module

Use the smallest structure that preserves the dependency rule:

1. Define the domain model and aggregate boundary.
2. Define consumer-owned ports for external interactions.
3. Implement use cases or application workflows against those ports.
4. Add thin adapters for HTTP, persistence, messaging, or other drivers.
5. Wire concrete adapters only in the module composition root.

Avoid creating layers, interfaces, or projections that the module does not yet need.
