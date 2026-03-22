package internal

type Stack struct {
	elements []interface{}
	top      int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
		top:      -1,
	}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
	s.top++
}

func (s *Stack) Pop() interface{} {
	if s.top < 0 {
		return nil
	}

	element := s.elements[s.top]
	s.elements = s.elements[:s.top]
	s.top--
	return element
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}
