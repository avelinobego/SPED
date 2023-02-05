package estrategias

type Strategy[T any] interface {
	execute() T
}

type context[T any] struct {
	s Strategy[T]
}

func New[T any](e Strategy[T]) (result *context[T]) {
	result = &context[T]{s: e}
	return
}
