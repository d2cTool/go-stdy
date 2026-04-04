package internal

import "fmt"

type empty struct{}

type Set[T comparable] struct {
	Items map[T]empty
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		Items: make(map[T]empty),
	}
}

func (s *Set[T]) Add(set *Set[T]) {
	for k := range set.Items {
		s.Items[k] = empty{}
	}
}

func (s *Set[T]) Remove(set *Set[T]) {
	for k := range set.Items {
		delete(s.Items, k)
	}
}

func (s *Set[T]) Intersect(set *Set[T]) *Set[T] {
	newSet := NewSet[T]()
	for k := range set.Items {
		if _, ok := s.Items[k]; ok {
			newSet.Items[k] = empty{}
		}
	}
	return newSet
}

func (s *Set[T]) Format(f fmt.State, c rune) {
	keys := make([]T, 0, len(s.Items))
	for k := range s.Items {
		keys = append(keys, k)
	}
	fmt.Fprintf(f, "%v", keys)
}
