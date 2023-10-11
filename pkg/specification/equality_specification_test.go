package specification

import "testing"

func Test_IsMoreOrEqualThan_Specification(t *testing.T) {

	p := func(t equalityObjectUnderTest) int64 {
		return t.age
	}
	spec := NewIsMoreOrEqualThan[equalityObjectUnderTest, int64](18, predicate[equalityObjectUnderTest, int64](p))

	tests := map[string]struct {
		input int64
		want  bool
	}{
		"younger":  {input: 17, want: false},
		"same age": {input: 18, want: true},
		"older":    {input: 19, want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			eout := equalityObjectUnderTest{age: tc.input}
			got := spec.IsValid(eout)
			if got != tc.want {
				t.Errorf("IsMoreOrEqualThan spec expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func Test_IsMoreThan_Specification(t *testing.T) {

	p := func(t equalityObjectUnderTest) int64 {
		return t.age
	}
	spec := NewIsMoreThan[equalityObjectUnderTest, int64](18, predicate[equalityObjectUnderTest, int64](p))

	tests := map[string]struct {
		input int64
		want  bool
	}{
		"younger":  {input: 17, want: false},
		"same age": {input: 18, want: false},
		"older":    {input: 19, want: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			eout := equalityObjectUnderTest{age: tc.input}
			got := spec.IsValid(eout)
			if got != tc.want {
				t.Errorf("IsMoreThan spec expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func Test_IsLessOrEqualThan_Specification(t *testing.T) {

	p := func(t equalityObjectUnderTest) int64 {
		return t.age
	}
	spec := NewIsLessOrEqualThan[equalityObjectUnderTest, int64](18, predicate[equalityObjectUnderTest, int64](p))

	tests := map[string]struct {
		input int64
		want  bool
	}{
		"younger":  {input: 17, want: true},
		"same age": {input: 18, want: true},
		"older":    {input: 19, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			eout := equalityObjectUnderTest{age: tc.input}
			got := spec.IsValid(eout)
			if got != tc.want {
				t.Errorf("IsLessOrEqualThan spec expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

func Test_IsLessThan_Specification(t *testing.T) {

	p := func(t equalityObjectUnderTest) int64 {
		return t.age
	}
	spec := NewIsLessThan[equalityObjectUnderTest, int64](18, predicate[equalityObjectUnderTest, int64](p))

	tests := map[string]struct {
		input int64
		want  bool
	}{
		"younger":  {input: 17, want: true},
		"same age": {input: 18, want: false},
		"older":    {input: 19, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			eout := equalityObjectUnderTest{age: tc.input}
			got := spec.IsValid(eout)
			if got != tc.want {
				t.Errorf("IsLessThan spec expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

type equalityObjectUnderTest struct {
	age int64
}
