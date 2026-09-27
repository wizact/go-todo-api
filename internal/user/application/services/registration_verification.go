package service

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const registrationVerificationHashCost = 14

func issueRegistrationVerification() (secret string, digest string, err error) {
	secret = uuid.NewString()
	hash, err := bcrypt.GenerateFromPassword([]byte(secret), registrationVerificationHashCost)
	if err != nil {
		return "", "", fmt.Errorf("hash registration verification secret: %w", err)
	}

	return secret, string(hash), nil
}

func matchesRegistrationVerification(digest, secret string) bool {
	return bcrypt.CompareHashAndPassword([]byte(digest), []byte(secret)) == nil
}
