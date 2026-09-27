package domain

import "errors"

var (
	ErrInvalidUser              = errors.New("user info is not valid")
	ErrRegistrationFailed       = errors.New("failed to register user")
	ErrUserPersistence          = errors.New("failed to persist user")
	ErrEmailAlreadyExists       = errors.New("email already registered")
	ErrUserLookupFailed         = errors.New("cannot get user")
	ErrUserIDNotFound           = errors.New("user id does not exist")
	ErrUserEmailNotFound        = errors.New("user email does not exist")
	ErrVerificationHashMismatch = errors.New("verification hash does not match")
)
