package service_test

import (
	"context"
	"errors"
	"testing"

	repository "github.com/wizact/go-todo-api/internal/user/adapters/repositories"
	"github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	svc "github.com/wizact/go-todo-api/internal/user/domain/services"
	event "github.com/wizact/go-todo-api/pkg/event-library/user/events"
)

func TestUserAccountService_RegisterNewUser_DuplicateEmail_ReturnsDomainError(t *testing.T) {
	seedUserList := init_users(t)
	u := seedUserList[0]

	ur := repository.NewUserMemoryRepository(seedUserList)
	uecm := event.UserEventClientMock{}

	uas := svc.NewUserAccountService(ur, uecm)

	_, err := uas.RegisterNewUser(context.Background(), u)

	if !errors.Is(err, domain.ErrEmailAlreadyExists) {
		t.Errorf("error = %v, want %v", err, domain.ErrEmailAlreadyExists)
	}
}

func init_users(t *testing.T) []aggregate.User {
	ua := aggregate.NewUser()
	u := ua.User()
	u.SetName("John", "Doe")
	u.SetEmail("john.doe@example.com")

	ua.SetUser(u)

	var seedUserList []aggregate.User
	seedUserList = append(seedUserList, ua)
	return seedUserList
}
