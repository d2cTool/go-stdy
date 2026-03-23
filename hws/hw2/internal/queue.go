package internal

import "errors"

var (
	ErrQueueEmpty = errors.New("queue is empty")
)

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue[T]) Enqueue(data *T) {
	item := NewNode(data, nil, nil)
	if q.size == 0 {
		q.head = item
		q.tail = item
	} else {
		q.tail.prev = item
		item.next = q.tail
		q.tail = item
	}

	q.size++
}

func (q *Queue[T]) Dequeue() (*T, error) {
	if q.size == 0 {
		return nil, ErrQueueEmpty
	}

	data := q.head.data
	q.head = q.head.prev
	if q.head != nil {
		q.head.next = nil
	} else {
		q.tail = nil
	}
	q.size--
	return data, nil
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
