package service

type Registration interface {
	NewUserRegisteredListener() error
	Done()
}
