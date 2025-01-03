package module

import (
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	repository "github.com/wizact/go-todo-api/internal/user/adapters/repositories"
	pubsubinfra "github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	UserDomainEvent "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
	event "github.com/wizact/go-todo-api/pkg/event-library/user/events"

	app_svc "github.com/wizact/go-todo-api/internal/user/application/services"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/services"
	app_svc_port "github.com/wizact/go-todo-api/internal/user/ports/applications"
	usecase_port "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	repository_port "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
	event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/events"
)

// A UserModule is the dependency container for the User module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type UserModule struct {
	userRepository     repository_port.UserRepository
	userEventClient    event_port.UserEventClient
	appRegistrationSvc app_svc_port.Registration
	userAccountUseCase usecase_port.UserAccountUseCase
}

// New UserModule is the factory method for the UserModule container
func NewUserModule(useDatabase bool) *UserModule {
	userRepo := instantiateUserRepository(useDatabase)
	userEventCli := instantiateUserEventClient()
	userAccountUseCase := instantiateUserAccountUseCase(userRepo, userEventCli)
	appSvc := instantiateAppSvc(userEventCli, userAccountUseCase)
	return &UserModule{
		userRepository:     userRepo,
		userEventClient:    userEventCli,
		appRegistrationSvc: appSvc,
		userAccountUseCase: userAccountUseCase,
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
	nf := pubsubinfra.NatsClientFactory[event.UserEventClient, UserDomainEvent.UserDomainEvent, *event.UserEventClient]{}
	uec, err := nf.Get()
	if err != nil {
		panic(err)
	}

	return uec
}

func instantiateAppSvc(ev event_port.UserEventClient, uc usecase_port.UserAccountUseCase) app_svc_port.Registration {
	return app_svc.NewRegisteration(ev, uc)
}

func instantiateUserAccountUseCase(r repository_port.UserRepository, ev event_port.UserEventClient) usecase_port.UserAccountUseCase {
	return usecase.NewUserAccountService(r, ev)
}

// UserAccountUseCase returns the concrete implementation of the UserAccountUseCase
func (u *UserModule) UserAccountUseCase() usecase_port.UserAccountUseCase {
	return u.userAccountUseCase
}

// UserRegistrationAppService returns the concrete implementation of the Registration service
func (u *UserModule) UserRegistrationAppService() app_svc_port.Registration {
	return u.appRegistrationSvc
}

// Done cleans up all the underlying resources for a graceful shotdown
func (u *UserModule) Done() {
	u.appRegistrationSvc.Done()
}
