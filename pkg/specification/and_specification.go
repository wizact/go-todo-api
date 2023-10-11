package specification

type AndSpecification[T any] struct {
	specifications []Specification[T]
}

func NewAndSpecification[T any](specifications ...Specification[T]) Specification[T] {
	return AndSpecification[T]{
		specifications: specifications,
	}
}

func (s AndSpecification[T]) IsValid(o T) bool {
	for _, specification := range s.specifications {
		if !specification.IsValid(o) {
			return false
		}
	}
	return true
}
