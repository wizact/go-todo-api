package specification

type And[T any] struct {
	specifications []Specification[T]
}

func NewAnd[T any](specifications ...Specification[T]) Specification[T] {
	return And[T]{
		specifications: specifications,
	}
}

func (a And[T]) IsValid(o T) bool {
	for _, specification := range a.specifications {
		if !specification.IsValid(o) {
			return false
		}
	}
	return true
}
