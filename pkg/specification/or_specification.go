package specification

type OrSpecification[T any] struct {
	specifications []Specification[T]
}

func NewOrSpecification[T any](specifications ...Specification[T]) Specification[T] {
	return OrSpecification[T]{
		specifications: specifications,
	}
}

func (s OrSpecification[T]) IsValid(o T) bool {
	for _, specification := range s.specifications {
		if specification.IsValid(o) {
			return true
		}
	}
	return false
}
