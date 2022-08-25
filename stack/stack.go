package stack

import "fmt"

type Stack[T any] struct {
	items []T
}

func New[T any]() Stack[T] {
	return Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Size() < 1 {
		var zeroValue T
		return zeroValue, fmt.Errorf("Pop on empty stack")
	}
	poppedItem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return poppedItem, nil
}

func (s *Stack[T]) Pop2() (T, T, error) {
	if s.Size() < 2 {
		var zeroValue T
		return zeroValue, zeroValue, fmt.Errorf("Pop2 on stack of size less than two")
	}
	poppedItem1 := s.items[len(s.items)-1]
	poppedItem2 := s.items[len(s.items)-2]
	s.items = s.items[:len(s.items)-2]
	return poppedItem1, poppedItem2, nil
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Items() []T {
	return s.items
}
