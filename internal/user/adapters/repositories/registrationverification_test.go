package repository

import "testing"

func TestSqliteRegistrationVerification_TableName_ReturnsRegistrationVerificationTable(t *testing.T) {
	t.Parallel()

	got := (SqliteRegistrationVerification{}).TableName()
	want := "user_registration_verifications"
	if got != want {
		t.Fatalf("TableName() = %q, want %q", got, want)
	}
}
