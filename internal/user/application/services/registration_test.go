package service

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

	"github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	"github.com/wizact/go-todo-api/internal/user/ports/mocks"
)

func TestRegistration_VerifyUserRegistration_HashMismatch_ReturnsDomainError(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	userID := uuid.New()
	userAccount.EXPECT().
		GetUserById(gomock.Any(), userID).
		Return(aggregate.NewUser(), nil)
	registration := NewRegisteration(userAccount)

	err := registration.VerifyUserRegistration(context.Background(), userID, "invalid-hash")

	if !errors.Is(err, domain.ErrVerificationHashMismatch) {
		t.Errorf("error = %v, want %v", err, domain.ErrVerificationHashMismatch)
	}
}

func TestRegistration_VerifyUserRegistration_PersistsVerifiedUser(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	user := aggregate.NewUser()
	token := user.Token()
	token.RefreshVerificationToken()
	user.SetToken(token)
	hash, err := hashRegistrationVerificationSecret(token.VerificationToken())
	if err != nil {
		t.Fatalf("create verification hash: %v", err)
	}

	userAccount.EXPECT().
		GetUserById(gomock.Any(), user.UserId()).
		Return(user, nil)
	userAccount.EXPECT().
		UpdateUser(gomock.Any(), gomock.Cond(func(updated aggregate.User) bool {
			return updated.IsActive() && updated.HasVerifiedEmail()
		})).
		Return(user, nil)
	registration := NewRegisteration(userAccount)

	err = registration.VerifyUserRegistration(context.Background(), user.UserId(), hash)

	if err != nil {
		t.Fatalf("verify registration: %v", err)
	}
}
