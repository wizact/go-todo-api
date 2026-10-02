package registration

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrVerificationExpired        = errors.New("registration verification expired")
	ErrVerificationSecretMismatch = errors.New("registration verification secret does not match")
)

// Verification is the application-managed credential used to verify a user registration.
type Verification struct {
	UserID       uuid.UUID
	SecretDigest string
	ExpiresAt    time.Time
}
