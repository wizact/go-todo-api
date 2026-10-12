package user

import (
	"context"
	repository "github.com/wizact/go-todo-api/internal/user/adapters/repository"
	"github.com/wizact/go-todo-api/internal/user/domain"
	"github.com/wizact/go-todo-api/internal/user/domain/aggregate"
	"strings"
	"testing"
	"time"

	appService "github.com/wizact/go-todo-api/internal/user/application/service"
)

func TestModule_Done_ReturnsWithoutRegistrationWorker(t *testing.T) {
	t.Parallel()

	module := Module{
		registrationService: appService.NewRegistration(nil, nil, appService.DefaultRegistrationConfig()),
	}
	returned := make(chan struct{})
	go func() {
		module.Done()
		close(returned)
	}()

	select {
	case <-returned:
	case <-time.After(time.Second):
		t.Fatal("Module.Done() blocked on registration cleanup")
	}
}

type testPublisher struct{}

func (testPublisher) PublishNewUserRegisteredEvent(context.Context, domain.UserRegisteredEvent) error {
	return nil
}

func TestNewModuleWithDependenciesRejectsInvalidAssembly(t *testing.T) {
	valid := Dependencies{Repository: repository.NewMemoryUserRepository(nil), EventPublisher: testPublisher{}}
	for _, tc := range []struct {
		name    string
		deps    Dependencies
		options []Option
		want    string
	}{
		{"repository", Dependencies{EventPublisher: testPublisher{}}, nil, "repository"},
		{"publisher", Dependencies{Repository: valid.Repository}, nil, "publisher"},
		{"nil option", valid, []Option{nil}, "nil option"},
		{"empty URL", valid, []Option{WithPublicBaseURL("")}, "absolute HTTP"},
		{"invalid final option", valid, []Option{WithPublicBaseURL("https://example.com"), WithPublicBaseURL("ftp://example.com")}, "absolute HTTP"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m, err := NewModuleWithDependencies(tc.deps, tc.options...)
			if m != nil || err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("module=%v error=%v", m, err)
			}
		})
	}
}

func TestModuleOptionsConfigureRegistrationWithoutEnvironment(t *testing.T) {
	t.Setenv("TODOAPI_DBPATH", "")
	t.Setenv("TODOAPI_NATSURL", "")
	for _, tc := range []struct {
		name    string
		options []Option
		want    string
	}{
		{"defaults", nil, "http://localhost:8080"},
		{"last option wins", []Option{WithPublicBaseURL("invalid"), WithPublicBaseURL(" https://example.com/app/ ")}, "https://example.com/app"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			u := aggregate.NewUser()
			repo := repository.NewMemoryUserRepository([]aggregate.User{u})
			m, err := NewModuleWithDependencies(Dependencies{Repository: repo, EventPublisher: testPublisher{}}, tc.options...)
			if err != nil {
				t.Fatal(err)
			}
			data, err := m.UserRegistrationAppService().FetchRegistrationVerificationEmailData(context.Background(), u.ID())
			if err != nil {
				t.Fatal(err)
			}
			if data["base_url"] != tc.want || !strings.HasPrefix(data["verify_email_link"], tc.want+"/users/verify-registration?") {
				t.Fatalf("unexpected email URL: %v", data["base_url"])
			}
			if _, err := repo.FindRegistrationVerification(context.Background(), u.ID()); err != nil {
				t.Fatalf("injected repository did not receive verification: %v", err)
			}
			m.Done()
		})
	}
}
