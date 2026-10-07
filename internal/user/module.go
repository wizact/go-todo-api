package module

import (
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	repository "github.com/wizact/go-todo-api/internal/user/adapters/repositories"
	user_domain "github.com/wizact/go-todo-api/internal/user/domain"
	pubsubinfra "github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	event "github.com/wizact/go-todo-api/pkg/event-library/user/events"

	app_svc "github.com/wizact/go-todo-api/internal/user/application/services"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/services"
	app_svc_port "github.com/wizact/go-todo-api/internal/user/ports/applications"
	usecase_port "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	event_port "github.com/wizact/go-todo-api/internal/user/ports/output/events"
	repository_port "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
)

// A UserModule is the dependency container for the User module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type UserModule struct {
	userRepository     repository_port.UserRepository
	userEventPublisher event_port.UserEventPublisher
	appRegistrationSvc app_svc_port.Registration
	userAccountUseCase usecase_port.UserAccountUseCase
}

type userRepository interface {
	repository_port.UserRepository
	repository_port.RegistrationRepository
}

// New UserModule is the factory method for the UserModule container
func NewUserModule(useDatabase bool, registrationConfig app_svc.RegistrationConfig) *UserModule {
	userRepo := instantiateUserRepository(useDatabase)
	userEventPublisher := instantiateUserEventPublisher()
	userAccountUseCase := instantiateUserAccountUseCase(userRepo, userEventPublisher)
	appSvc := instantiateAppSvc(userAccountUseCase, userRepo, registrationConfig)
	return &UserModule{
		userRepository:     userRepo,
		userEventPublisher: userEventPublisher,
		appRegistrationSvc: appSvc,
		userAccountUseCase: userAccountUseCase,
	}
}

func instantiateUserRepository(useDatabase bool) userRepository {
	var userRepo userRepository

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

func instantiateUserEventPublisher() event_port.UserEventPublisher {
	nf := pubsubinfra.NatsClientFactory[event.UserEventClient, user_domain.UserRegisteredEvent, *event.UserEventClient]{}
	uec, err := nf.Get()
	if err != nil {
		panic(err)
	}

	return uec
}

func instantiateAppSvc(
	uc usecase_port.UserAccountUseCase,
	repository repository_port.RegistrationRepository,
	config app_svc.RegistrationConfig,
) app_svc_port.Registration {
	return app_svc.NewRegistration(uc, repository, config)
}

func instantiateUserAccountUseCase(r repository_port.UserRepository, ev event_port.UserEventPublisher) usecase_port.UserAccountUseCase {
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
