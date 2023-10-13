package model

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUser_IsValid(t *testing.T) {
	type user struct {
		ID          uuid.UUID
		FirstName   string
		LastName    string
		DateOfBirth time.Time
		Email       string
		Phone       PhoneNumber
	}
	tests := []struct {
		name   string
		fields user
		want   bool
	}{
		{"invalid user with no email", user{FirstName: "foo", LastName: "bar", Email: ""}, false},
		{"invalid user with no valid email", user{FirstName: "foo", LastName: "bar", Email: "invalidemail"}, false},
		{"invalid user with no name", user{FirstName: "", LastName: "", Email: "foo@bar.baz"}, false},
		{"valid user", user{FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:          tt.fields.ID,
				FirstName:   tt.fields.FirstName,
				LastName:    tt.fields.LastName,
				DateOfBirth: tt.fields.DateOfBirth,
				Email:       tt.fields.Email,
				Phone:       tt.fields.Phone,
			}
			if got := u.IsValid(); got != tt.want {
				t.Errorf("User.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_IsTheSameUserAs(t *testing.T) {
	type user struct {
		ID          uuid.UUID
		FirstName   string
		LastName    string
		DateOfBirth time.Time
		Email       string
		Phone       PhoneNumber
	}

	user2 := NewUser()

	tests := []struct {
		name   string
		fields user
		want   bool
	}{
		{"Users with the same id", user{ID: user2.ID, FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, true},
		{"Users with the different id", user{ID: uuid.New(), FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, false},
		{"Users with the different id and same email", user{ID: uuid.New(), FirstName: "foo", LastName: "bar", Email: "foo@bar.baz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:          tt.fields.ID,
				FirstName:   tt.fields.FirstName,
				LastName:    tt.fields.LastName,
				DateOfBirth: tt.fields.DateOfBirth,
				Email:       tt.fields.Email,
				Phone:       tt.fields.Phone,
			}
			if got := u.IsTheSameUserAs(user2); got != tt.want {
				t.Errorf("User.IsTheSameUserAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasName(t *testing.T) {
	user1 := NewUser()
	user2 := user1
	user2.FirstName = "foo"
	user3 := user1
	user3.LastName = "bar"
	tests := []struct {
		name string
		user User
		want bool
	}{
		{"user with no FirstName and LastName", user1, false},
		{"user with FirstName but no LastName", user2, true},
		{"user with LastName but no FirstName", user3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasName(tt.user); got != tt.want {
				t.Errorf("HasName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasValidEmail(t *testing.T) {
	tests := []struct {
		name string
		user User
		want bool
	}{
		{"user with no email", User{Email: ""}, false},
		{"user with no valid email", User{Email: "invalidemail"}, false},
		{"user with valid email", User{Email: "foo@bar.baz"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValidEmail(tt.user); got != tt.want {
				t.Errorf("HasValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPhoneNumber_IsEqual(t *testing.T) {
	type fields struct {
		CountryCode string
		AreaCode    string
		Number      string
	}
	p2 := NewPhoneNumber()
	p2.CountryCode = "+64"
	p2.AreaCode = "021"
	p2.Number = "123456"

	tests := []struct {
		name   string
		fields fields
		phone  PhoneNumber
		want   bool
	}{
		{"valid phone number", fields{CountryCode: "+64", AreaCode: "021", Number: "123456"}, p2, true},
		{"country code mismatch", fields{CountryCode: "+61", AreaCode: "021", Number: "123456"}, p2, false},
		{"area code mismatch", fields{CountryCode: "+64", AreaCode: "022", Number: "123456"}, p2, false},
		{"number mismatch", fields{CountryCode: "+64", AreaCode: "021", Number: "123457"}, p2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PhoneNumber{
				CountryCode: tt.fields.CountryCode,
				AreaCode:    tt.fields.AreaCode,
				Number:      tt.fields.Number,
			}
			if got := p.IsEqual(tt.phone); got != tt.want {
				t.Errorf("PhoneNumber.IsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
