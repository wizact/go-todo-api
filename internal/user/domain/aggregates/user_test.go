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
	user := model.RehydrateUser(
		userID,
		"Ada",
		"Lovelace",
		time.Date(1815, time.December, 10, 0, 0, 0, 0, time.UTC),
		"ada@example.com",
		model.NewPhoneNumber("+44", "20", "12345678"),
	)
	root.SetUser(user)
	root.VerifyRegistration()

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

func TestUser_VerifyRegistration_ActivatesVerifiedUser(t *testing.T) {
	t.Parallel()

	user := NewUser()

	user.VerifyRegistration()

	want := [2]bool{true, true}
	if got := [2]bool{user.IsActive(), user.HasVerifiedEmail()}; got != want {
		t.Fatalf("registration state = %v, want %v", got, want)
	}
}

func TestUser_IsValid_DoesNotRequireRegistrationCredential(t *testing.T) {
	t.Parallel()

	user := model.RehydrateUser(
		uuid.New(),
		"Ada",
		"Lovelace",
		time.Date(1815, time.December, 10, 0, 0, 0, 0, time.UTC),
		"ada@example.com",
		model.NewPhoneNumber("+44", "20", "12345678"),
	)
	location := model.NewLocation(0, 0)
	root := User{user: &user, location: &location}

	if !root.IsValid() {
		t.Fatal("IsValid() = false, want true without registration credential")
	}
}

func TestRehydrateUser_RestoresRegistrationStatus(t *testing.T) {
	t.Parallel()

	want := RegistrationStatus{IsActive: true, HasVerifiedEmail: false}
	user := RehydrateUser(
		model.User{},
		model.NewLocation(0, 0),
		want,
	)

	got := RegistrationStatus{
		IsActive:         user.IsActive(),
		HasVerifiedEmail: user.HasVerifiedEmail(),
	}
	if got != want {
		t.Fatalf("registration status = %v, want %v", got, want)
	}
}
