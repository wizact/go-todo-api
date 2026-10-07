package communication

import (
	"fmt"
	"strings"
)

// Config contains delivery-provider settings for the communication module.
type Config struct {
	SendGridEnabled        bool
	SendGridKey            string
	SendGridFromName       string
	SendGridFromEmail      string
	VerificationTemplateID string
}

// Validate checks provider settings before the communication module starts.
func (c Config) Validate() error {
	if !c.SendGridEnabled {
		return nil
	}

	required := []struct {
		name  string
		value string
	}{
		{name: "SendGrid key", value: c.SendGridKey},
		{name: "SendGrid sender name", value: c.SendGridFromName},
		{name: "SendGrid sender email", value: c.SendGridFromEmail},
		{name: "verification template ID", value: c.VerificationTemplateID},
	}
	for _, setting := range required {
		if strings.TrimSpace(setting.value) == "" {
			return fmt.Errorf("%s is required when SendGrid is enabled", setting.name)
		}
	}

	return nil
}
