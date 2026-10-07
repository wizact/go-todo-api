package communication

import "testing"

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{name: "memory delivery needs no provider settings"},
		{
			name: "complete SendGrid configuration is valid",
			config: Config{
				SendGridEnabled:        true,
				SendGridKey:            "key",
				SendGridFromName:       "TODO API",
				SendGridFromEmail:      "todo@example.com",
				VerificationTemplateID: "verification-template",
			},
		},
		{
			name:    "missing SendGrid key is invalid",
			config:  Config{SendGridEnabled: true, SendGridFromName: "TODO API", SendGridFromEmail: "todo@example.com", VerificationTemplateID: "verification-template"},
			wantErr: true,
		},
		{
			name:    "missing sender name is invalid",
			config:  Config{SendGridEnabled: true, SendGridKey: "key", SendGridFromEmail: "todo@example.com", VerificationTemplateID: "verification-template"},
			wantErr: true,
		},
		{
			name:    "missing sender email is invalid",
			config:  Config{SendGridEnabled: true, SendGridKey: "key", SendGridFromName: "TODO API", VerificationTemplateID: "verification-template"},
			wantErr: true,
		},
		{
			name:    "missing template ID is invalid",
			config:  Config{SendGridEnabled: true, SendGridKey: "key", SendGridFromName: "TODO API", SendGridFromEmail: "todo@example.com"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			err := test.config.Validate()
			if (err != nil) != test.wantErr {
				t.Fatalf("Validate() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}
