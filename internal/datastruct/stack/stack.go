package stack

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(v int) {
	s.head = &Node{
		Value: v,
		Next:  s.head,
	}
}

func (s *Stack) Pop() (int, bool) {
	if s.head == nil {
		return 0, false
	}

	v := s.head.Value
	s.head = s.head.Next

	return v, true
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
