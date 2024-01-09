package repository

import (
	"context"

	"github.com/google/uuid"
	ua "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
)

type UserSqlLiteRepository struct {
}

func NewUserSqlLiteRepository() *UserSqlLiteRepository {
	return &UserSqlLiteRepository{}
}

func (r *UserSqlLiteRepository) FindById(ctx context.Context, id uuid.UUID) (ua.User, error) {
	return ua.User{}, nil
}

func (r *UserSqlLiteRepository) FindByEmail(ctx context.Context, email string) (ua.User, error) {
	return ua.User{}, nil
}

func (r *UserSqlLiteRepository) Create(ctx context.Context, user ua.User) (ua.User, error) {
	return ua.User{}, nil
}

func (r *UserSqlLiteRepository) Update(ctx context.Context, user ua.User) (ua.User, error) {
	return ua.User{}, nil
}
