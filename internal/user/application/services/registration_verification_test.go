package service

import "testing"

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
