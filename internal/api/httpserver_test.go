package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"

	userportmocks "github.com/wizact/go-todo-api/internal/user/ports/mocks"
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
