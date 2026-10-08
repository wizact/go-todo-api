package model

import (
	"errors"
	"net/mail"
	"strings"
	"time"

	"github.com/google/uuid"
	sp "github.com/wizact/go-todo-api/pkg/specification"
)

type User struct {
	id          uuid.UUID
	firstName   string
	lastName    string
	dateOfBirth time.Time
	email       string
	phone       PhoneNumber
}

var ErrInvalidUser = errors.New("user info is not valid")

func NewUser(
	firstName string,
	lastName string,
	dateOfBirth time.Time,
	email string,
	phone PhoneNumber,
) (User, error) {
	u := User{
		id:          uuid.New(),
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dateOfBirth,
		email:       email,
		phone:       phone,
	}
	if !u.IsValid() {
		return User{}, ErrInvalidUser
	}

	return u, nil
}

// RehydrateUser restores a user from persisted state without generating a new identity.
func RehydrateUser(
	id uuid.UUID,
	firstName string,
	lastName string,
	dateOfBirth time.Time,
	email string,
	phone PhoneNumber,
) User {
	return User{
		id:          id,
		firstName:   firstName,
		lastName:    lastName,
		dateOfBirth: dateOfBirth,
		email:       email,
		phone:       phone,
	}
}

func (u User) ID() uuid.UUID { return u.id }

func (u User) Name() (string, string) { return u.firstName, u.lastName }

// ConcatenatedName returns the full name of the user by concatenating the first name and last name
func (u User) ConcatenatedName() string { return u.firstName + u.lastName }

func (u User) DateOfBirth() time.Time { return u.dateOfBirth }

func (u User) Email() string { return u.email }

func (u User) Phone() PhoneNumber { return u.phone }

func (u User) IsValid() bool {
	spec := sp.NewAnd[User](
		sp.Func[User](HasName),
		sp.Func[User](HasValidEmail),
	)

	return spec.IsValid(u)

}

func (u User) IsTheSameUserAs(u2 User) bool {
	return u.id == u2.id
}

func HasName(user User) bool {
	f := strings.Trim(user.firstName, " ")
	l := strings.Trim(user.lastName, " ")
	return f != "" || l != ""
}

func HasValidEmail(user User) bool {
	if user.email == "" {
		return false
	}

	if _, err := mail.ParseAddress(user.email); err != nil {
		return false
	}
	return true
}

type PhoneNumber struct {
	countryCode string
	areaCode    string
	number      string
}

func NewEmptyPhoneNumber() PhoneNumber {
	return PhoneNumber{}
}

func NewPhoneNumber(countryCode, areaCode, number string) PhoneNumber {
	return PhoneNumber{
		countryCode: countryCode,
		areaCode:    areaCode,
		number:      number,
	}
}

func (p PhoneNumber) CountryCode() string { return p.countryCode }

func (p PhoneNumber) AreaCode() string { return p.areaCode }

func (p PhoneNumber) Number() string { return p.number }

// IsEqual checks whether or not two instances of PhoneNumber value object are equal or not by comparing all elements of the value objects with each other
func (p PhoneNumber) IsEqual(p2 PhoneNumber) bool {
	return p.countryCode == p2.countryCode && p.areaCode == p2.areaCode && p.number == p2.number
}
