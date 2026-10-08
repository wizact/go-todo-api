package specification

type Specification[T any] interface {
	IsValid(o T) bool
}

type Func[T any] func(o T) bool

func (f Func[T]) IsValid(o T) bool {
	return f(o)
}
