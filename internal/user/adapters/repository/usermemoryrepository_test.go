package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/application/registration"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/output/repository"
)

func TestMemoryUserRepository_FindRegistrationVerification_Missing_ReturnsNotFound(t *testing.T) {
	t.Parallel()

	repository := NewMemoryUserRepository(nil)

	_, err := repository.FindRegistrationVerification(context.Background(), uuid.New())

	if !errors.Is(err, repositoryport.ErrRegistrationVerificationNotFound) {
		t.Fatalf("FindRegistrationVerification() error = %v, want %v", err, repositoryport.ErrRegistrationVerificationNotFound)
	}
}

func TestMemoryUserRepository_SaveRegistrationVerification_ReplacesCredential(t *testing.T) {
	t.Parallel()

	repository := NewMemoryUserRepository(nil)
	userID := uuid.New()
	original := registration.Verification{
		UserID:       userID,
		SecretDigest: "original-digest",
		ExpiresAt:    time.Date(2026, time.September, 29, 12, 0, 0, 0, time.UTC),
	}
	replacement := registration.Verification{
		UserID:       userID,
		SecretDigest: "replacement-digest",
		ExpiresAt:    time.Date(2026, time.September, 30, 12, 0, 0, 0, time.UTC),
	}

	if err := repository.SaveRegistrationVerification(context.Background(), original); err != nil {
		t.Fatalf("save original registration verification: %v", err)
	}
	if err := repository.SaveRegistrationVerification(context.Background(), replacement); err != nil {
		t.Fatalf("save replacement registration verification: %v", err)
	}
	got, err := repository.FindRegistrationVerification(context.Background(), userID)
	if err != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", err)
	}

	if got != replacement {
		t.Fatalf("FindRegistrationVerification() = %#v, want %#v", got, replacement)
	}
}

func TestMemoryUserRepository_CompleteRegistration_UpdatesStateAndConsumesCredential(t *testing.T) {
	t.Parallel()

	user := newUserAggregate(t)
	repository := NewMemoryUserRepository([]aggregate.User{user})
	verification := registration.Verification{
		UserID:       user.ID(),
		SecretDigest: "verified-digest",
		ExpiresAt:    time.Date(2026, time.September, 29, 12, 0, 0, 0, time.UTC),
	}
	if err := repository.SaveRegistrationVerification(context.Background(), verification); err != nil {
		t.Fatalf("SaveRegistrationVerification() error = %v", err)
	}
	user.VerifyRegistration()

	if _, err := repository.CompleteRegistration(context.Background(), user, verification.SecretDigest); err != nil {
		t.Fatalf("CompleteRegistration() error = %v", err)
	}
	_, replayError := repository.CompleteRegistration(context.Background(), user, verification.SecretDigest)
	persistedUser, err := repository.FindByID(context.Background(), user.ID())
	if err != nil {
		t.Fatalf("FindById() error = %v", err)
	}
	_, credentialError := repository.FindRegistrationVerification(context.Background(), user.ID())

	got := registrationCompletionState{
		Active:               persistedUser.IsActive(),
		AggregateEmail:       persistedUser.HasVerifiedEmail(),
		ProjectionEmail:      persistedUser.HasVerifiedEmail(),
		CredentialWasRemoved: errors.Is(credentialError, repositoryport.ErrRegistrationVerificationNotFound),
		ReplayWasRejected:    errors.Is(replayError, repositoryport.ErrRegistrationVerificationNotFound),
	}
	want := registrationCompletionState{true, true, true, true, true}
	if got != want {
		t.Fatalf("registration completion state = %#v, want %#v", got, want)
	}
}

func TestMemoryUserRepository_CompleteRegistration_StaleDigestPreservesState(t *testing.T) {
	t.Parallel()

	user := newUserAggregate(t)
	repository := NewMemoryUserRepository([]aggregate.User{user})
	verification := registration.Verification{
		UserID:       user.ID(),
		SecretDigest: "current-digest",
		ExpiresAt:    time.Date(2026, time.September, 29, 12, 0, 0, 0, time.UTC),
	}
	if err := repository.SaveRegistrationVerification(context.Background(), verification); err != nil {
		t.Fatalf("SaveRegistrationVerification() error = %v", err)
	}
	user.VerifyRegistration()

	_, completionError := repository.CompleteRegistration(context.Background(), user, "stale-digest")
	persistedUser, err := repository.FindByID(context.Background(), user.ID())
	if err != nil {
		t.Fatalf("FindById() error = %v", err)
	}
	persistedVerification, err := repository.FindRegistrationVerification(context.Background(), user.ID())
	if err != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", err)
	}

	got := registrationRollbackState{
		ErrorWasNotFound:      errors.Is(completionError, repositoryport.ErrRegistrationVerificationNotFound),
		Active:                persistedUser.IsActive(),
		AggregateEmail:        persistedUser.HasVerifiedEmail(),
		ProjectionEmail:       persistedUser.HasVerifiedEmail(),
		CredentialWasRetained: persistedVerification == verification,
	}
	want := registrationRollbackState{ErrorWasNotFound: true, CredentialWasRetained: true}
	if got != want {
		t.Fatalf("registration rollback state = %#v, want %#v", got, want)
	}
}
