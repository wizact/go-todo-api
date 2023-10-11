package specification

import (
	"testing"
)

func Test_OrSpecification_Given_at_least_one_valid_condition_When_or_spec_is_executed_Then_return_true(t *testing.T) {
	spec := NewOrSpecification[UserUnderTest](
		FunctionSpecification[UserUnderTest](HasName),
		FunctionSpecification[UserUnderTest](HasValidEmail),
	)

	uut := UserUnderTest{FirstName: "", LastName: "", FullName: "", Email: "foo@bar.baz"}

	v := spec.IsValid(uut)

	if !v {
		t.Errorf("OrSpecification expected: %v, got: %v", true, v)
	}
}

func Test_OrSpecification_Given_all_valid_conditions_When_or_spec_is_executed_Then_return_true(t *testing.T) {
	spec := NewOrSpecification[UserUnderTest](
		FunctionSpecification[UserUnderTest](HasName),
		FunctionSpecification[UserUnderTest](HasValidEmail),
	)

	uut := UserUnderTest{FirstName: "foo", LastName: "bar", FullName: "", Email: "foo@bar.baz"}

	v := spec.IsValid(uut)

	if !v {
		t.Errorf("OrSpecification expected: %v, got: %v", true, v)
	}
}

func Test_OrSpecification_Given_no_valid_field_in_object_When_or_spec_is_executed_Then_return_false(t *testing.T) {
	spec := NewOrSpecification[UserUnderTest](
		FunctionSpecification[UserUnderTest](HasName),
		FunctionSpecification[UserUnderTest](HasValidEmail),
	)

	uut := UserUnderTest{FirstName: "", LastName: "", FullName: "", Email: ""}

	v := spec.IsValid(uut)

	if v {
		t.Errorf("OrSpecification expected: %v, got: %v", false, v)
	}
}
