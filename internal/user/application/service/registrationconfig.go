package service

const defaultPublicBaseURL = "http://localhost:8080"

// RegistrationConfig contains deployment-specific registration workflow values.
type RegistrationConfig struct {
	PublicBaseURL string
}

// DefaultRegistrationConfig returns the local development configuration.
func DefaultRegistrationConfig() RegistrationConfig {
	return RegistrationConfig{PublicBaseURL: defaultPublicBaseURL}
}
