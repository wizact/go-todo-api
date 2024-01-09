package module

import (
	repository "github.com/wizact/go-todo-api/internal/user/adapters/repositories"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/services"
	usecase_port "github.com/wizact/go-todo-api/internal/user/ports/input/use_cases"
	repository_port "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
)

// A UserModule is the dependency container for the User module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type UserModule struct {
	userRepository repository_port.UserRepository
}

// New UserModule is the factory method for the UserModule container
func NewUserModule(useDatabase bool) *UserModule {
	var userRepo repository_port.UserRepository

	if useDatabase {
		userRepo = repository.NewUserSqlLiteRepository()
	} else {
		// for testing purposes
		ua := make([]aggregate.User, 0)
		userRepo = repository.NewUserMemoryRepository(ua)
	}

	return &UserModule{
		userRepository: userRepo,
	}
}

func (u *UserModule) ResolveUserAccountUseCase() usecase_port.UserAccountUseCase {
	return usecase.NewUserAccountService(u.userRepository)
}
