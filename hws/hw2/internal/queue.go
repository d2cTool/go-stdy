package internal

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		elements: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(v interface{}) {
	q.elements = append(q.elements, v)
}

func (q *Queue) Dequeue() interface{} {
	if len(q.elements) == 0 {
		return nil
	}

	v := q.elements[0]
	q.elements = q.elements[1:]
	return v
}
