package model

import (
	"net/http"
	"testing"

	domainmodel "github.com/wizact/go-todo-api/internal/user/domain/model"
)

func TestUserToDomainModel_InvalidDetails_ReturnsBadRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		firstName string
		lastName  string
		email     string
	}{
		{name: "missing name", email: "ada@example.com"},
		{name: "invalid email", firstName: "Ada", lastName: "Lovelace", email: "invalid"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			input := User{
				FirstName:   tt.firstName,
				LastName:    tt.lastName,
				DateOfBirth: "1815-12-10T00:00:00Z",
				Email:       tt.email,
			}

			_, appErr := input.ToDomainModel()
			if appErr == nil || appErr.Code != http.StatusBadRequest {
				t.Fatalf("error = %#v, want status %d", appErr, http.StatusBadRequest)
			}
		})
	}
}

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
