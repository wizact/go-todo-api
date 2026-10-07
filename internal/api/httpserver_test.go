package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"

	userportmocks "github.com/wizact/go-todo-api/internal/user/ports/mocks"
	"github.com/wizact/go-todo-api/pkg/communication"
)

func TestRegisterRoutes_UsesProvidedUserServices(t *testing.T) {
	controller := gomock.NewController(t)
	userAccount := userportmocks.NewMockUserAccountUseCase(controller)
	registration := userportmocks.NewMockRegistration(controller)
	userID := uuid.New()
	registration.EXPECT().
		VerifyUserRegistration(gomock.Any(), userID, "verification-token").
		Return(nil)
	router := mux.NewRouter()
	registerRoutes(router, userAccount, registration)
	request := httptest.NewRequest(
		http.MethodPost,
		"/users/verify-registration?uid="+userID.String()+"&token=verification-token",
		nil,
	)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
}

func TestRegisterBackgroundServices_InvalidConfigReturnsError(t *testing.T) {
	controller := gomock.NewController(t)
	registration := userportmocks.NewMockRegistration(controller)

	err := registerBackgroundServices(registration, communication.Config{SendGridEnabled: true})

	if err == nil {
		t.Fatal("registerBackgroundServices() error = nil, want configuration error")
	}
}
