package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	repository "github.com/wizact/go-todo-api/internal/user/adapters/repositories"
	"github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
	svc "github.com/wizact/go-todo-api/internal/user/domain/services"
)

func TestUserAccountService_RegisterNewUser_DuplicateEmail_ReturnsDomainError(t *testing.T) {
	seedUserList := init_users(t)
	u := seedUserList[0]

	ur := repository.NewUserMemoryRepository(seedUserList)
	publisher := userEventPublisherStub{}

	uas := svc.NewUserAccountService(ur, publisher)

	_, err := uas.RegisterNewUser(context.Background(), u)

	if !errors.Is(err, domain.ErrEmailAlreadyExists) {
		t.Errorf("error = %v, want %v", err, domain.ErrEmailAlreadyExists)
	}
}

type userEventPublisherStub struct{}

func (userEventPublisherStub) PublishNewUserRegisteredEvent(context.Context, domain.UserRegisteredEvent) error {
	return nil
}

func init_users(t *testing.T) []aggregate.User {
	t.Helper()

	ua := aggregate.NewUser()
	u, err := model.NewUser("John", "Doe", time.Time{}, "john.doe@example.com", model.PhoneNumber{})
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	ua.SetUser(u)

	var seedUserList []aggregate.User
	seedUserList = append(seedUserList, ua)
	return seedUserList
}
