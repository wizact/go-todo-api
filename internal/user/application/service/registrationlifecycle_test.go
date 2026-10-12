package service

import (
	"sync"
	"testing"
	"time"
)

func TestRegistration_Done_RepeatedCallsReturnWithoutWorker(t *testing.T) {
	t.Parallel()

	registration := NewRegistration(nil, nil, DefaultRegistrationConfig())
	returned := make(chan struct{})
	go func() {
		registration.Done()
		registration.Done()
		close(returned)
	}()

	select {
	case <-returned:
	case <-time.After(time.Second):
		t.Fatal("Done() blocked without a background worker")
	}
}

func TestRegistration_Done_ConcurrentCallsReturn(t *testing.T) {
	t.Parallel()

	registration := NewRegistration(nil, nil, DefaultRegistrationConfig())
	var calls sync.WaitGroup
	returned := make(chan struct{})
	for range 8 {
		calls.Go(registration.Done)
	}
	go func() {
		calls.Wait()
		close(returned)
	}()

	select {
	case <-returned:
	case <-time.After(time.Second):
		t.Fatal("concurrent Done() calls blocked")
	}
}
