package service

import "testing"

func TestRegistrationConfigValidate(t *testing.T) {
	for _, value := range []string{"http://localhost:8080", " https://example.com/app "} {
		if err := (RegistrationConfig{PublicBaseURL: value}).Validate(); err != nil {
			t.Errorf("Validate(%q): %v", value, err)
		}
	}
	for _, value := range []string{"", "/relative", "ftp://example.com", "https://example.com?", "https://example.com?q=1", "https://example.com#", "http://%"} {
		if err := (RegistrationConfig{PublicBaseURL: value}).Validate(); err == nil {
			t.Errorf("Validate(%q) accepted invalid URL", value)
		}
	}
}
