package service

import "github.com/google/uuid"

type Registration interface {
	GetRegistrationVerificationEmailData(uid uuid.UUID) (map[string]string, error)
	Done()
}
