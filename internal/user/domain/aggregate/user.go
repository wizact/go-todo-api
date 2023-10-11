package aggregate

import (
	"errors"

	"github.com/google/uuid"
	model "github.com/wizact/go-todo-api/internal/user/domain/model"
)

// User aggregate with User as it's root entity
type User struct {
	user             *model.User
	role             *model.Role
	hasVerifiedEmail bool
	isActive         bool
	isDeleted        bool
}

// NewDefaultUser creates a new user with an auto generated uuid and limited role
func NewDefaultUser() User {
	u := model.NewUser()
	r := model.RoleFactory{}.CreateUserWithLimitedRole()
	return User{
		user: &u,
		role: &r,
	}
}

// UserId gets the id of the user as aggregate root identity
func (u *User) UserId() uuid.UUID {
	return u.user.ID
}

// User gets the user as aggregate root
func (u *User) User() model.User {
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

// Role gets the user role entity
func (u *User) Role() model.Role {
	return *u.role
}

// SetRole sets the user role
func (u *User) SetRole(nr model.Role) {
	u.role = &nr
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

// IsDeleted gets the user is deleted flag
func (u *User) IsDeleted() bool {
	return u.isDeleted
}

// SetIsDeleted sets the user is deleted flag
func (u *User) SetIsDeleted(b bool) {
	u.isDeleted = b
}

// IsValid checks if the user is valid
func (u *User) IsValid() bool {
	return u.user.IsValid() && u.role.IsValid()
}
