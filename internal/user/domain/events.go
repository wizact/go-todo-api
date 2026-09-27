package domain

import "github.com/google/uuid"

// UserRegisteredEvent represents the user aggregate after registration.
type UserRegisteredEvent struct {
	ID               uuid.UUID
	FirstName        string
	LastName         string
	Email            string
	IsActive         bool
	HasVerifiedEmail bool
}
