package service

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const defaultPublicBaseURL = "http://localhost:8080"

// RegistrationConfig contains deployment-specific registration workflow values.
type RegistrationConfig struct {
	PublicBaseURL string
}

// DefaultRegistrationConfig returns the local development configuration.
func DefaultRegistrationConfig() RegistrationConfig {
	return RegistrationConfig{PublicBaseURL: defaultPublicBaseURL}
}

// Validate checks the public URL before assembling registration services.
func (c RegistrationConfig) Validate() error {
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
	return nil
}
