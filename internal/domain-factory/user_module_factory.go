package domainfactory

import userModule "github.com/wizact/go-todo-api/internal/user"

var userModuleInstance *userModule.UserModule

// CreateUserModule creates a new user module
func CreateNewUserModule() *userModule.UserModule {
	if userModuleInstance != nil {
		return userModuleInstance
	}

	userModuleInstance = userModule.NewUserModule(true)
	return userModuleInstance
}
