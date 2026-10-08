package specification

type Or[T any] struct {
	specifications []Specification[T]
}

func NewOr[T any](specifications ...Specification[T]) Specification[T] {
	return Or[T]{
		specifications: specifications,
	}
}

func (o Or[T]) IsValid(value T) bool {
	for _, specification := range o.specifications {
		if specification.IsValid(value) {
			return true
		}
	}
	return false
}
