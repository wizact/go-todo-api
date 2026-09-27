package service

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const registrationVerificationHashCost = 14

func hashRegistrationVerificationSecret(secret string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(secret), registrationVerificationHashCost)
	if err != nil {
		return "", fmt.Errorf("hash registration verification secret: %w", err)
	}

	return string(hash), nil
}

func issueRegistrationVerification() (secret string, digest string, err error) {
	secret = uuid.NewString()
	digest, err = hashRegistrationVerificationSecret(secret)
	if err != nil {
		return "", "", err
	}

	return secret, digest, nil
}

func matchesRegistrationVerification(digest, secret string) bool {
	return bcrypt.CompareHashAndPassword([]byte(digest), []byte(secret)) == nil
}
