package model

import (
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

func NewEmptyUser() User {
	return User{
		id:    uuid.New(),
		phone: NewEmptyPhoneNumber(),
	}
}

func NewUser(
	id uuid.UUID,
	firstName string,
	lastName string,
	dateOfBirth time.Time,
	email string,
	phone PhoneNumber) User {

	u := User{}

	u.SetID(id)
	u.SetName(firstName, lastName)
	u.SetDateOfBirth(dateOfBirth)
	u.SetEmail(email)
	u.SetPhone(phone)

	return u

}

func (u *User) ID() uuid.UUID      { return u.id }
func (u *User) SetID(id uuid.UUID) { u.id = id }

func (u *User) Name() (string, string) { return u.firstName, u.lastName }
func (u *User) SetName(fn, ln string) {
	u.firstName = fn
	u.lastName = ln
}

func (u *User) DateOfBirth() time.Time       { return u.dateOfBirth }
func (u *User) SetDateOfBirth(dob time.Time) { u.dateOfBirth = dob }

func (u *User) Email() string      { return u.email }
func (u *User) SetEmail(em string) { u.email = em }

func (u *User) Phone() PhoneNumber      { return u.phone }
func (u *User) SetPhone(ph PhoneNumber) { u.phone = ph }

func (u User) IsValid() bool {
	spec := sp.NewAndSpecification[User](
		sp.FunctionSpecification[User](HasName),
		sp.FunctionSpecification[User](HasValidEmail),
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

// TODO: Validation of value object should happen only on the creation side.
func NewPhoneNumber(countryCode, areaCode, number string) PhoneNumber {
	ph := PhoneNumber{}
	ph.SetCountryCode(countryCode)
	ph.SetAreaCode(areaCode)
	ph.SetNumber(number)
	return ph
}

func (p *PhoneNumber) CountryCode() string      { return p.countryCode }
func (p *PhoneNumber) SetCountryCode(cc string) { p.countryCode = cc }

func (p *PhoneNumber) AreaCode() string      { return p.areaCode }
func (p *PhoneNumber) SetAreaCode(ac string) { p.areaCode = ac }

func (p *PhoneNumber) Number() string      { return p.number }
func (p *PhoneNumber) SetNumber(ph string) { p.number = ph }

// IsEqual checks whether or not two instances of PhoneNumber value object are equal or not by comparing all elements of the value objects with each other
func (p PhoneNumber) IsEqual(p2 PhoneNumber) bool {
	return p.countryCode == p2.countryCode && p.areaCode == p2.areaCode && p.number == p2.number
}
