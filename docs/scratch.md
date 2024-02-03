

## Interface Adapters
* cmd/api: Driver Actor
* domain: Repositories

## Use Cases
* services: Application Business Rules

## Enterprise Business Rules (entities, contexts, aggregates)
* aggregate
* entity
* valueobject



Input And Output Ports are interfaces to the UseCase and Repositories

Driving And Driven Adapters are Controllers/Views/Presentations and Repositories that use / to be used through input and output adapters


Driving Adapter > Input Port > UseCase > Output Port > Driven Adapter
HttpController > UserServiceUseCase (Interface) > UserService > UserRepository (Interface) > UserMemoryRepository

Adapters:
convert data from the format most convenient for the use cases and entities, to the format most convenient for some external agency such as the Database or the Web. It is this layer, for example, that will wholly contain the MVC architecture of a GUI. The Presenters, Views, and Controllers all belong in here.



Ports must only be consist of interface type. 
From inner circles nothing can reference outer circles expect the port folder