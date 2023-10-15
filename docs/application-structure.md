
# Internal

## api
`api` folder is where http handlers and middlewares are configured. This layer is analogous to the `Frameworks and Drivers` of Clean Architecture and is the outermost layer of the app.

This layer should stay as thin as possible an utilise the input ports in the [interface/controller](#interface) layer to communicate with each domain.

### handlers
An example of input adapter where http handlers for the api can live. Handler connects to the the domains using the inpot ports.

### middleware
Http handler middlewares such as logging, rate limiting, error handling.

## Domain Modules
Each domain has it's own folder in the internal folder.

This is a snapshot of a domain strcuture as the higher level folders can have dependency on lower level folders, but not the other way around:
```
├── internal
│   ├── user
│   │   ├── infrastructure
│   │   │       ├── UserDBRepository.go
│   │   │       └── UserMemoryRepository.go
│   │   ├── interfaces
│   │   │       └── controller
│   │   │       └── presenter
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
│   │   │       └── role.go
```

### Domain
Domain folder in each module includes few sub folders, grouped by their hexagonal architecture layers:

#### Enterprise Business Rules
This is where all value objects, entities, aggregates are located:

##### `model`
This folder contains our domain `Entities` and `Value Objects`. It is essential to make sure our enterprise business rules are validated here as much as possible and avoid anemic objects. The enterprise business rules that are validated here are applicable to the whole domain and enterprise and are essential to the integrity of our domain and data.

##### `aggregate`
Aggregate has one aggregate root entity, `user` in this [case](../internal/user/domain/aggregate/user.go). All the changes to the other sub-entities (i.e. `role`) should be done through the aggregate root entity by referencing aggrgate root entity id.

Aggregate Root and Repository relationship should be strictly 1:1. This is because on write, using one repository makes it possible to maintain domain variants. This can be additionally enforced by a transaction and/or unit of work implementation. One read, however, there is a possiblity to use separate means to read data since that would not change the state of data assuming read does not have any side effect.

#### Application Business Rules
##### `repository`
Define one repository per aggregate root to enforice business invaritants while writing the objects to the storage. The interface for the repository is defined here, so that it can be referenced and implemented in infrastructure module to be aligned with clean architecture constraints.

##### `services`
These are domain services responsible for applying application use cases by glueing our aggregates and repositories together. This layer should not be impacted by changes in the database layer, or external depencies.

The business rules validated here as opposed to the validation in model layer are more application to the specific application business rules and use cases.

#### interface
interface layer is the port between outside world (API, UI, Views, Database) and our domain (Use Cases) and vice versa. Essentially it consists of [two types of interfaces](https://crosp.net/blog/software-architecture/clean-architecture-part-2-the-clean-architecture/):

* Use Case Input Ports (e.g. Controller): A button click, or API handler invocation will invoke a method that implements the input adapter (interface) to pass the control to the Use Case layer.
* Use Case Output Ports (e.g. Database, Presenter): Use case layer, after processing the request, invokes a method of implementation of the interface to hand over the control flow to the views (i.e. html rendere view)

To be pragmatic, we merged the input/output ports into one `controller` folder. Additionally, `views` can be placed here too. 

Adapters live here too. Mapping between api models / database entities to  domain entities will happen here. It means layers from use cases inwards should not know about anything from interface layer outwards.

#### Infrastructure 
Infrastructure layer is an adapter to the outer world. In our architecture the concrette implementation of repository interface defined in the Application Business Rules layer is implemented here. These repositories can be used for CRUD operations on data stores, or other external APIs. It is a common practice to have more than one implementation of the repository interface here for different scenarios (i.e. SQL vs in-memory store for testing purposes)

##### Repository
Structs in this module are the implementations of the repository interface in the `UserRepository` and has a 1:1 relation with the root aggregate `user`

#### Other
##### Application / Services
Application services are different from domain services in a sense that they are not required as part of the lifecylce of a domain. However there are necessary for the application lifecycle. An example of that is session or cart management.