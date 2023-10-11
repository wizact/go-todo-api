

## Interface Adapters
* cmd/api: Driver Actor
* domain: Repositories

## Use Cases
* services: Application Business Rules

## Enterprise Business Rules (entities, contexts, aggregates)
* aggregate
* entity
* valueobject


├── internal
│   ├── user
│   │   ├── infrastructure
│   │   │       ├── UserDBRepository.go
│   │   │       └── UserMemoryRepository.go
│   │   ├── presentation
│   │   ├── application
│   │   │   └── service
│   │   │       ├── authorization.go
│   │   │       └── registration.go
│   │   ├── domain     
│   │   │   ├── service    
│   │   │   ├── factory    
│   │   │   ├── repository
│   │   │   │    └── user.go
│   │   │   └── model  
│   │   │       ├── user.go
│   │   │       ├── group.go
│   │   │       └── role.go
