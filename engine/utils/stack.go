package utils

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

func (s *Stack[T]) Pop() (*T, bool) {
	if len(s.data) == 0 {
		return nil, false
	}
	data := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return &data, true
}

func (s *Stack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Peek() (*T, bool) {
	if len(s.data) == 0 {
		return nil, false
	}
	return &s.data[len(s.data)-1], true
}
