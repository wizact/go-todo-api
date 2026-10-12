package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/application/registration"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
)

var ErrRegistrationVerificationNotFound = errors.New("registration verification not found")

// UserRepository provides persistence for the user aggregate.
type UserRepository interface {
	// FindByID returns an error matching domain.ErrUserIDNotFound when the user is missing.
	FindByID(ctx context.Context, id uuid.UUID) (ua.User, error)
	// FindByEmail returns an error matching domain.ErrUserEmailNotFound when the email is missing.
	FindByEmail(ctx context.Context, email string) (ua.User, error)
	Create(ctx context.Context, user ua.User) (ua.User, error)
	// Update returns an error matching domain.ErrUserIDNotFound when the user is missing.
	Update(ctx context.Context, user ua.User) (ua.User, error)
}

// RegistrationRepository provides persistence for the registration workflow.
// Implementations must atomically persist aggregate and projection changes when completing registration.
type RegistrationRepository interface {
	SaveRegistrationVerification(ctx context.Context, verification registration.Verification) error
	FindRegistrationVerification(ctx context.Context, userID uuid.UUID) (registration.Verification, error)
	CompleteRegistration(ctx context.Context, user ua.User, verificationDigest string) (ua.User, error)
}
