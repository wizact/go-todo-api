package api

import (
	"testing"

	userservice "github.com/wizact/go-todo-api/internal/user/application/service"
	"github.com/wizact/go-todo-api/pkg/communication"
)

func TestLoadConfig_MapsVerificationDeliverySettings(t *testing.T) {
	t.Setenv("TODOAPI_PUBLICBASEURL", "https://todo.example.com/")
	t.Setenv("TODOAPI_SENDGRIDENABLED", "true")
	t.Setenv("TODOAPI_SENDGRIDKEY", "sendgrid-key")
	t.Setenv("TODOAPI_SENDGRIDFROMNAME", "TODO API")
	t.Setenv("TODOAPI_SENDGRIDFROMEMAIL", "todo@example.com")
	t.Setenv("TODOAPI_SENDGRIDVERIFICATIONTEMPLATEID", "verification-template")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}

	got := struct {
		registration  userservice.RegistrationConfig
		communication communication.Config
	}{
		registration:  config.registrationConfig(),
		communication: config.communicationConfig(),
	}
	want := struct {
		registration  userservice.RegistrationConfig
		communication communication.Config
	}{
		registration: userservice.RegistrationConfig{PublicBaseURL: "https://todo.example.com/"},
		communication: communication.Config{
			SendGridEnabled:        true,
			SendGridKey:            "sendgrid-key",
			SendGridFromName:       "TODO API",
			SendGridFromEmail:      "todo@example.com",
			VerificationTemplateID: "verification-template",
		},
	}
	if got != want {
		t.Fatalf("module configs = %#v, want %#v", got, want)
	}
}

func TestLoadConfig_PublicBaseURLWithWhitespace_Normalizes(t *testing.T) {
	t.Setenv("TODOAPI_PUBLICBASEURL", " https://todo.example.com/app/ ")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() error = %v", err)
	}
	if got, want := config.registrationConfig().PublicBaseURL, "https://todo.example.com/app/"; got != want {
		t.Fatalf("PublicBaseURL = %q, want %q", got, want)
	}
}

func TestLoadConfig_InvalidDeliverySettingsReturnsError(t *testing.T) {
	tests := []struct {
		name            string
		publicBaseURL   string
		sendGridEnabled string
		sendGridKey     string
	}{
		{name: "missing public base URL", sendGridEnabled: "false"},
		{name: "non-HTTP public base URL", publicBaseURL: "ftp://todo.example.com", sendGridEnabled: "false"},
		{name: "public base URL with query", publicBaseURL: "https://todo.example.com/app?tenant=x", sendGridEnabled: "false"},
		{name: "public base URL with fragment", publicBaseURL: "https://todo.example.com/app#welcome", sendGridEnabled: "false"},
		{name: "enabled SendGrid without key", publicBaseURL: "https://todo.example.com", sendGridEnabled: "true"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("TODOAPI_PUBLICBASEURL", test.publicBaseURL)
			t.Setenv("TODOAPI_SENDGRIDENABLED", test.sendGridEnabled)
			t.Setenv("TODOAPI_SENDGRIDKEY", test.sendGridKey)
			t.Setenv("TODOAPI_SENDGRIDFROMNAME", "TODO API")
			t.Setenv("TODOAPI_SENDGRIDFROMEMAIL", "todo@example.com")
			t.Setenv("TODOAPI_SENDGRIDVERIFICATIONTEMPLATEID", "verification-template")

			_, err := LoadConfig()

			if err == nil {
				t.Fatal("LoadConfig() error = nil, want validation error")
			}
		})
	}
}
