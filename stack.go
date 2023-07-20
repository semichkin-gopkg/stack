package stack

type Stack[T any] struct {
	values []T
	length int
}

func New[T any](cap int) *Stack[T] {
	return &Stack[T]{make([]T, 0, cap), 0}
}

func (s *Stack[T]) Push(key T) {
	s.values = append(s.values, key)
	s.length++
}

func (s *Stack[T]) Top() (T, bool) {
	var x T
	if s.length <= 0 {
		return x, false
	}

	return s.values[s.length-1], true
}

func (s *Stack[T]) Pop() (T, bool) {
	var x T
	if s.length <= 0 {
		return x, false
	}

	x, s.values = s.values[s.length-1], s.values[:s.length-1]
	s.length--
	return x, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}
