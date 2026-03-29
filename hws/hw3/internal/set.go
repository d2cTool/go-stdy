package internal

type empty struct{}

type Set[T comparable] struct {
	items map[T]empty
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]empty),
	}
}

func (s *Set[T]) Add(set *Set[T]) {
	for k := range set.items {
		s.items[k] = empty{}
	}
}

func (s *Set[T]) Remove(set *Set[T]) {
	for k := range set.items {
		delete(s.items, k)
	}
}

func (s *Set[T]) Intersect(set *Set[T]) *Set[T] {
	newSet := NewSet[T]()
	for k := range set.items {
		if _, ok := s.items[k]; ok {
			newSet.items[k] = empty{}
		}
	}
	return newSet
}
