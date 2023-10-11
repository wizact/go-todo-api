package specification

type predicate[T any, V Number] func(T) V

type Number interface {
	int64 | float64
}

type IsMoreOrEqualThan[T any, V Number] struct {
	minValue  V
	predicate predicate[T, V]
}

func NewIsMoreOrEqualThan[T any, V Number](minValue V, pr predicate[T, V]) Specification[T] {
	return IsMoreOrEqualThan[T, V]{
		minValue:  minValue,
		predicate: pr,
	}
}

func (h IsMoreOrEqualThan[T, V]) IsValid(o T) bool {
	return h.predicate(o) >= h.minValue
}

type IsMoreThan[T any, V Number] struct {
	minValue  V
	predicate predicate[T, V]
}

func NewIsMoreThan[T any, V Number](minValue V, pr predicate[T, V]) Specification[T] {
	return IsMoreThan[T, V]{
		minValue:  minValue,
		predicate: pr,
	}
}

func (h IsMoreThan[T, V]) IsValid(o T) bool {
	return h.predicate(o) > h.minValue
}

type IsLessOrEqualThan[T any, V Number] struct {
	maxValue  V
	predicate predicate[T, V]
}

func NewIsLessOrEqualThan[T any, V Number](maxValue V, pr predicate[T, V]) Specification[T] {
	return IsLessOrEqualThan[T, V]{
		maxValue:  maxValue,
		predicate: pr,
	}
}

func (h IsLessOrEqualThan[T, V]) IsValid(o T) bool {
	return h.predicate(o) <= h.maxValue
}

type IsLessThan[T any, V Number] struct {
	maxValue  V
	predicate predicate[T, V]
}

func NewIsLessThan[T any, V Number](maxValue V, pr predicate[T, V]) Specification[T] {
	return IsLessThan[T, V]{
		maxValue:  maxValue,
		predicate: pr,
	}
}

func (h IsLessThan[T, V]) IsValid(o T) bool {
	return h.predicate(o) < h.maxValue
}
