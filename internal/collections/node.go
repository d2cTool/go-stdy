package internal

type Node[T any] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

func NewNode[T any](data T) *Node[T] {
	return &Node[T]{data: data, next: nil, prev: nil}
}
