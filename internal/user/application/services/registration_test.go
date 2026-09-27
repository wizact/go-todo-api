package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"

	service "github.com/wizact/go-todo-api/internal/user/application/services"
	"github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregates"
	"github.com/wizact/go-todo-api/internal/user/ports/mocks"
	event "github.com/wizact/go-todo-api/pkg/event-library/user/events"
)

func TestRegistration_VerifyUserRegistration_HashMismatch_ReturnsDomainError(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(controller)
	userID := uuid.New()
	userAccount.EXPECT().
		GetUserById(gomock.Any(), userID).
		Return(aggregate.NewUser(), nil)
	registration := service.NewRegisteration(event.UserEventClientMock{}, userAccount)

	err := registration.VerifyUserRegistration(context.Background(), userID, "invalid-hash")

	if !errors.Is(err, domain.ErrVerificationHashMismatch) {
		t.Errorf("error = %v, want %v", err, domain.ErrVerificationHashMismatch)
	}
}
