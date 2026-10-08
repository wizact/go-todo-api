package communication

import (
	"context"
	"fmt"

	applicationport "github.com/wizact/go-todo-api/internal/user/ports/service"
	userDomainListener "github.com/wizact/go-todo-api/pkg/communication/application/listeners/user"
	appService "github.com/wizact/go-todo-api/pkg/communication/application/service"
	ports "github.com/wizact/go-todo-api/pkg/communication/ports/service"
	userEventPort "github.com/wizact/go-todo-api/pkg/eventlibrary/ports/input/event"
	pubsubinfra "github.com/wizact/go-todo-api/pkg/eventlibrary/pubsub"
	userDomainEvent "github.com/wizact/go-todo-api/pkg/eventlibrary/user/domain"
	userEvent "github.com/wizact/go-todo-api/pkg/eventlibrary/user/event"
)

// A Module is the dependency container for the communication module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type Module struct {
	emailClient     ports.Emailer
	userEventClient userEventPort.UserSubscriber

	// Listeners
	newUserRegisteredListener *userDomainListener.RegisteredEventListener
}

// New Module is the factory method for the comms container
func NewModule(config Config, registration applicationport.Registration) (*Module, error) {
	emailClient, err := instantiateApplicationService(config)
	if err != nil {
		return nil, fmt.Errorf("configure communication module: %w", err)
	}
	userEventCli := instantiateUserEventClient()

	udl := instantiateUserDomainListenersAndListen(userEventCli, registration, emailClient, config.VerificationTemplateID)

	return &Module{
		userEventClient:           userEventCli,
		emailClient:               emailClient,
		newUserRegisteredListener: udl,
	}, nil
}

func instantiateUserEventClient() userEventPort.UserSubscriber {
	nf := pubsubinfra.NATSClientFactory[userEvent.UserEventClient, userDomainEvent.UserDomainEvent, *userEvent.UserEventClient]{}
	uec, err := nf.Create()
	if err != nil {
		panic(err)
	}

	return uec
}

func instantiateApplicationService(config Config) (ports.Emailer, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	if !config.SendGridEnabled {
		return appService.NewMemoryEmailClient(), nil
	}

	return appService.NewSendGridEmailClient(
		config.SendGridKey,
		config.SendGridFromName,
		config.SendGridFromEmail,
	), nil
}

func instantiateUserDomainListenersAndListen(
	uec userEventPort.UserSubscriber,
	registration applicationport.Registration,
	ecas ports.Emailer,
	templateID string,
) *userDomainListener.RegisteredEventListener {
	nurel := userDomainListener.NewRegisteredEventListener(uec, registration, ecas, templateID)
	err := nurel.Listen(context.Background())
	if err != nil {
		panic(err)
	}
	return nurel
}

// Done cleans up all the underlying resources for a graceful shotdown
func (m *Module) Done() {
	m.newUserRegisteredListener.Done()
}
