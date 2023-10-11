package specification

type NotSpecification[T any] struct {
	specification Specification[T]
}

func NewNotSpecification[T any](specification Specification[T]) Specification[T] {
	return NotSpecification[T]{
		specification: specification,
	}
}

func (s NotSpecification[T]) IsValid(o T) bool {
	return !s.specification.IsValid(o)
}
