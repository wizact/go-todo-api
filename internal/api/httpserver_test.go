package api

import (
	"errors"
	dbinfra "github.com/wizact/go-todo-api/internal/infra/db"
	"github.com/wizact/go-todo-api/pkg/eventlibrary/pubsub"
	"net/http"
	"net/http/httptest"
	"path/filepath"
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

func TestNewUserModuleReturnsAdapterConfigurationErrors(t *testing.T) {
	t.Setenv("TODOAPI_DBPATH", "")
	t.Setenv("TODOAPI_NATSURL", "")
	config := Config{PublicBaseURL: "https://example.com"}
	if m, err := newUserModule(config); m != nil || !errors.Is(err, dbinfra.ErrResolveConnectionString) {
		t.Fatalf("missing database: module=%v error=%v", m, err)
	}
	t.Setenv("TODOAPI_DBPATH", filepath.Join(t.TempDir(), "users.db"))
	if m, err := newUserModule(config); m != nil || !errors.Is(err, pubsub.ErrResolveNATSOptions) {
		t.Fatalf("missing NATS: module=%v error=%v", m, err)
	}
	t.Setenv("TODOAPI_NATSURL", "nats://localhost:4222")
	m, err := newUserModule(config)
	if err != nil {
		t.Fatal(err)
	}
	m.Done()
}

func TestStartServerRejectsInvalidConfigurationBeforeListening(t *testing.T) {
	if err := StartServer("localhost", "0", false, Config{}); err == nil {
		t.Fatal("expected configuration error")
	}
}
