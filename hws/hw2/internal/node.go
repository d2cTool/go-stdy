package internal

type Node[T any] struct {
	data *T
	next *Node[T]
	prev *Node[T]
}

type Option[T any] func(*Node[T])

func NewNode[T any](data *T, n *Node[T], p *Node[T]) *Node[T] {
	return &Node[T]{data: data, next: n, prev: p}
}

func NewNodeWithOptions[T any](data *T, opts ...Option[T]) *Node[T] {
	node := NewNode(data, nil, nil)
	for _, opt := range opts {
		opt(node)
	}
	return node
}
