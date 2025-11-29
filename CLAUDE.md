# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go TODO API implementing Domain-Driven Design (DDD) and Clean Architecture patterns. The application demonstrates a hexagonal architecture with clear separation of concerns across different layers.

## Architecture

The application follows Clean Architecture with these layers:
- **Enterprise Business Rules**: Domain models, entities, value objects, aggregates
- **Application Business Rules**: Domain services, repository interfaces, use cases
- **Interface Adapters**: Controllers, presenters, adapters between layers
- **Infrastructure**: Database implementations, external service adapters
- **Frameworks & Drivers**: HTTP handlers, middleware

Each domain module (e.g., `user`) follows this structure:
```
internal/user/
├── domain/
│   ├── models/          # Entities and value objects
│   ├── aggregates/      # Aggregate roots
│   ├── services/        # Domain services (use cases)
├── application/
│   └── services/        # Application services
├── adapters/
│   ├── controllers/     # HTTP controllers and models
│   └── repositories/    # Repository implementations
├── ports/               # Interface definitions
│   ├── input/          # Use case interfaces
│   └── output/         # Repository interfaces
└── module.go           # Dependency injection container
```

## Development Commands

### Running the Application
```bash
# Start the HTTP server locally on port 9000
make run-server
```

### Building
```bash
# Build all binaries (server and db-migration)
make all-build

# Build specific binary for current OS/ARCH
make build-server
make build-db-migration

# Build for specific platform
make build-server OS=linux ARCH=amd64
```

### Testing
```bash
# Run all tests
make test

# Run tests directly with Go (without Docker)
go test ./...
```

### Database Management
```bash
# Generate embedded migration resources
make gen-db-resource

# Run database migrations (generates resources first)
make run-db-migration
```

### Development Environment
```bash
# Launch development container
devcontainer up --workspace-folder .

# Open shell in containerized build environment
make shell
```

## Key Technologies

- **Web Framework**: Gorilla Mux for HTTP routing
- **Database**: SQLite with GORM ORM
- **Migrations**: golang-migrate for database versioning
- **Events**: NATS for domain event publishing
- **Email**: SendGrid for email notifications
- **CLI**: yacli for command-line interface
- **Testing**: Standard Go testing framework

## Environment Variables

Key environment variables (see `docs/environment-variables.md` for complete list):
- `TODOAPI_DBPATH`: SQLite database file path
- `TODOAPI_NATSURL`: NATS server URL (e.g., `nats:4222`)
- `TODOAPI_SENDGRIDKEY`: SendGrid API key
- `TODOAPI_SENDGRIDFROMNAME`: SendGrid sender name
- `TODOAPI_SENDGRIDFROMEMAIL`: SendGrid sender email

## Testing Domain Events

To test NATS domain events:
```bash
# Subscribe to user registration events
nats subscribe "User.NewUserRegistered" -s nats:4222
```

## Module Structure

The application uses a modular architecture where each domain has its own module with dependency injection. The `UserModule` in `internal/user/module.go` demonstrates the pattern:

- Repository implementations can be swapped between in-memory and database versions
- Event clients are injected for domain event publishing
- Use cases and application services are wired together in the module factory

## Database

- SQLite database located at `db/todo.db`
- Migrations in `db/migrations/` with up/down SQL files
- Migration scripts are embedded into Go using resource generation
- Database connection and repository abstractions in `internal/infra/db/`

## API Structure

- HTTP server setup in `internal/api/http_server.go`
- Route registration follows domain-based organization
- Health check endpoint at `/__health-check`
- User endpoints at `/users`
- HTTP client examples in `http-client/` directory

## Build System

The project uses a sophisticated Docker-based build system:
- All builds run in containerized environment for consistency
- Supports cross-compilation for different OS/ARCH combinations
- Artifacts output to `out/` directory
- Build cache in `.go/` directory for faster subsequent builds