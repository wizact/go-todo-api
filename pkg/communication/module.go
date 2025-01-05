package communication

import (
	umf "github.com/wizact/go-todo-api/internal/domain-factory"
	user_domain_listener "github.com/wizact/go-todo-api/pkg/communication/application/listeners/user"
	app_svc "github.com/wizact/go-todo-api/pkg/communication/application/services"
	ports "github.com/wizact/go-todo-api/pkg/communication/ports/applications"
	user_event_port "github.com/wizact/go-todo-api/pkg/event-library/ports/events"
	pubsubinfra "github.com/wizact/go-todo-api/pkg/event-library/pubsub"
	UserDomainEvent "github.com/wizact/go-todo-api/pkg/event-library/user/domain"
	user_event "github.com/wizact/go-todo-api/pkg/event-library/user/events"

	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/wizact/go-todo-api/pkg/version"
)

// A CommsModule is the dependency container for the communication module
// and if the use* flags are set to true, then it returns the concrete
// implementation of the interface instead of the memory or fake implementation.
type CommsModule struct {
	emailClientAppSvc ports.Emailer
	userEventClient   user_event_port.UserEventClient

	// Listeners
	newUserRegisteredListener *user_domain_listener.NewUserRegisteredEventListener
}

// New CommsModule is the factory method for the comms container
func NewCommsModule(useSendGrid bool) *CommsModule {
	userEventCli := instantiateUserEventClient()
	emailClientAppSvc := instantiateAppSvc(useSendGrid)

	udl := instantiateUserDomainListenersAndListen(userEventCli, emailClientAppSvc)

	return &CommsModule{
		userEventClient:           userEventCli,
		emailClientAppSvc:         emailClientAppSvc,
		newUserRegisteredListener: udl,
	}
}

func instantiateUserEventClient() user_event_port.UserEventClient {
	nf := pubsubinfra.NatsClientFactory[user_event.UserEventClient, UserDomainEvent.UserDomainEvent, *user_event.UserEventClient]{}
	uec, err := nf.Get()
	if err != nil {
		panic(err)
	}

	return uec
}

func instantiateAppSvc(useSendGrid bool) ports.Emailer {
	if !useSendGrid {
		return app_svc.NewMemoryEmailClient()
	}

	sg := &SendGridConfig{}
	err := sg.LoadConfig()
	if err != nil {
		panic(err)
	}

	return app_svc.NewSendGridEmailClient(sg.SendGridKey, sg.SendGridFromName, sg.SendGridFromEmail)
}

func instantiateUserDomainListenersAndListen(uec user_event_port.UserEventClient, ecas ports.Emailer) *user_domain_listener.NewUserRegisteredEventListener {

	um := umf.CreateNewUserModule()
	nurel := user_domain_listener.NewNewUserRegisteredEventListener(uec, um.UserRegistrationAppService(), ecas)
	err := nurel.Listen()
	if err != nil {
		panic(err)
	}
	return nurel
}

// SendGridConfig holds the configuration for sendgrid
type SendGridConfig struct {
	SendGridKey       string
	SendGridFromName  string
	SendGridFromEmail string
}

// LoadConfig gets the configuration from env variables for sendgrid
func (s *SendGridConfig) LoadConfig() error {
	err := envconfig.Process(version.APPNAME, s)
	if err != nil {
		panic(err)
	}

	if s.SendGridKey == "" && s.SendGridFromEmail == "" && s.SendGridFromName == "" {
		return errors.New("cannot resolve sendgrid configuration")
	}

	return nil
}

// Done cleans up all the underlying resources for a graceful shotdown
func (cm *CommsModule) Done() {
	cm.newUserRegisteredListener.Done()
}
