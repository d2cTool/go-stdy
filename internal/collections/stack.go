package internal

import "errors"

var (
	ErrStackEmpty = errors.New("stack is empty")
)

type Stack[T any] struct {
	node *Node[T]
	size int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		node: nil,
		size: 0,
	}
}

func (s *Stack[T]) Push(data T) {
	item := NewNode(data)
	item.next = s.node
	if s.node != nil {
		s.node.prev = item
	}
	s.node = item
	s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size == 0 {
		var zero T
		return zero, ErrStackEmpty
	}
	data := s.node.data
	s.node = s.node.next
	if s.node != nil {
		s.node.prev = nil
	}
	s.size--
	return data, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}
