package service_test

import (
	"context"
	"testing"

	repository "github.com/wizact/go-todo-api/internal/user/adapters/repositories"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	svc "github.com/wizact/go-todo-api/internal/user/domain/services"
	event "github.com/wizact/go-todo-api/pkg/event-library/user/events"
)

func Test_NewUserAccountService(t *testing.T) {
	seedUserList := init_users(t)
	u := seedUserList[0]

	ur := repository.NewUserMemoryRepository(seedUserList)
	uecm := event.UserEventClientMock{}

	uas := svc.NewUserAccountService(ur, uecm)

	_, err := uas.RegisterNewUser(context.Background(), u)

	if err == nil {
		t.Error(err)
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
