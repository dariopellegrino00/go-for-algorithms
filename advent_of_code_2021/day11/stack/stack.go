package stack

import "fmt"

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		panic("Called peek() on empty stack")
	} else {
		return s.stack[0]
	}
}

func (s *Stack[T]) Pop() T {
	e := s.Peek()
	s.stack = s.stack[1:]
	return e
}

func (s *Stack[T]) Push(e T) {
	s.stack = append(s.stack, e)
}

func (s Stack[T]) String() string {
	str := "["
	for i, e := range s.stack {
		if i < len(s.stack)-1 {
			str += fmt.Sprint(e) + " -> "
		} else {
			str += fmt.Sprint(e)
		}
	}
	return str + "]"
}
