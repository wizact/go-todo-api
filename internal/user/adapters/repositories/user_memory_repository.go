package repository

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrFailedToAddUser    = errors.New("failed to create user")
	ErrFailedToUpdateUser = errors.New("failed to update user")
)

type UserMemoryRepository struct {
	Users map[uuid.UUID]ua.User
	*sync.Mutex
}

func NewUserMemoryRepository(seedUserList []ua.User) UserMemoryRepository {
	sul := UserMemoryRepository{
		Users: make(map[uuid.UUID]ua.User),
	}

	for _, user := range seedUserList {
		sul.Users[uuid.UUID(user.UserId())] = user
	}

	return sul
}

func (r UserMemoryRepository) FindById(ctx context.Context, id uuid.UUID) (ua.User, error) {
	if user, ok := r.Users[uuid.UUID(id)]; ok {
		return user, nil
	}
	return ua.User{}, ErrUserNotFound
}

func (r UserMemoryRepository) FindByEmail(ctx context.Context, email string) (ua.User, error) {
	for _, v := range r.Users {
		if v.Email() == email {
			return v, nil
		}
	}

	return ua.User{}, ErrUserNotFound
}
func (r UserMemoryRepository) Create(ctx context.Context, user ua.User) (ua.User, error) {
	if r.Users == nil {
		r.Lock()
		r.Users = make(map[uuid.UUID]ua.User)
		r.Unlock()
	}

	if _, ok := r.Users[uuid.UUID(user.UserId())]; ok {
		return ua.User{}, fmt.Errorf("user already exists: %w", ErrFailedToAddUser)
	}

	r.Lock()
	r.Users[uuid.UUID(user.UserId())] = user
	r.Unlock()

	return user, nil
}
func (r UserMemoryRepository) Update(ctx context.Context, user ua.User) (ua.User, error) {
	if r.Users == nil {
		r.Lock()
		r.Users = make(map[uuid.UUID]ua.User)
		r.Unlock()
	}

	if _, ok := r.Users[uuid.UUID(user.UserId())]; !ok {
		return ua.User{}, fmt.Errorf("user does not exist: %w", ErrFailedToUpdateUser)
	}

	r.Lock()
	r.Users[uuid.UUID(user.UserId())] = user
	r.Unlock()

	return user, nil
}
