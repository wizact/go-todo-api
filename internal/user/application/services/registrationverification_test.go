package service

import "testing"

func TestHashRegistrationVerificationSecret_ReturnsMatchingDigest(t *testing.T) {
	t.Parallel()

	digest, err := hashRegistrationVerificationSecret("raw-secret")
	if err != nil {
		t.Fatalf("hashRegistrationVerificationSecret() error = %v", err)
	}

	if !matchesRegistrationVerification(digest, "raw-secret") {
		t.Fatal("registration verification digest does not match secret")
	}
}

func TestIssueRegistrationVerification_ReturnsMatchingSecret(t *testing.T) {
	t.Parallel()

	secret, digest, err := issueRegistrationVerification()
	if err != nil {
		t.Fatalf("issueRegistrationVerification() error = %v", err)
	}

	if !matchesRegistrationVerification(digest, secret) {
		t.Fatal("issued secret does not match verification digest")
	}
}

func TestIssueRegistrationVerification_DoesNotStoreRawSecret(t *testing.T) {
	t.Parallel()

	secret, digest, err := issueRegistrationVerification()
	if err != nil {
		t.Fatalf("issueRegistrationVerification() error = %v", err)
	}

	if digest == secret {
		t.Fatal("verification digest contains the raw secret")
	}
}

func TestMatchesRegistrationVerification_WrongSecret_DoesNotMatch(t *testing.T) {
	t.Parallel()

	_, digest, err := issueRegistrationVerification()
	if err != nil {
		t.Fatalf("issueRegistrationVerification() error = %v", err)
	}

	if matchesRegistrationVerification(digest, "wrong-secret") {
		t.Fatal("wrong secret matched verification digest")
	}
}
