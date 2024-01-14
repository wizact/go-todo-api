package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	repository "github.com/wizact/go-todo-api/internal/user/ports/output/repositories"
	hsm "github.com/wizact/go-todo-api/pkg/http-server-model"
)

var (
	ErrInvalidUser               = hsm.NewAppError(errors.New("user info is not valid"), "user info is not valid", http.StatusBadRequest)
	ErrFailedToRegisterUser      = hsm.NewAppError(errors.New("user info is not valid"), "user info is not valid", http.StatusBadRequest)
	ErrServerErrorToRegisterUser = hsm.NewAppError(errors.New("internal server error"), "internal server error", http.StatusInternalServerError)
	ErrEmailAlreadyExists        = hsm.NewAppError(errors.New("email already registered"), "email already registered", http.StatusBadRequest)

	ErrFailedToGetUser         = hsm.NewAppError(errors.New("cannot get user"), "cannot get user", http.StatusNotFound)
	ErrUserIdDoesNotExist      = hsm.NewAppError(errors.New("user id does not exist"), "user id does not exist", http.StatusNotFound)
	ErrUserByEmailDoesNotExist = hsm.NewAppError(errors.New("user email does not exist"), "user email does not exist", http.StatusNotFound)
)

type UserAccountService struct {
	// repositories and other services
	userRepository repository.UserRepository
}

func NewUserAccountService(ur repository.UserRepository) *UserAccountService {
	ua := &UserAccountService{
		userRepository: ur,
	}

	return ua
}

func (ua *UserAccountService) RegisterNewUser(ctx context.Context, user aggregate.User) (aggregate.User, *hsm.AppError) {
	// Verify the account
	if !user.IsValid() {
		return user, ErrInvalidUser
	}

	// Check if the user does not exist
	u, e := ua.userRepository.FindByEmail(ctx, user.Email())
	if e != nil && !errors.Is(e, ErrUserByEmailDoesNotExist) {
		return user, ErrFailedToRegisterUser
	}

	if u.Email() == user.Email() {
		return user, ErrEmailAlreadyExists
	}

	// Create user with an active status
	user.SetIsActive(true)
	u, e = ua.userRepository.Create(ctx, user)
	if e != nil {
		return user, ErrServerErrorToRegisterUser
	}

	return u, nil
}

func (ua *UserAccountService) GetUserById(ctx context.Context, uid uuid.UUID) (aggregate.User, *hsm.AppError) {
	var u aggregate.User
	u, e := ua.userRepository.FindById(ctx, uid)

	if e != nil {
		// Fallback to generic error
		return u, ErrFailedToGetUser
	}

	return u, nil
}
