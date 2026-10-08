package specification

import "testing"

func Test_Not_Given_invalid_object_with_no_name_provided_When_and_spec_is_executed_Then_return_true(t *testing.T) {
	spec := NewNot[UserUnderTest](
		Func[UserUnderTest](HasName),
	)

	uut := UserUnderTest{FirstName: "", LastName: "", FullName: "", Email: "foo@bar.baz"}

	v := spec.IsValid(uut)

	if !v {
		t.Errorf("AndSpecification expected: %v, got: %v", true, v)
	}
}
