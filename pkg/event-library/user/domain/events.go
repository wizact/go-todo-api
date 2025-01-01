package event

import (
	"encoding/json"

	"github.com/google/uuid"
)

// UserDomainEvent a value object represents the domain event for the User aggregate
type UserDomainEvent struct {
	ID               uuid.UUID `json:"ID"`
	FirstName        string    `json:"FirstName"`
	LastName         string    `json:"LastName"`
	Email            string    `json:"Email"`
	IsActive         bool      `json:"IsActive"`
	HasVerifiedEmail bool      `json:"HasVerifiedEmail"`
}

// LoadDomainEventObject unmarshal a byte array and returns UserDomainEvent if successful
func (ude *UserDomainEvent) LoadDomainEventObject(p []byte) error {
	if err := json.Unmarshal(p, ude); err != nil {
		return err
	}
	return nil
}
