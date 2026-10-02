package registration

import (
	"time"

	"github.com/google/uuid"
)

// Verification is the application-managed credential used to verify a user registration.
type Verification struct {
	UserID       uuid.UUID
	SecretDigest string
	ExpiresAt    time.Time
}
