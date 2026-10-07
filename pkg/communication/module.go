package communication

import (
	"context"
	"fmt"

	applicationport "github.com/wizact/go-todo-api/internal/user/ports/applications"
	user_domain_listener "github.com/wizact/go-todo-api/pkg/communication/application/listeners/user"
	app_svc "github.com/wizact/go-todo-api/pkg/communication/application/services"
	ports "github.com/wizact/go-todo-api/pkg/communication/ports/applications"
	user_event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/input/events"
	pubsubinfra "github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	UserDomainEvent "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
	user_event "github.com/wizact/go-todo-api/pkg/event-library/user/events"
)

// A CommsModule is the dependency container for the communication module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type CommsModule struct {
	emailClientAppSvc ports.Emailer
	userEventClient   user_event_port.UserEventClientInput

	// Listeners
	newUserRegisteredListener *user_domain_listener.NewUserRegisteredEventListener
}

// New CommsModule is the factory method for the comms container
func NewCommsModule(config Config, registration applicationport.Registration) (*CommsModule, error) {
	emailClientAppSvc, err := instantiateAppSvc(config)
	if err != nil {
		return nil, fmt.Errorf("configure communication module: %w", err)
	}
	userEventCli := instantiateUserEventClient()

	udl := instantiateUserDomainListenersAndListen(userEventCli, registration, emailClientAppSvc, config.VerificationTemplateID)

	return &CommsModule{
		userEventClient:           userEventCli,
		emailClientAppSvc:         emailClientAppSvc,
		newUserRegisteredListener: udl,
	}, nil
}

func instantiateUserEventClient() user_event_port.UserEventClientInput {
	nf := pubsubinfra.NatsClientFactory[user_event.UserEventClient, UserDomainEvent.UserDomainEvent, *user_event.UserEventClient]{}
	uec, err := nf.Get()
	if err != nil {
		panic(err)
	}

	return uec
}

func instantiateAppSvc(config Config) (ports.Emailer, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	if !config.SendGridEnabled {
		return app_svc.NewMemoryEmailClient(), nil
	}

	return app_svc.NewSendGridEmailClient(
		config.SendGridKey,
		config.SendGridFromName,
		config.SendGridFromEmail,
	), nil
}

func instantiateUserDomainListenersAndListen(
	uec user_event_port.UserEventClientInput,
	registration applicationport.Registration,
	ecas ports.Emailer,
	templateID string,
) *user_domain_listener.NewUserRegisteredEventListener {
	nurel := user_domain_listener.NewNewUserRegisteredEventListener(uec, registration, ecas, templateID)
	err := nurel.Listen(context.Background())
	if err != nil {
		panic(err)
	}
	return nurel
}

// Done cleans up all the underlying resources for a graceful shotdown
func (cm *CommsModule) Done() {
	cm.newUserRegisteredListener.Done()
}
