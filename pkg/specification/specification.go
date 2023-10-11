package specification

type Specification[T any] interface {
	IsValid(o T) bool
}

type FunctionSpecification[T any] func(o T) bool

func (fs FunctionSpecification[T]) IsValid(o T) bool {
	return fs(o)
}
