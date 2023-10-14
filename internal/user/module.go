package module

import (
	"github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	"github.com/wizact/go-todo-api/internal/user/domain/repository"
	"github.com/wizact/go-todo-api/internal/user/domain/service"
	infraRepository "github.com/wizact/go-todo-api/internal/user/infrastructure/repository"
)

// A UserModule is the dependency container for the User module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type UserModule struct {
	userRepository         repository.UserRepository
	emailGatewayRepository repository.EmailGatewayRepository
	// userAccountService       service.UserAccountService
	// emailVerificationService service.EmailVerificationService
}

// New UserModule is the factory method for the UserModule container
func NewUserModule(useDatabase, useEmailGateway bool) *UserModule {
	var userRepo repository.UserRepository
	var emailGatewayRepo repository.EmailGatewayRepository

	if useDatabase {
		userRepo = infraRepository.NewUserSqlLiteRepository()
	} else {
		// for testing purposes
		ua := make([]aggregate.User, 0)
		userRepo = infraRepository.NewUserMemoryRepository(ua)
	}

	if useEmailGateway {
		emailGatewayRepo = infraRepository.NewEmailGatewayRepository()
	} else {
		emailGatewayRepo = infraRepository.NewFakeEmailGatewayRepository()
	}

	return &UserModule{
		userRepository:         userRepo,
		emailGatewayRepository: emailGatewayRepo,
	}
}

func (u *UserModule) ResolveUserAccountService() service.UserAccountService {
	return *service.NewUserAccountService(u.userRepository, u.ResolveEmailVerificationService())
}

func (u *UserModule) ResolveEmailVerificationService() service.EmailVerificationService {
	return *service.NewEmailVerificationService(u.emailGatewayRepository)
}
