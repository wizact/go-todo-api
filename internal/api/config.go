package api

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/kelseyhightower/envconfig"
	userservice "github.com/wizact/go-todo-api/internal/user/application/service"
	"github.com/wizact/go-todo-api/pkg/communication"
	"github.com/wizact/go-todo-api/pkg/version"
)

// Config contains deployment settings required to assemble the application.
type Config struct {
	PublicBaseURL                  string `required:"true"`
	SendGridEnabled                bool   `default:"false"`
	SendGridKey                    string
	SendGridFromName               string
	SendGridFromEmail              string
	SendGridVerificationTemplateID string
}

// LoadConfig reads application configuration from the environment.
func LoadConfig() (Config, error) {
	config := Config{}
	if err := envconfig.Process(version.AppName, &config); err != nil {
		return Config{}, fmt.Errorf("load application configuration: %w", err)
	}
	config.PublicBaseURL = strings.TrimSpace(config.PublicBaseURL)
	if err := config.Validate(); err != nil {
		return Config{}, fmt.Errorf("validate application configuration: %w", err)
	}

	return config, nil
}

// Validate checks runtime configuration before constructing application modules.
func (c Config) Validate() error {
	publicBaseURL, err := url.Parse(strings.TrimSpace(c.PublicBaseURL))
	if err != nil {
		return fmt.Errorf("parse public base URL: %w", err)
	}
	if publicBaseURL.Host == "" || (publicBaseURL.Scheme != "http" && publicBaseURL.Scheme != "https") {
		return errors.New("public base URL must be an absolute HTTP or HTTPS URL")
	}
	if publicBaseURL.RawQuery != "" || publicBaseURL.ForceQuery || strings.Contains(c.PublicBaseURL, "#") {
		return errors.New("public base URL must not contain a query or fragment")
	}
	if err := c.communicationConfig().Validate(); err != nil {
		return fmt.Errorf("validate communication configuration: %w", err)
	}

	return nil
}

func (c Config) registrationConfig() userservice.RegistrationConfig {
	return userservice.RegistrationConfig{PublicBaseURL: c.PublicBaseURL}
}

func (c Config) communicationConfig() communication.Config {
	return communication.Config{
		SendGridEnabled:        c.SendGridEnabled,
		SendGridKey:            c.SendGridKey,
		SendGridFromName:       c.SendGridFromName,
		SendGridFromEmail:      c.SendGridFromEmail,
		VerificationTemplateID: c.SendGridVerificationTemplateID,
	}
}
