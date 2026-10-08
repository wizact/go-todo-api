package domain

import "github.com/wizact/go-todo-api/internal/user/domain"

// UserDomainEvent preserves the public event-library contract while ownership
// moves to the user domain.
type UserDomainEvent = domain.UserRegisteredEvent
