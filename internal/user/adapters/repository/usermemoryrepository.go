package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/wizact/go-todo-api/internal/user/application/registration"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/output/repository"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrFailedToAddUser    = errors.New("failed to create user")
	ErrFailedToUpdateUser = errors.New("failed to update user")
)

type MemoryUserRepository struct {
	Users                     map[uuid.UUID]ua.User
	registrationVerifications map[uuid.UUID]registration.Verification
	mutex                     sync.RWMutex
}

func NewMemoryUserRepository(seedUserList []ua.User) *MemoryUserRepository {
	repository := &MemoryUserRepository{
		Users:                     make(map[uuid.UUID]ua.User),
		registrationVerifications: make(map[uuid.UUID]registration.Verification),
	}

	for _, user := range seedUserList {
		repository.Users[user.ID()] = user
	}

	return repository
}

func (r *MemoryUserRepository) FindByID(ctx context.Context, id uuid.UUID) (ua.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if user, ok := r.Users[id]; ok {
		return user, nil
	}
	return ua.User{}, ErrUserNotFound
}

func (r *MemoryUserRepository) FindByEmail(ctx context.Context, email string) (ua.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, v := range r.Users {
		if v.Email() == email {
			return v, nil
		}
	}

	return ua.User{}, ErrUserNotFound
}

func (r *MemoryUserRepository) Create(ctx context.Context, user ua.User) (ua.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.Users == nil {
		r.Users = make(map[uuid.UUID]ua.User)
	}

	if _, ok := r.Users[user.ID()]; ok {
		return ua.User{}, fmt.Errorf("user already exists: %w", ErrFailedToAddUser)
	}

	r.Users[user.ID()] = user

	return user, nil
}

func (r *MemoryUserRepository) Update(ctx context.Context, user ua.User) (ua.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.Users == nil {
		r.Users = make(map[uuid.UUID]ua.User)
	}

	if _, ok := r.Users[user.ID()]; !ok {
		return ua.User{}, fmt.Errorf("user does not exist: %w", ErrFailedToUpdateUser)
	}

	r.Users[user.ID()] = user

	return user, nil
}

func (r *MemoryUserRepository) SaveRegistrationVerification(
	ctx context.Context,
	verification registration.Verification,
) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.registrationVerifications == nil {
		r.registrationVerifications = make(map[uuid.UUID]registration.Verification)
	}
	r.registrationVerifications[verification.UserID] = verification
	return nil
}

func (r *MemoryUserRepository) FindRegistrationVerification(
	ctx context.Context,
	userID uuid.UUID,
) (registration.Verification, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	verification, ok := r.registrationVerifications[userID]
	if !ok {
		return registration.Verification{}, repositoryport.ErrRegistrationVerificationNotFound
	}

	return verification, nil
}

func (r *MemoryUserRepository) CompleteRegistration(
	ctx context.Context,
	user ua.User,
	verificationDigest string,
) (ua.User, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, ok := r.Users[user.ID()]; !ok {
		return ua.User{}, ErrUserNotFound
	}
	verification, ok := r.registrationVerifications[user.ID()]
	if !ok || verification.SecretDigest != verificationDigest {
		return ua.User{}, repositoryport.ErrRegistrationVerificationNotFound
	}

	r.Users[user.ID()] = user
	delete(r.registrationVerifications, user.ID())
	return user, nil
}
