package model

import (
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
	sp "github.com/wizact/go-todo-api/pkg/specification"
)

type User struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	DateOfBirth time.Time
	Email       string
	Phone       PhoneNumber
}

func NewUser() User {
	return User{
		ID:    uuid.New(),
		Phone: NewPhoneNumber(),
	}
}

func (u User) IsValid() bool {
	spec := sp.NewAndSpecification[User](
		sp.FunctionSpecification[User](HasName),
		sp.FunctionSpecification[User](HasValidEmail),
	)

	return spec.IsValid(u)

}

func (u User) IsTheSameUserAs(u2 User) bool {
	return u.ID == u2.ID
}

func HasName(user User) bool {
	f := strings.Trim(user.FirstName, " ")
	l := strings.Trim(user.LastName, " ")
	return f != "" || l != ""
}

func HasValidEmail(user User) bool {
	if user.Email == "" {
		return false
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return false
	}
	return true
}

type PhoneNumber struct {
	CountryCode string
	AreaCode    string
	Number      string
}

func NewPhoneNumber() PhoneNumber {
	return PhoneNumber{}
}

// IsEqual checks whether or not two instances of PhoneNumber value object are equal or not by comparing all elements of the value objects with each other
func (p PhoneNumber) IsEqual(p2 PhoneNumber) bool {
	return p.CountryCode == p2.CountryCode && p.AreaCode == p2.AreaCode && p.Number == p2.Number
}
