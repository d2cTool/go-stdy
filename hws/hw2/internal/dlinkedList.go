package internal

type DLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func NewDLinkedList() *DLinkedList {
	return &DLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// нельзя делать функции с одним названием и разным набором параметров
// нельзя использовать дефолтные значения для аргументов
func NewNode(data interface{}, n *Node, p *Node) *Node {
	return &Node{data: data, next: n, prev: p}
}

func (d *DLinkedList) AddBegin(data interface{}) {
	if d.size == 0 {
		d.head = NewNode(data, nil, nil)
		d.tail = d.head
		d.size = 1
	} else {
		node := NewNode(data, d.head, nil)
		d.head.prev = node
		d.head = node
		d.size++
	}
}

func (d *DLinkedList) AddEnd(data interface{}) {
	if d.size == 0 {
		d.head = NewNode(data, nil, nil)
		d.tail = d.head
		d.size = 1
	} else {
		node := NewNode(data, nil, d.head)
		d.tail.next = node
		d.tail = node
		d.size++
	}
}

func (d *DLinkedList) RemoveBegin(data interface{}) *Node {
	if d.size == 0 {
		return nil
	}

	node := d.head
	d.size--
	d.head = d.head.next
	d.head.prev = nil

	return node
}

func (d *DLinkedList) RemoveEnd(data interface{}) *Node {
	if d.size == 0 {
		return nil
	}

	node := d.tail
	d.size--
	d.tail = d.tail.prev
	d.tail.next = nil

	return node
}
