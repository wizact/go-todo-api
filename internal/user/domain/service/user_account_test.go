package service

import (
	"testing"

	"github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	"github.com/wizact/go-todo-api/internal/user/infrastructure/repository"
)

func Test_NewUserAccountService(t *testing.T) {
	seedUserList := init_users(t)
	u := seedUserList[0]

	evs := NewEmailVerificationService(repository.NewFakeEmailGatewayRepository())

	ur := repository.NewUserMemoryRepository(seedUserList)

	uas := NewUserAccountService(ur, evs)

	_, err := uas.RegisterNewUser(u)

	if err == nil {
		t.Error(err)
	}
}

func init_users(t *testing.T) []aggregate.User {
	ua := aggregate.NewUser()
	u := ua.User()
	u.FirstName = "John"
	u.LastName = "Doe"
	u.Email = "john.doe@example.com"
	ua.SetUser(u)

	var seedUserList []aggregate.User
	seedUserList = append(seedUserList, ua)
	return seedUserList
}
