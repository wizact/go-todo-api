package specification

import (
	"net/mail"
	"strings"
	"testing"
)

func Test_AndSpecification_Given_valid_object_When_and_spec_is_executed_Then_return_true(t *testing.T) {
	spec := NewAndSpecification[UserUnderTest](
		FunctionSpecification[UserUnderTest](HasName),
		FunctionSpecification[UserUnderTest](HasValidEmail),
	)

	uut := UserUnderTest{FirstName: "Foo", LastName: "Bar", FullName: "", Email: "foo@bar.baz"}

	v := spec.IsValid(uut)

	if !v {
		t.Errorf("AndSpecification expected: %v, got: %v", true, v)
	}
}

func Test_AndSpecification_Given_at_least_one_invalid_field_in_object_When_and_spec_is_executed_Then_return_false(t *testing.T) {
	spec := NewAndSpecification[UserUnderTest](
		FunctionSpecification[UserUnderTest](HasName),
		FunctionSpecification[UserUnderTest](HasValidEmail),
	)

	uut := UserUnderTest{FirstName: "", LastName: "", FullName: "", Email: "foo@bar.baz"}

	v := spec.IsValid(uut)

	if v {
		t.Errorf("AndSpecification expected: %v, got: %v", false, v)
	}
}

type UserUnderTest struct {
	FirstName string
	LastName  string
	FullName  string
	Email     string
}

func HasValidEmail(user UserUnderTest) bool {
	if user.Email == "" {
		return false
	}

	if _, err := mail.ParseAddress(user.Email); err != nil {
		return false
	}
	return true
}

func HasName(user UserUnderTest) bool {
	f := strings.Trim(user.FirstName, " ")
	l := strings.Trim(user.LastName, " ")
	fn := strings.Trim(user.FullName, " ")
	return f != "" || l != "" || fn != ""
}
