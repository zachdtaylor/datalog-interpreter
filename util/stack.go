package util

type Stack struct {
	stack []int
}

func (s *Stack) Push(value int) {
	s.stack = append(s.stack, value)
}

func (s *Stack) Pop() int {
	l := len(s.stack)
	value := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return value
}

func (s *Stack) Values() []int {
	return s.stack
}
