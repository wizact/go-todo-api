package user

import (
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	repository "github.com/wizact/go-todo-api/internal/user/adapters/repository"
	appService "github.com/wizact/go-todo-api/internal/user/application/service"
	userDomain "github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	eventPort "github.com/wizact/go-todo-api/internal/user/ports/event"
	pubsubinfra "github.com/wizact/go-todo-api/pkg/eventlibrary/pubsub"
	event "github.com/wizact/go-todo-api/pkg/eventlibrary/user/event"
)

// NewModule assembles legacy environment-backed adapters and panics on configuration errors.
// Deprecated: use NewModuleWithDependencies and construct adapters at application bootstrap.
func NewModule(useDatabase bool, config appService.RegistrationConfig) *Module {
	m, err := NewModuleWithDependencies(Dependencies{Repository: instantiateUserRepository(useDatabase), EventPublisher: instantiateUserEventPublisher()}, WithPublicBaseURL(config.PublicBaseURL))
	if err != nil {
		panic(err)
	}
	return m
}

func instantiateUserRepository(useDatabase bool) userRepository {
	var userRepo userRepository

	if useDatabase {
		rf := dbinfra.SQLiteRepositoryFactory[repository.SQLiteUserRepository, *repository.SQLiteUserRepository]{}
		repo, err := rf.Create()
		if err != nil {
			panic(err)
		}
		userRepo = repo
	} else {

		ua := make([]aggregate.User, 0)
		userRepo = repository.NewMemoryUserRepository(ua)
	}
	return userRepo
}

func instantiateUserEventPublisher() eventPort.UserEventPublisher {
	nf := pubsubinfra.NATSClientFactory[event.UserEventClient, userDomain.UserRegisteredEvent, *event.UserEventClient]{}
	uec, err := nf.Create()
	if err != nil {
		panic(err)
	}

	return uec
}
