package models

import (
	"testing"

	domainmodel "github.com/wizact/go-todo-api/internal/user/domain/models"
)

func TestUserToDomainModel_PreservesPhoneNumber(t *testing.T) {
	t.Parallel()

	input := User{
		FirstName:        "Ada",
		LastName:         "Lovelace",
		DateOfBirth:      "1815-12-10T00:00:00Z",
		Email:            "ada@example.com",
		PhoneCountryCode: "+44",
		PhoneAreaCode:    "20",
		PhoneNumber:      "12345678",
	}

	user, appErr := input.ToDomainModel()
	if appErr != nil {
		t.Fatalf("ToDomainModel() error = %v", appErr)
	}

	want := domainmodel.NewPhoneNumber("+44", "20", "12345678")
	domainUser := user.User()
	if got := domainUser.Phone(); !got.IsEqual(want) {
		t.Fatalf("phone number = %#v, want %#v", got, want)
	}
}
