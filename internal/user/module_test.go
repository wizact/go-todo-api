package user

import (
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
