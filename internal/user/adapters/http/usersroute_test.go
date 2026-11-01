package httpadapter_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/mock/gomock"

	controller "github.com/wizact/go-todo-api/internal/user/adapters/controller"
	handlers "github.com/wizact/go-todo-api/internal/user/adapters/http"
	applicationregistration "github.com/wizact/go-todo-api/internal/user/application/registration"
	"github.com/wizact/go-todo-api/internal/user/domain"
	aggregate "github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	model "github.com/wizact/go-todo-api/internal/user/domain/model"
	"github.com/wizact/go-todo-api/internal/user/ports/mocks"
	repositoryport "github.com/wizact/go-todo-api/internal/user/ports/repository"
)

const registeredUserID = "50dff1d7-957c-4498-b286-a3f6c8f517a2"

type routeMocks struct {
	userAccount  *mocks.MockUserAccountUseCase
	registration *mocks.MockRegistration
}

func TestUserRouteFactory_CreateUserRoute_UsesProvidedDependencies(t *testing.T) {
	gomockController := gomock.NewController(t)
	userAccount := mocks.NewMockUserAccountUseCase(gomockController)
	registration := mocks.NewMockRegistration(gomockController)
	userAccount.EXPECT().
		RegisterNewUser(gomock.Any(), gomock.Any()).
		Return(registeredUser(), nil)

	userRoute := handlers.UserRouteFactory{
		UserAccountUseCase: userAccount,
		Registration:       registration,
	}.CreateUserRoute()
	router := mux.NewRouter()
	userRoute.SetupRoutes("/users", router)
	request := httptest.NewRequest(http.MethodPost, "/users", loadFixture(t, "registeruser.json"))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusCreated)
	}
}

func TestUserRoute_RegisterUser_Status(t *testing.T) {
	tests := []struct {
		name       string
		fixture    string
		wantStatus int
		setupMocks func(routeMocks)
	}{
		{
			name:       "returns created for valid registration",
			fixture:    "registeruser.json",
			wantStatus: http.StatusCreated,
			setupMocks: func(m routeMocks) {
				m.userAccount.EXPECT().
					RegisterNewUser(gomock.Any(), gomock.Any()).
					Return(registeredUser(), nil)
			},
		},
		{
			name:       "returns bad request for malformed JSON",
			fixture:    "registerusermalformed.json",
			wantStatus: http.StatusBadRequest,
			setupMocks: func(routeMocks) {},
		},
		{
			name:       "returns internal server error when persistence fails",
			fixture:    "registeruser.json",
			wantStatus: http.StatusInternalServerError,
			setupMocks: func(m routeMocks) {
				m.userAccount.EXPECT().
					RegisterNewUser(gomock.Any(), gomock.Any()).
					Return(aggregate.User{}, domain.ErrUserPersistence)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newUserRouter(t, tt.setupMocks)
			request := httptest.NewRequest(http.MethodPost, "/users", loadFixture(t, tt.fixture))
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", response.Code, tt.wantStatus)
			}
		})
	}
}

func TestUserRoute_VerifyRegistration_Status(t *testing.T) {
	validID := uuid.MustParse(registeredUserID)
	tests := []struct {
		name       string
		query      string
		wantStatus int
		setupMocks func(routeMocks)
	}{
		{
			name:       "returns OK for valid verification",
			query:      "?uid=" + validID.String() + "&token=verification-token",
			wantStatus: http.StatusOK,
			setupMocks: func(m routeMocks) {
				m.registration.EXPECT().
					VerifyUserRegistration(gomock.Any(), validID, "verification-token").
					Return(nil)
			},
		},
		{
			name:       "returns bad request for invalid user ID",
			query:      "?uid=invalid&token=verification-token",
			wantStatus: http.StatusBadRequest,
			setupMocks: func(routeMocks) {},
		},
		{
			name:       "returns bad request when token is missing",
			query:      "?uid=" + validID.String(),
			wantStatus: http.StatusBadRequest,
			setupMocks: func(routeMocks) {},
		},
		{
			name:       "returns bad request for mismatched verification token",
			query:      "?uid=" + validID.String() + "&token=invalid-token",
			wantStatus: http.StatusBadRequest,
			setupMocks: func(m routeMocks) {
				m.registration.EXPECT().
					VerifyUserRegistration(gomock.Any(), validID, "invalid-token").
					Return(applicationregistration.ErrVerificationSecretMismatch)
			},
		},
		{
			name:       "returns bad request for consumed verification token",
			query:      "?uid=" + validID.String() + "&token=consumed-token",
			wantStatus: http.StatusBadRequest,
			setupMocks: func(m routeMocks) {
				m.registration.EXPECT().
					VerifyUserRegistration(gomock.Any(), validID, "consumed-token").
					Return(repositoryport.ErrRegistrationVerificationNotFound)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newUserRouter(t, tt.setupMocks)
			request := httptest.NewRequest(http.MethodPost, "/users/verify-registration"+tt.query, nil)
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", response.Code, tt.wantStatus)
			}
		})
	}
}

func newUserRouter(t *testing.T, setupMocks func(routeMocks)) http.Handler {
	t.Helper()

	gomockController := gomock.NewController(t)
	m := routeMocks{
		userAccount:  mocks.NewMockUserAccountUseCase(gomockController),
		registration: mocks.NewMockRegistration(gomockController),
	}

	setupMocks(m)

	userController := controller.New(m.userAccount, m.registration)
	userRoute := handlers.NewUserRoute(userController)
	router := mux.NewRouter()
	userRoute.SetupRoutes("/users", router)

	return router
}

func loadFixture(t *testing.T, name string) *bytes.Reader {
	t.Helper()

	fixture, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatalf("read fixture %s: %v", name, err)
	}

	return bytes.NewReader(fixture)
}

func registeredUser() aggregate.User {
	phone := model.NewPhoneNumber("+64", "23", "123456")
	user := model.RehydrateUser(
		uuid.MustParse(registeredUserID),
		"Foo",
		"Bar",
		time.Date(1983, time.January, 2, 15, 4, 5, 0, time.UTC),
		"foo.bar@example.com",
		phone,
	)
	location := model.NewLocation(173.3002574488138, -41.26595602617756)

	result := aggregate.NewUser()
	result.SetUser(user)
	result.SetLocation(location)
	return result
}

func TestUserRoute_FetchUserByID_Status(t *testing.T) {
	validID := uuid.MustParse(registeredUserID)
	for _, tc := range []struct {
		name, id, authorization string
		want                    int
		setup                   func(routeMocks)
	}{
		{"requires authorization", registeredUserID, "", http.StatusUnauthorized, func(routeMocks) {}},
		{"rejects invalid ID", "invalid", "'Bearer bearer token'", http.StatusBadRequest, func(routeMocks) {}},
		{"returns user", registeredUserID, "'Bearer bearer token'", http.StatusOK, func(m routeMocks) {
			m.userAccount.EXPECT().FetchUserByID(gomock.Any(), validID).Return(registeredUser(), nil)
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			router := newUserRouter(t, tc.setup)
			request := httptest.NewRequest(http.MethodGet, "/users/"+tc.id, nil)
			request.Header.Set("Authorization", tc.authorization)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)
			if response.Code != tc.want {
				t.Fatalf("status = %d, want %d", response.Code, tc.want)
			}
		})
	}
}

func TestUserRoute_RegisterUser_Location(t *testing.T) {
	router := newUserRouter(t, func(m routeMocks) {
		m.userAccount.EXPECT().RegisterNewUser(gomock.Any(), gomock.Any()).Return(registeredUser(), nil)
	})
	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/users", loadFixture(t, "registeruser.json")))
	if got := response.Header().Get("Location"); got != "/users/"+registeredUserID {
		t.Fatalf("Location = %q", got)
	}
}

func TestUserRoute_RegisterUser_Response(t *testing.T) {
	router := newUserRouter(t, func(m routeMocks) {
		m.userAccount.EXPECT().RegisterNewUser(gomock.Any(), gomock.Any()).Return(registeredUser(), nil)
	})
	response := httptest.NewRecorder()
	router.ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/users", loadFixture(t, "registeruser.json")))
	var body map[string]any
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["user_id"] != registeredUserID {
		t.Fatalf("response user_id = %v", body["user_id"])
	}
}
