package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	model "github.com/wizact/go-todo-api/internal/user/domain/model"
)

type sqliteUserUpdateState struct {
	ReturnedEmail        string
	AggregateEmail       string
	ProjectionEmail      string
	PreviousEmailMissing bool
}

func TestUserSQLiteRepository_Update_PersistsAggregateAndEmailView(t *testing.T) {
	t.Parallel()

	repository, database := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})
	original, err := repository.Create(context.Background(), newUserAggregate(t))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	updated := userAggregateWithEmail(original, "ada.updated@example.com")

	returned, err := repository.Update(context.Background(), updated)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	persisted, err := repository.FindByID(context.Background(), original.ID())
	if err != nil {
		t.Fatalf("FindById() error = %v", err)
	}
	var projection SQLiteUserEmailView
	if err := database.First(&projection, "user_id = ?", original.ID().String()).Error; err != nil {
		t.Fatalf("find user email view: %v", err)
	}
	_, previousEmailError := repository.FindByEmail(context.Background(), original.Email())

	got := sqliteUserUpdateState{
		ReturnedEmail:        returned.Email(),
		AggregateEmail:       persisted.Email(),
		ProjectionEmail:      projection.Email,
		PreviousEmailMissing: errors.Is(previousEmailError, domain.ErrUserEmailNotFound),
	}
	want := sqliteUserUpdateState{
		ReturnedEmail:        updated.Email(),
		AggregateEmail:       updated.Email(),
		ProjectionEmail:      updated.Email(),
		PreviousEmailMissing: true,
	}

	if got != want {
		t.Fatalf("Update() state = %#v, want %#v", got, want)
	}
}

func TestUserSQLiteRepository_Update_MissingUserReturnsNotFound(t *testing.T) {
	t.Parallel()

	repository, _ := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})

	_, err := repository.Update(context.Background(), newUserAggregate(t))

	if !errors.Is(err, domain.ErrUserIDNotFound) {
		t.Fatalf("Update() error = %v, want %v", err, domain.ErrUserIDNotFound)
	}
}

func TestUserSQLiteRepository_Update_EmailViewFailureRollsBackAggregate(t *testing.T) {
	t.Parallel()

	repository, database := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})
	original, err := repository.Create(context.Background(), newUserAggregate(t))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if err := database.Migrator().DropTable(&SQLiteUserEmailView{}); err != nil {
		t.Fatalf("DropTable() error = %v", err)
	}
	updated := userAggregateWithEmail(original, "ada.updated@example.com")

	_, updateError := repository.Update(context.Background(), updated)
	persisted, findError := repository.FindByID(context.Background(), original.ID())
	if findError != nil {
		t.Fatalf("FindById() error = %v", findError)
	}
	got := struct {
		UpdateFailed   bool
		PersistedEmail string
	}{UpdateFailed: updateError != nil, PersistedEmail: persisted.Email()}
	want := struct {
		UpdateFailed   bool
		PersistedEmail string
	}{UpdateFailed: true, PersistedEmail: original.Email()}

	if got != want {
		t.Fatalf("Update() rollback state = %#v, want %#v", got, want)
	}
}

func TestUserSQLiteRepository_Update_CanceledContextReturnsError(t *testing.T) {
	t.Parallel()

	repository, _ := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})
	user, err := repository.Create(context.Background(), newUserAggregate(t))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = repository.Update(ctx, user)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Update() error = %v, want %v", err, context.Canceled)
	}
}

func userAggregateWithEmail(user aggregate.User, email string) aggregate.User {
	entity := user.User()
	firstName, lastName := entity.Name()
	updated := model.RehydrateUser(
		user.ID(),
		firstName,
		lastName,
		entity.DateOfBirth(),
		email,
		entity.Phone(),
	)
	return aggregate.RehydrateUser(
		updated,
		user.Location(),
		aggregate.RegistrationStatus{
			IsActive:         user.IsActive(),
			HasVerifiedEmail: user.HasVerifiedEmail(),
		},
	)
}
