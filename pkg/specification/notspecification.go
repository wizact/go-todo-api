package specification

type Not[T any] struct {
	specification Specification[T]
}

func NewNot[T any](specification Specification[T]) Specification[T] {
	return Not[T]{
		specification: specification,
	}
}

func (n Not[T]) IsValid(o T) bool {
	return !n.specification.IsValid(o)
}
