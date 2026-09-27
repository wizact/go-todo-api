package aggregate

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/domain"
	"github.com/wizact/go-todo-api/internal/user/domain/models"
)

func TestUser_GetDomainEventPayload_ReturnsRegisteredEvent(t *testing.T) {
	t.Parallel()

	userID := uuid.New()
	root := NewUser()
	user := model.NewUser(
		userID,
		"Ada",
		"Lovelace",
		time.Date(1815, time.December, 10, 0, 0, 0, 0, time.UTC),
		"ada@example.com",
		model.NewPhoneNumber("+44", "20", "12345678"),
	)
	root.SetUser(user)
	root.SetIsActive(true)
	root.SetHasVerifiedEmail(true)

	want := domain.UserRegisteredEvent{
		ID:               userID,
		FirstName:        "Ada",
		LastName:         "Lovelace",
		Email:            "ada@example.com",
		IsActive:         true,
		HasVerifiedEmail: true,
	}
	var got domain.UserRegisteredEvent = root.GetDomainEventPayload()

	if got != want {
		t.Fatalf("domain event = %#v, want %#v", got, want)
	}
}
