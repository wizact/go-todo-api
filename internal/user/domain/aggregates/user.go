package aggregate

import (
	"errors"

	"github.com/google/uuid"
	model "github.com/wizact/go-todo-api/internal/user/domain/models"
)

// User aggregate with User as it's root entity
type User struct {
	user             *model.User
	location         *model.Location
	hasVerifiedEmail bool
	isActive         bool
}

// NewUser creates a new user with an auto generated uuid and limited role
func NewUser() User {
	u := model.NewUser()
	l := model.Location{}
	return User{
		user:     &u,
		location: &l,
	}
}

// GetAggregateEventPayload returns a representation of the aggregate for event processing
func (u *User) GetAggregateEventPayload() interface{} {
	ue := u.User()
	ae := struct {
		ID               uuid.UUID `json:"ID"`
		FirstName        string    `json:"FirstName"`
		LastName         string    `json:"LastName"`
		Email            string    `json:"Email"`
		IsActive         bool      `json:"IsActive"`
		HasVerifiedEmail bool      `json:"HasVerifiedEmail"`
	}{
		ID:               u.UserId(),
		FirstName:        ue.FirstName,
		LastName:         ue.LastName,
		Email:            ue.Email,
		IsActive:         u.isActive,
		HasVerifiedEmail: u.HasVerifiedEmail(),
	}

	return ae
}

// UserId gets the id of the user as aggregate root identity
func (u *User) UserId() uuid.UUID {
	return u.user.ID
}

// User gets the user as aggregate root
func (u *User) User() model.User {
	if u.user == nil {
		um := model.NewUser()
		u.user = &um
	}
	return *u.user
}

// SetUser sets the user
func (u *User) SetUser(nu model.User) {
	u.user = &nu
}

// Email gets the user email
func (u *User) Email() string {
	if u.user != nil {
		return u.user.Email
	}

	return ""
}

// SetUser sets the user
func (u *User) SetEmail(email string) error {
	if u.user == nil {
		return errors.New("user is not instantiated")
	}

	cloned := u.user
	cloned.Email = email

	if !model.HasValidEmail(*cloned) {
		return errors.New("email is not valid")
	}

	u.user.Email = email
	return nil
}

// Location gets the user location value object
func (u *User) Location() model.Location {
	if u.location == nil {
		l := model.Location{}
		u.location = &l
	}

	return *u.location
}

// SetLocation sets the user role
func (u *User) SetLocation(nl model.Location) {
	u.location = &nl
}

// HasVerifiedEmail gets the user has verified email flag
func (u *User) HasVerifiedEmail() bool {
	return u.hasVerifiedEmail
}

// SetHasVerifiedEmail sets the user has verified email flag
func (u *User) SetHasVerifiedEmail(b bool) {
	u.hasVerifiedEmail = b
}

// IsActive gets the user is active flag
func (u *User) IsActive() bool {
	return u.isActive
}

// SetIsActive sets the user is active flag
func (u *User) SetIsActive(b bool) {
	u.isActive = b
}

// IsValid checks if the user is valid
func (u *User) IsValid() bool {
	return (u.user != nil && u.user.IsValid()) && (u.location != nil && u.location.IsValid())
}

// UserEmailView is a snapshot of email information for user aggregate for read-only purposes
type UserEmailView struct {
	id               uuid.UUID
	email            string
	hasVerifiedEmail bool
}

func NewUserEmailView(id uuid.UUID, email string, hasVerifiedEmail bool) UserEmailView {
	return UserEmailView{id: id, email: email, hasVerifiedEmail: hasVerifiedEmail}
}

func (uev UserEmailView) Id() uuid.UUID {
	return uev.id
}

func (uev UserEmailView) Email() string {
	return uev.email
}

func (uev UserEmailView) IsEmailVerified() bool {
	return uev.hasVerifiedEmail
}
