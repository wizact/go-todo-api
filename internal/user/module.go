package module

import (
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	pubsubinfra "github.com/wizact/go-todo-api/internal/infra/pubsub"
	event "github.com/wizact/go-todo-api/internal/user/adapters/events"
	repository "github.com/wizact/go-todo-api/internal/user/adapters/repositories"

	app_svc "github.com/wizact/go-todo-api/internal/user/application/services" // TODO: should be replaced with interface
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/services"
	app_svc_port "github.com/wizact/go-todo-api/internal/user/ports/applications"
	event_port "github.com/wizact/go-todo-api/internal/user/ports/events"
	usecase_port "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	repository_port "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
)

// A UserModule is the dependency container for the User module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type UserModule struct {
	userRepository     repository_port.UserRepository
	userEventClient    event_port.UserEventClient
	appRegistrationSvc app_svc_port.Registration
}

// New UserModule is the factory method for the UserModule container
func NewUserModule(useDatabase bool) *UserModule {
	userRepo := instantiateUserRepository(useDatabase)
	userEventCli := instantiateUserEventClient()
	appSvc := instantiateEventListeners(userEventCli)
	return &UserModule{
		userRepository:     userRepo,
		userEventClient:    userEventCli,
		appRegistrationSvc: appSvc,
	}
}

func instantiateUserRepository(useDatabase bool) repository_port.UserRepository {
	var userRepo repository_port.UserRepository

	if useDatabase {
		rf := dbinfra.SqliteRepositoryFactory[repository.UserSqliteRepository, *repository.UserSqliteRepository]{}
		repo, err := rf.Get()
		if err != nil {
			panic(err)
		}
		userRepo = repo
	} else {

		ua := make([]aggregate.User, 0)
		userRepo = repository.NewUserMemoryRepository(ua)
	}
	return userRepo
}

func instantiateUserEventClient() event_port.UserEventClient {
	nf := pubsubinfra.NatsClientFactory[event.UserEventClient, aggregate.User, *event.UserEventClient]{}
	uec, err := nf.Get()
	if err != nil {
		panic(err)
	}

	return uec
}

func instantiateEventListeners(ev event_port.UserEventClient) app_svc_port.Registration {
	appSvc := app_svc.NewRegisteration(ev)
	err := appSvc.NewUserRegisteredListener()
	if err != nil {
		panic(err)
	}

	return appSvc
}

func (u *UserModule) ResolveUserAccountUseCase() usecase_port.UserAccountUseCase {
	return usecase.NewUserAccountService(u.userRepository, u.userEventClient)
}

// Done cleans up all the underlying resources for a graceful shotdown
func (u *UserModule) Done() {
	u.appRegistrationSvc.Done()
}
