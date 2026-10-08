package user

import (
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	repository "github.com/wizact/go-todo-api/internal/user/adapters/repository"
	userDomain "github.com/wizact/go-todo-api/internal/user/domain"
	pubsubinfra "github.com/wizact/go-todo-api/pkg/eventlibrary/pubsub"
	event "github.com/wizact/go-todo-api/pkg/eventlibrary/user/event"

	appService "github.com/wizact/go-todo-api/internal/user/application/service"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/service"
	usecasePort "github.com/wizact/go-todo-api/internal/user/ports/input/usecase"
	eventPort "github.com/wizact/go-todo-api/internal/user/ports/output/event"
	repositoryPort "github.com/wizact/go-todo-api/internal/user/ports/output/repository"
	appServicePort "github.com/wizact/go-todo-api/internal/user/ports/service"
)

// A Module is the dependency container for the User module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type Module struct {
	userRepository      repositoryPort.UserRepository
	userEventPublisher  eventPort.UserEventPublisher
	registrationService appServicePort.Registration
	userAccountUseCase  usecasePort.UserAccountUseCase
}

type userRepository interface {
	repositoryPort.UserRepository
	repositoryPort.RegistrationRepository
}

// New Module is the factory method for the Module container
func NewModule(useDatabase bool, registrationConfig appService.RegistrationConfig) *Module {
	userRepo := instantiateUserRepository(useDatabase)
	userEventPublisher := instantiateUserEventPublisher()
	userAccountUseCase := instantiateUserAccountUseCase(userRepo, userEventPublisher)
	appSvc := instantiateApplicationService(userAccountUseCase, userRepo, registrationConfig)
	return &Module{
		userRepository:      userRepo,
		userEventPublisher:  userEventPublisher,
		registrationService: appSvc,
		userAccountUseCase:  userAccountUseCase,
	}
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

func instantiateApplicationService(
	uc usecasePort.UserAccountUseCase,
	repository repositoryPort.RegistrationRepository,
	config appService.RegistrationConfig,
) appServicePort.Registration {
	return appService.NewRegistration(uc, repository, config)
}

func instantiateUserAccountUseCase(r repositoryPort.UserRepository, ev eventPort.UserEventPublisher) usecasePort.UserAccountUseCase {
	return usecase.NewUserAccount(r, ev)
}

// UserAccountUseCase returns the concrete implementation of the UserAccountUseCase
func (m *Module) UserAccountUseCase() usecasePort.UserAccountUseCase {
	return m.userAccountUseCase
}

// UserRegistrationAppService returns the concrete implementation of the Registration service
func (m *Module) UserRegistrationAppService() appServicePort.Registration {
	return m.registrationService
}

// Done cleans up all the underlying resources for a graceful shotdown
func (m *Module) Done() {
	m.registrationService.Done()
}
