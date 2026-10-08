package service

import (
	"context"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	repositoryadapter "github.com/wizact/go-todo-api/internal/user/adapters/repository"
	"github.com/wizact/go-todo-api/internal/user/domain"
	domainservice "github.com/wizact/go-todo-api/internal/user/domain/service"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/output/repository"
	"gorm.io/gorm"
)

type registrationIntegrationState struct {
	TokenPresent           bool
	AggregateOmitsToken    bool
	CredentialOmitsToken   bool
	EventOmitsToken        bool
	DigestMatches          bool
	ExpiresAt              time.Time
	Active                 bool
	EmailVerified          bool
	CredentialWasConsumed  bool
	TokenReplayWasRejected bool
}

type recordingUserEventPublisher struct {
	event domain.UserRegisteredEvent
}

func (p *recordingUserEventPublisher) PublishNewUserRegisteredEvent(
	_ context.Context,
	event domain.UserRegisteredEvent,
) error {
	p.event = event
	return nil
}

func TestRegistration_SqliteJourneyPersistsDigestAndConsumesCredential(t *testing.T) {
	userID := uuid.New()
	user := registrationTestUser(userID)
	now := time.Date(2026, time.October, 4, 12, 0, 0, 0, time.UTC)
	repository, database := newRegistrationIntegrationRepository(t)
	publisher := &recordingUserEventPublisher{}
	userAccount := domainservice.NewUserAccount(repository, publisher)

	if _, err := userAccount.RegisterNewUser(context.Background(), user); err != nil {
		t.Fatalf("RegisterNewUser() error = %v", err)
	}
	registration := NewRegistration(userAccount, repository, DefaultRegistrationConfig())
	registration.now = func() time.Time { return now }
	emailData, err := registration.FetchRegistrationVerificationEmailData(context.Background(), userID)
	if err != nil {
		t.Fatalf("GetRegistrationVerificationEmailData() error = %v", err)
	}
	token := emailData["token"]
	verification, err := repository.FindRegistrationVerification(context.Background(), userID)
	if err != nil {
		t.Fatalf("FindRegistrationVerification() error = %v", err)
	}
	aggregateJSON := persistedAggregateJSON(t, database, userID)
	eventJSON, err := json.Marshal(publisher.event)
	if err != nil {
		t.Fatalf("marshal registered event: %v", err)
	}

	if err := registration.VerifyUserRegistration(context.Background(), userID, token); err != nil {
		t.Fatalf("VerifyUserRegistration() error = %v", err)
	}
	replayError := registration.VerifyUserRegistration(context.Background(), userID, token)
	persistedUser, err := repository.FindByID(context.Background(), userID)
	if err != nil {
		t.Fatalf("FindById() error = %v", err)
	}
	_, credentialError := repository.FindRegistrationVerification(context.Background(), userID)

	got := registrationIntegrationState{
		TokenPresent:           token != "",
		AggregateOmitsToken:    !strings.Contains(aggregateJSON, token),
		CredentialOmitsToken:   verification.SecretDigest != token && !strings.Contains(verification.SecretDigest, token),
		EventOmitsToken:        !strings.Contains(string(eventJSON), token),
		DigestMatches:          matchesRegistrationVerification(verification.SecretDigest, token),
		ExpiresAt:              verification.ExpiresAt,
		Active:                 persistedUser.IsActive(),
		EmailVerified:          persistedUser.HasVerifiedEmail(),
		CredentialWasConsumed:  errors.Is(credentialError, repositoryport.ErrRegistrationVerificationNotFound),
		TokenReplayWasRejected: errors.Is(replayError, repositoryport.ErrRegistrationVerificationNotFound),
	}
	want := registrationIntegrationState{
		TokenPresent:           true,
		AggregateOmitsToken:    true,
		CredentialOmitsToken:   true,
		EventOmitsToken:        true,
		DigestMatches:          true,
		ExpiresAt:              now.Add(registrationVerificationLifetime),
		Active:                 true,
		EmailVerified:          true,
		CredentialWasConsumed:  true,
		TokenReplayWasRejected: true,
	}
	if got != want {
		t.Fatalf("registration integration state = %#v, want %#v", got, want)
	}
}

func newRegistrationIntegrationRepository(t *testing.T) (*repositoryadapter.SQLiteUserRepository, *gorm.DB) {
	t.Helper()

	connection, err := dbinfra.NewSQLiteConnection(filepath.Join(t.TempDir(), "users.db"))
	if err != nil {
		t.Fatalf("NewSqliteConnection() error = %v", err)
	}
	database, err := connection.Open(gorm.Config{})
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}
	if err := database.AutoMigrate(
		&repositoryadapter.SQLiteUserAggregate{},
		&repositoryadapter.SQLiteUserEmailView{},
		&repositoryadapter.SQLiteRegistrationVerification{},
	); err != nil {
		t.Fatalf("AutoMigrate() error = %v", err)
	}
	repository := &repositoryadapter.SQLiteUserRepository{}
	repository.SetConnection(connection)
	return repository, database
}

func persistedAggregateJSON(t *testing.T, database *gorm.DB, userID uuid.UUID) string {
	t.Helper()

	var value string
	result := database.Raw(
		"SELECT value_data FROM users_aggregate WHERE user_id = ?",
		userID.String(),
	).Scan(&value)
	if result.Error != nil {
		t.Fatalf("find persisted aggregate JSON: %v", result.Error)
	}
	if result.RowsAffected != 1 {
		t.Fatalf("persisted aggregate rows = %d, want 1", result.RowsAffected)
	}
	return value
}
