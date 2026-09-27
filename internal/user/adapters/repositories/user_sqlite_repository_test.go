package repository

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
	"gorm.io/gorm"
)

func TestUserSqliteRepository_Create_TokenViewFailureReturnsError(t *testing.T) {
	t.Parallel()

	repository, _ := newUserSqliteRepository(t, &SqliteUserAggregate{}, &SqliteUserEmailView{})
	user := newUserAggregate(t)

	_, err := repository.Create(context.Background(), user)

	if err == nil {
		t.Fatal("Create() error = nil, want token view persistence error")
	}
}

func TestUserSqliteRepository_Create_TokenViewFailureRollsBackAggregate(t *testing.T) {
	t.Parallel()

	repository, database := newUserSqliteRepository(t, &SqliteUserAggregate{}, &SqliteUserEmailView{})
	user := newUserAggregate(t)

	_, _ = repository.Create(context.Background(), user)

	var aggregateCount int64
	result := database.Unscoped().
		Model(&SqliteUserAggregate{}).
		Where("user_id = ?", user.UserId().String()).
		Count(&aggregateCount)
	if result.Error != nil {
		t.Fatalf("count user aggregates: %v", result.Error)
	}
	if aggregateCount != 0 {
		t.Fatalf("persisted aggregate count = %d, want 0", aggregateCount)
	}
}

func TestUserSqliteRepository_Create_EmailViewFailureReturnsError(t *testing.T) {
	t.Parallel()

	repository, _ := newUserSqliteRepository(t, &SqliteUserAggregate{})
	user := newUserAggregate(t)

	_, err := repository.Create(context.Background(), user)

	if err == nil {
		t.Fatal("Create() error = nil, want email view persistence error")
	}
}

func TestUserSqliteRepository_Create_PersistsAggregateAndViews(t *testing.T) {
	t.Parallel()

	repository, database := newUserSqliteRepository(
		t,
		&SqliteUserAggregate{},
		&SqliteUserEmailView{},
		&SqliteUserTokenView{},
	)
	user := newUserAggregate(t)

	if _, err := repository.Create(context.Background(), user); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	models := []any{&SqliteUserAggregate{}, &SqliteUserEmailView{}, &SqliteUserTokenView{}}
	var got [3]int64
	for index, model := range models {
		result := database.Model(model).
			Where("user_id = ?", user.UserId().String()).
			Count(&got[index])
		if result.Error != nil {
			t.Fatalf("count model %T: %v", model, result.Error)
		}
	}

	want := [3]int64{1, 1, 1}
	if got != want {
		t.Fatalf("persisted row counts = %v, want %v", got, want)
	}
}

func newUserSqliteRepository(t *testing.T, models ...any) (*UserSqliteRepository, *gorm.DB) {
	t.Helper()

	connection, err := dbinfra.NewSqliteConnection(filepath.Join(t.TempDir(), "users.db"))
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

	repository := &UserSqliteRepository{}
	repository.Connection(connection)
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
