package stack

import (
	"fmt"
)

var (
	ErrStackIsNil = fmt.Errorf("stack is nil\n")
)

type Stack[T any] struct {
	values []T
}

// Empty if stack is empty return true, return false
func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

// Len get length of stack
func (s *Stack[T]) Len() int {
	return len(s.values)
}

// Clear clear stack
func (s *Stack[T]) Clear() {
	s.values = nil
}

// Top get top value of stack
func (s *Stack[T]) Top() T {
	if s.Len() > 0 {
		return s.values[s.Len()-1]
	}
	panic(ErrStackIsNil)
}

// Pop pop
func (s *Stack[T]) Pop() error {
	if s.Len() > 0 {
		s.values = s.values[:s.Len()-1]
		return nil
	}
	return ErrStackIsNil
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}
