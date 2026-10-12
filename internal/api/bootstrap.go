package api

import (
	"fmt"
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	usermodule "github.com/wizact/go-todo-api/internal/user"
	"github.com/wizact/go-todo-api/internal/user/adapters/repository"
	"github.com/wizact/go-todo-api/internal/user/domain"
	"github.com/wizact/go-todo-api/pkg/eventlibrary/pubsub"
	"github.com/wizact/go-todo-api/pkg/eventlibrary/user/event"
)

// newUserModule selects production adapters at the application composition root.
func newUserModule(config Config) (*usermodule.Module, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("validate application configuration: %w", err)
	}
	repositoryFactory := dbinfra.SQLiteRepositoryFactory[repository.SQLiteUserRepository, *repository.SQLiteUserRepository]{}
	repo, err := repositoryFactory.Create()
	if err != nil {
		return nil, fmt.Errorf("create user repository: %w", err)
	}
	publisherFactory := pubsub.NATSClientFactory[event.UserEventClient, domain.UserRegisteredEvent, *event.UserEventClient]{}
	publisher, err := publisherFactory.Create()
	if err != nil {
		return nil, fmt.Errorf("create user event publisher: %w", err)
	}
	return usermodule.NewModuleWithDependencies(usermodule.Dependencies{Repository: repo, EventPublisher: publisher}, usermodule.WithPublicBaseURL(config.PublicBaseURL))
}
