package aggregate

import (
	"github.com/google/uuid"
	domainEvent "github.com/wizact/go-todo-api/internal/user/domain"
	model "github.com/wizact/go-todo-api/internal/user/domain/model"
)

// User aggregate with User as it's root entity
type User struct {
	user             *model.User
	location         *model.Location
	hasVerifiedEmail bool
	isActive         bool
}

// RegistrationStatus is the persisted registration state of a user.
type RegistrationStatus struct {
	IsActive         bool
	HasVerifiedEmail bool
}

// NewUser creates an empty aggregate for assembling a new registration.
func NewUser() User {
	u := model.User{}
	l := model.Location{}
	return User{
		user:     &u,
		location: &l,
	}
}

// RehydrateUser restores a user aggregate from persisted state.
func RehydrateUser(user model.User, location model.Location, status RegistrationStatus) User {
	return User{
		user:             &user,
		location:         &location,
		hasVerifiedEmail: status.HasVerifiedEmail,
		isActive:         status.IsActive,
	}
}

// GetAggregateEventPayload returns a representation of the aggregate for event processing
func (u *User) DomainEventPayload() domainEvent.UserRegisteredEvent {
	ue := u.User()
	fn, ln := ue.Name()
	ae := domainEvent.UserRegisteredEvent{
		ID:               u.ID(),
		FirstName:        fn,
		LastName:         ln,
		Email:            ue.Email(),
		IsActive:         u.isActive,
		HasVerifiedEmail: u.HasVerifiedEmail(),
	}

	return ae
}

// ID gets the id of the user as aggregate root identity
func (u *User) ID() uuid.UUID {
	return u.user.ID()
}

// User gets the user as aggregate root
func (u *User) User() model.User {
	if u.user == nil {
		um := model.User{}
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
		return u.user.Email()
	}

	return ""
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

// IsActive gets the user is active flag
func (u *User) IsActive() bool {
	return u.isActive
}

// VerifyRegistration activates the user after their email is verified.
func (u *User) VerifyRegistration() {
	u.isActive = true
	u.hasVerifiedEmail = true
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

func (u UserEmailView) ID() uuid.UUID {
	return u.id
}

func (u UserEmailView) Email() string {
	return u.email
}

func (u UserEmailView) IsEmailVerified() bool {
	return u.hasVerifiedEmail
}
