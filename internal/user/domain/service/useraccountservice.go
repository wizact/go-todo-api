package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	event "github.com/wizact/go-todo-api/internal/user/ports/output/event"
	repository "github.com/wizact/go-todo-api/internal/user/ports/output/repository"
)

type UserAccount struct {
	// repositories and other services
	userRepository     repository.UserRepository
	userEventPublisher event.UserEventPublisher
}

func NewUserAccount(ur repository.UserRepository, publisher event.UserEventPublisher) *UserAccount {
	ua := &UserAccount{
		userRepository:     ur,
		userEventPublisher: publisher,
	}

	return ua
}

func (ua *UserAccount) RegisterNewUser(ctx context.Context, user aggregate.User) (aggregate.User, error) {
	// Verify the account
	if !user.IsValid() {
		return user, domain.ErrInvalidUser
	}

	// Check if the user does not exist
	u, e := ua.userRepository.FindByEmail(ctx, user.Email())
	if e != nil && !errors.Is(e, domain.ErrUserEmailNotFound) {
		return user, fmt.Errorf("%w: find user by email: %w", domain.ErrRegistrationFailed, e)
	}

	if u.Email() == user.Email() {
		return user, domain.ErrEmailAlreadyExists
	}

	u, e = ua.userRepository.Create(ctx, user)
	if e != nil {
		return user, fmt.Errorf("%w: create user: %w", domain.ErrUserPersistence, e)
	}

	// emit events
	err := ua.userEventPublisher.PublishNewUserRegisteredEvent(ctx, user.DomainEventPayload())

	if err != nil {
		log.Printf("failed PublishNewUserRegisteredEvent for %v \n", u.ID())
	}

	return u, nil
}

// FetchUserByID gets a user aggregate by id
func (ua *UserAccount) FetchUserByID(ctx context.Context, uid uuid.UUID) (aggregate.User, error) {
	var u aggregate.User
	u, e := ua.userRepository.FindByID(ctx, uid)

	if e != nil {
		// Fallback to generic error
		return u, fmt.Errorf("%w: find user by ID: %w", domain.ErrUserLookupFailed, e)
	}

	return u, nil
}

// UpdateUser updates a user aggregate
func (ua *UserAccount) UpdateUser(ctx context.Context, user aggregate.User) (aggregate.User, error) {
	u, e := ua.userRepository.Update(ctx, user)

	if e != nil {
		return u, fmt.Errorf("%w: update user: %w", domain.ErrUserPersistence, e)
	}

	return u, nil
}
