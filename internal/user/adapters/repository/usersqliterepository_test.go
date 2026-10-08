package repository

import (
	"context"
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"
	"testing"
	"time"

	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	model "github.com/wizact/go-todo-api/internal/user/domain/model"
	"gorm.io/gorm"
)

func TestSQLiteUserModel_LegacyVerificationFieldsAreDropped(t *testing.T) {
	t.Parallel()

	legacy := []byte(`{"ID":"user-id","VerificationToken":"legacy-token","VerificationSalt":"legacy-salt"}`)
	var user SQLiteUserModel
	if err := json.Unmarshal(legacy, &user); err != nil {
		t.Fatalf("unmarshal legacy user: %v", err)
	}
	reencoded, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("marshal user: %v", err)
	}
	var fields map[string]any
	if err := json.Unmarshal(reencoded, &fields); err != nil {
		t.Fatalf("unmarshal reencoded user: %v", err)
	}
	_, tokenExists := fields["VerificationToken"]
	_, saltExists := fields["VerificationSalt"]

	if tokenExists || saltExists {
		t.Fatalf("reencoded fields = %v, want legacy verification fields omitted", fields)
	}
}

func TestSQLiteUserAggregate_toDomain_MalformedUserIDReturnsError(t *testing.T) {
	t.Parallel()

	_, err := (SQLiteUserAggregate{UserID: "not-a-uuid"}).toDomain()

	if err == nil {
		t.Fatal("ToDomain() error = nil, want malformed user ID error")
	}
}

func TestSQLiteUserEmailView_toDomain_MalformedUserIDReturnsError(t *testing.T) {
	t.Parallel()

	_, err := (SQLiteUserEmailView{UserID: "not-a-uuid"}).toDomain()

	if err == nil {
		t.Fatal("ToDomain() error = nil, want malformed user ID error")
	}
}

func TestSQLiteUserRepository_Create_DoesNotRequireTokenView(t *testing.T) {
	t.Parallel()

	repository, _ := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})
	user := newUserAggregate(t)

	_, err := repository.Create(context.Background(), user)

	if err != nil {
		t.Fatalf("Create() error = %v, want nil", err)
	}
}

func TestSQLiteUserRepository_Create_EmailViewFailureReturnsError(t *testing.T) {
	t.Parallel()

	repository, database := newSQLiteUserRepository(t, &SQLiteUserAggregate{})
	user := newUserAggregate(t)

	_, err := repository.Create(context.Background(), user)
	var aggregateCount int64
	countError := database.Model(&SQLiteUserAggregate{}).
		Where("user_id = ?", user.ID().String()).
		Count(&aggregateCount).
		Error

	if err == nil || !strings.Contains(err.Error(), "persist user email view") || countError != nil || aggregateCount != 0 {
		t.Fatalf("Create() error = %v, aggregate count = %d, count error = %v; want contextual error and rolled-back aggregate", err, aggregateCount, countError)
	}
}

func TestSQLiteUserRepository_FindByID_CanceledContextReturnsError(t *testing.T) {
	t.Parallel()

	repository, _ := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})
	user, err := repository.Create(context.Background(), newUserAggregate(t))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = repository.FindByID(ctx, user.ID())

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("FindById() error = %v, want %v", err, context.Canceled)
	}
}

func TestSQLiteUserRepository_FindByEmail_CanceledContextReturnsError(t *testing.T) {
	t.Parallel()

	repository, _ := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})
	user, err := repository.Create(context.Background(), newUserAggregate(t))
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err = repository.FindByEmail(ctx, user.Email())

	if !errors.Is(err, context.Canceled) {
		t.Fatalf("FindByEmail() error = %v, want %v", err, context.Canceled)
	}
}

func TestSQLiteUserRepository_Create_PersistsAggregateAndEmailView(t *testing.T) {
	t.Parallel()

	repository, database := newSQLiteUserRepository(
		t,
		&SQLiteUserAggregate{},
		&SQLiteUserEmailView{},
	)
	user := newUserAggregate(t)

	if _, err := repository.Create(context.Background(), user); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	models := []any{&SQLiteUserAggregate{}, &SQLiteUserEmailView{}}
	var got [2]int64
	for index, model := range models {
		result := database.Model(model).
			Where("user_id = ?", user.ID().String()).
			Count(&got[index])
		if result.Error != nil {
			t.Fatalf("count model %T: %v", model, result.Error)
		}
	}

	want := [2]int64{1, 1}
	if got != want {
		t.Fatalf("persisted row counts = %v, want %v", got, want)
	}
}

func newSQLiteUserRepository(t *testing.T, models ...any) (*SQLiteUserRepository, *gorm.DB) {
	t.Helper()

	connection, err := dbinfra.NewSQLiteConnection(filepath.Join(t.TempDir(), "users.db"))
	if err != nil {
		t.Fatalf("NewSqliteConnection() error = %v", err)
	}

	database, err := connection.Open(gorm.Config{})
	if err != nil {
		t.Fatalf("Open() error = %v", err)
	}

	if err := database.AutoMigrate(models...); err != nil {
		t.Fatalf("AutoMigrate() error = %v", err)
	}

	repository := &SQLiteUserRepository{}
	repository.SetConnection(connection)
	return repository, database
}

func newUserAggregate(t *testing.T) aggregate.User {
	t.Helper()

	user, err := model.NewUser(
		"Ada",
		"Lovelace",
		time.Date(1815, time.December, 10, 0, 0, 0, 0, time.UTC),
		"ada@example.com",
		model.NewPhoneNumber("+44", "20", "12345678"),
	)
	if err != nil {
		t.Fatalf("NewUser() error = %v", err)
	}

	root := aggregate.NewUser()
	root.SetUser(user)
	root.SetLocation(model.NewLocation(0, 0))
	return root
}
