package internal

type DLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewDLinkedList[T any]() *DLinkedList[T] {
	return &DLinkedList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (d *DLinkedList[T]) AddBegin(data T) {
	if d.size == 0 {
		d.head = NewNode(data)
		d.tail = d.head
		d.size = 1
	} else {
		node := NewNode(data)
		node.next = d.head
		d.head.prev = node
		d.head = node
		d.size++
	}
}

func (d *DLinkedList[T]) AddEnd(data T) {
	if d.size == 0 {
		d.head = NewNode(data)
		d.tail = d.head
		d.size = 1
	} else {
		node := NewNode(data)
		node.prev = d.head
		d.tail.next = node
		d.tail = node
		d.size++
	}
}

func (d *DLinkedList[T]) RemoveBegin() *Node[T] {
	if d.size == 0 {
		return nil
	}

	node := d.head
	d.size--
	d.head = d.head.next
	d.head.prev = nil

	return node
}

func (d *DLinkedList[T]) RemoveEnd() *Node[T] {
	if d.size == 0 {
		return nil
	}

	node := d.tail
	d.size--
	d.tail = d.tail.prev
	d.tail.next = nil

	return node
}
