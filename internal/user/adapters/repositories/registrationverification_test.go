package repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/application/registration"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
)

type registrationCompletionState struct {
	Active               bool
	AggregateEmail       bool
	ProjectionEmail      bool
	CredentialWasRemoved bool
	ReplayWasRejected    bool
}

type registrationRollbackState struct {
	ErrorWasNotFound      bool
	Active                bool
	AggregateEmail        bool
	ProjectionEmail       bool
	CredentialWasRetained bool
}

func TestSqliteRegistrationVerification_TableName_ReturnsRegistrationVerificationTable(t *testing.T) {
	t.Parallel()

	got := (SqliteRegistrationVerification{}).TableName()
	want := "user_registration_verifications"
	if got != want {
		t.Fatalf("TableName() = %q, want %q", got, want)
	}
}

func TestUserSqliteRepository_FindRegistrationVerification_Missing_ReturnsNotFound(t *testing.T) {
	t.Parallel()

	repository, _ := newUserSqliteRepository(t, &SqliteRegistrationVerification{})

	_, err := repository.FindRegistrationVerification(context.Background(), uuid.New())

	if !errors.Is(err, repositoryport.ErrRegistrationVerificationNotFound) {
		t.Fatalf("FindRegistrationVerification() error = %v, want %v", err, repositoryport.ErrRegistrationVerificationNotFound)
	}
}

func TestUserSqliteRepository_SaveRegistrationVerification_PersistsCredential(t *testing.T) {
	t.Parallel()

	repository, _ := newUserSqliteRepository(t, &SqliteRegistrationVerification{})
	want := registration.Verification{
		UserID:       uuid.New(),
		SecretDigest: "verification-digest",
		ExpiresAt:    time.Date(2026, time.September, 29, 12, 0, 0, 0, time.UTC),
	}

	if err := repository.SaveRegistrationVerification(context.Background(), want); err != nil {
		t.Fatalf("SaveRegistrationVerification() error = %v", err)
	}
	got, err := repository.FindRegistrationVerification(context.Background(), want.UserID)
	if err != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", err)
	}

	if got != want {
		t.Fatalf("FindRegistrationVerification() = %#v, want %#v", got, want)
	}
}

func TestUserSqliteRepository_SaveRegistrationVerification_ReplacesCredential(t *testing.T) {
	t.Parallel()

	repository, _ := newUserSqliteRepository(t, &SqliteRegistrationVerification{})
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

func TestUserSqliteRepository_CompleteRegistration_UpdatesStateAndConsumesCredential(t *testing.T) {
	t.Parallel()

	repository, database := newUserSqliteRepository(
		t,
		&SqliteUserAggregate{},
		&SqliteUserEmailView{},
		&SqliteRegistrationVerification{},
	)
	user, err := repository.Create(context.Background(), newUserAggregate(t))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	verification := registration.Verification{
		UserID:       user.UserId(),
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
	persistedUser, err := repository.FindById(context.Background(), user.UserId())
	if err != nil {
		t.Fatalf("FindById() error = %v", err)
	}
	var emailView SqliteUserEmailView
	if err := database.First(&emailView, "user_id = ?", user.UserId().String()).Error; err != nil {
		t.Fatalf("find user email view: %v", err)
	}
	_, credentialError := repository.FindRegistrationVerification(context.Background(), user.UserId())

	got := registrationCompletionState{
		Active:               persistedUser.IsActive(),
		AggregateEmail:       persistedUser.HasVerifiedEmail(),
		ProjectionEmail:      emailView.HasVerifiedEmail,
		CredentialWasRemoved: errors.Is(credentialError, repositoryport.ErrRegistrationVerificationNotFound),
		ReplayWasRejected:    errors.Is(replayError, repositoryport.ErrRegistrationVerificationNotFound),
	}
	want := registrationCompletionState{true, true, true, true, true}
	if got != want {
		t.Fatalf("registration completion state = %#v, want %#v", got, want)
	}
}

func TestUserSqliteRepository_CompleteRegistration_StaleDigestRollsBack(t *testing.T) {
	t.Parallel()

	repository, database := newUserSqliteRepository(
		t,
		&SqliteUserAggregate{},
		&SqliteUserEmailView{},
		&SqliteRegistrationVerification{},
	)
	user, err := repository.Create(context.Background(), newUserAggregate(t))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	verification := registration.Verification{
		UserID:       user.UserId(),
		SecretDigest: "current-digest",
		ExpiresAt:    time.Date(2026, time.September, 29, 12, 0, 0, 0, time.UTC),
	}
	if err := repository.SaveRegistrationVerification(context.Background(), verification); err != nil {
		t.Fatalf("SaveRegistrationVerification() error = %v", err)
	}
	user.VerifyRegistration()

	_, completionError := repository.CompleteRegistration(context.Background(), user, "stale-digest")
	persistedUser, err := repository.FindById(context.Background(), user.UserId())
	if err != nil {
		t.Fatalf("FindById() error = %v", err)
	}
	var emailView SqliteUserEmailView
	if err := database.First(&emailView, "user_id = ?", user.UserId().String()).Error; err != nil {
		t.Fatalf("find user email view: %v", err)
	}
	persistedVerification, err := repository.FindRegistrationVerification(context.Background(), user.UserId())
	if err != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", err)
	}

	got := registrationRollbackState{
		ErrorWasNotFound:      errors.Is(completionError, repositoryport.ErrRegistrationVerificationNotFound),
		Active:                persistedUser.IsActive(),
		AggregateEmail:        persistedUser.HasVerifiedEmail(),
		ProjectionEmail:       emailView.HasVerifiedEmail,
		CredentialWasRetained: persistedVerification == verification,
	}
	want := registrationRollbackState{ErrorWasNotFound: true, CredentialWasRetained: true}
	if got != want {
		t.Fatalf("registration rollback state = %#v, want %#v", got, want)
	}
}

func TestUserSqliteRepository_CompleteRegistration_StaleDigestWithoutUserReturnsNotFound(t *testing.T) {
	t.Parallel()

	repository, _ := newUserSqliteRepository(
		t,
		&SqliteUserAggregate{},
		&SqliteUserEmailView{},
		&SqliteRegistrationVerification{},
	)

	_, err := repository.CompleteRegistration(
		context.Background(),
		newUserAggregate(t),
		"stale-digest",
	)

	if !errors.Is(err, repositoryport.ErrRegistrationVerificationNotFound) {
		t.Fatalf("CompleteRegistration() error = %v, want %v", err, repositoryport.ErrRegistrationVerificationNotFound)
	}
}
