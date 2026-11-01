package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/domain"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/repository"
)

func TestUserRepository_MissingUser_ReturnsDomainError(t *testing.T) {
	t.Parallel()

	factories := []struct {
		name          string
		newRepository func(*testing.T) repositoryport.UserRepository
	}{
		{"memory", func(t *testing.T) repositoryport.UserRepository {
			return NewMemoryUserRepository(nil)
		}},
		{"sqlite", func(t *testing.T) repositoryport.UserRepository {
			repository, _ := newSQLiteUserRepository(t, &SQLiteUserAggregate{}, &SQLiteUserEmailView{})
			return repository
		}},
	}
	operations := []struct {
		name   string
		invoke func(*testing.T, repositoryport.UserRepository) error
		want   error
	}{
		{"find by ID", func(t *testing.T, repository repositoryport.UserRepository) error {
			_, err := repository.FindByID(context.Background(), uuid.New())
			return err
		}, domain.ErrUserIDNotFound},
		{"find by email", func(t *testing.T, repository repositoryport.UserRepository) error {
			_, err := repository.FindByEmail(context.Background(), "missing@example.com")
			return err
		}, domain.ErrUserEmailNotFound},
		{"update", func(t *testing.T, repository repositoryport.UserRepository) error {
			_, err := repository.Update(context.Background(), newUserAggregate(t))
			return err
		}, domain.ErrUserIDNotFound},
	}
	for _, factory := range factories {
		t.Run(factory.name, func(t *testing.T) {
			t.Parallel()
			for _, operation := range operations {
				t.Run(operation.name, func(t *testing.T) {
					err := operation.invoke(t, factory.newRepository(t))
					if !errors.Is(err, operation.want) {
						t.Fatalf("error = %v, want errors.Is with %v", err, operation.want)
					}
				})
			}
		})
	}
}

func TestMemoryUserRepository_MissingUser_PreservesAdapterError(t *testing.T) {
	t.Parallel()

	repository := NewMemoryUserRepository(nil)
	tests := []struct {
		name   string
		invoke func() error
		want   error
	}{
		{"find by ID", func() error {
			_, err := repository.FindByID(context.Background(), uuid.New())
			return err
		}, ErrUserNotFound},
		{"find by email", func() error {
			_, err := repository.FindByEmail(context.Background(), "missing@example.com")
			return err
		}, ErrUserNotFound},
		{"update", func() error {
			_, err := repository.Update(context.Background(), newUserAggregate(t))
			return err
		}, ErrFailedToUpdateUser},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.invoke(); !errors.Is(err, test.want) {
				t.Fatalf("error = %v, want errors.Is with %v", err, test.want)
			}
		})
	}
}
