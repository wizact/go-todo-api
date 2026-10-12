package user

import (
	"errors"
	"fmt"
	"strings"

	appService "github.com/wizact/go-todo-api/internal/user/application/service"
	usecase "github.com/wizact/go-todo-api/internal/user/domain/service"
	usecasePort "github.com/wizact/go-todo-api/internal/user/ports/input/usecase"
	eventPort "github.com/wizact/go-todo-api/internal/user/ports/output/event"
	repositoryPort "github.com/wizact/go-todo-api/internal/user/ports/output/repository"
	appServicePort "github.com/wizact/go-todo-api/internal/user/ports/service"
)

// Module exposes the assembled user services.
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

// Dependencies contains required collaborators. One repository owns user and
// registration persistence so completing registration remains atomic.
type Dependencies struct {
	Repository     userRepository
	EventPublisher eventPort.UserEventPublisher
}

type moduleConfig struct{ registration appService.RegistrationConfig }

// Option configures optional module settings. Options apply in order.
type Option func(*moduleConfig)

// WithPublicBaseURL sets the URL used in registration verification emails.
func WithPublicBaseURL(value string) Option {
	return func(config *moduleConfig) { config.registration.PublicBaseURL = strings.TrimSpace(value) }
}

// NewModuleWithDependencies validates configuration and assembles services without
// reading environment variables, opening connections, or starting workers.
func NewModuleWithDependencies(deps Dependencies, options ...Option) (*Module, error) {
	config := moduleConfig{registration: appService.DefaultRegistrationConfig()}
	for i, option := range options {
		if option == nil {
			return nil, fmt.Errorf("nil option at index %d", i)
		}
		option(&config)
	}
	if deps.Repository == nil {
		return nil, errors.New("user repository is required")
	}
	if deps.EventPublisher == nil {
		return nil, errors.New("user event publisher is required")
	}
	if err := config.registration.Validate(); err != nil {
		return nil, fmt.Errorf("validate registration configuration: %w", err)
	}
	account := usecase.NewUserAccount(deps.Repository, deps.EventPublisher)
	return &Module{
		userRepository:      deps.Repository,
		userEventPublisher:  deps.EventPublisher,
		userAccountUseCase:  account,
		registrationService: appService.NewRegistration(account, deps.Repository, config.registration),
	}, nil
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
