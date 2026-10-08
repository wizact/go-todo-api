package specification

type predicate[T any, V Number] func(T) V

type Number interface {
	int64 | float64
}

type GreaterOrEqual[T any, V Number] struct {
	minValue  V
	predicate predicate[T, V]
}

func NewGreaterOrEqual[T any, V Number](minValue V, pr predicate[T, V]) Specification[T] {
	return GreaterOrEqual[T, V]{
		minValue:  minValue,
		predicate: pr,
	}
}

func (g GreaterOrEqual[T, V]) IsValid(o T) bool {
	return g.predicate(o) >= g.minValue
}

type GreaterThan[T any, V Number] struct {
	minValue  V
	predicate predicate[T, V]
}

func NewGreaterThan[T any, V Number](minValue V, pr predicate[T, V]) Specification[T] {
	return GreaterThan[T, V]{
		minValue:  minValue,
		predicate: pr,
	}
}

func (g GreaterThan[T, V]) IsValid(o T) bool {
	return g.predicate(o) > g.minValue
}

type LessOrEqual[T any, V Number] struct {
	maxValue  V
	predicate predicate[T, V]
}

func NewLessOrEqual[T any, V Number](maxValue V, pr predicate[T, V]) Specification[T] {
	return LessOrEqual[T, V]{
		maxValue:  maxValue,
		predicate: pr,
	}
}

func (l LessOrEqual[T, V]) IsValid(o T) bool {
	return l.predicate(o) <= l.maxValue
}

type LessThan[T any, V Number] struct {
	maxValue  V
	predicate predicate[T, V]
}

func NewLessThan[T any, V Number](maxValue V, pr predicate[T, V]) Specification[T] {
	return LessThan[T, V]{
		maxValue:  maxValue,
		predicate: pr,
	}
}

func (l LessThan[T, V]) IsValid(o T) bool {
	return l.predicate(o) < l.maxValue
}
