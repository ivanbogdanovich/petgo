package stack

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	i := len(s.items) - 1
	v := s.items[i]
	s.items = s.items[:i]
	return v, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func CallStack() {
	var s Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for !s.IsEmpty() {
		v, _ := s.Pop()
		fmt.Println(v)
	}
}