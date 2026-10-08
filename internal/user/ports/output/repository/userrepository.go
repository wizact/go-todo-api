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
	FindByID(ctx context.Context, id uuid.UUID) (ua.User, error)
	FindByEmail(ctx context.Context, email string) (ua.User, error)
	Create(ctx context.Context, user ua.User) (ua.User, error)
	Update(ctx context.Context, user ua.User) (ua.User, error)
}

// RegistrationRepository provides persistence for the registration workflow.
// Implementations must atomically persist aggregate and projection changes when completing registration.
type RegistrationRepository interface {
	SaveRegistrationVerification(ctx context.Context, verification registration.Verification) error
	FindRegistrationVerification(ctx context.Context, userID uuid.UUID) (registration.Verification, error)
	CompleteRegistration(ctx context.Context, user ua.User, verificationDigest string) (ua.User, error)
}
