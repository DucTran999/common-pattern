package stack

import "errors"

type Stack interface {
	Push(val int)
	Pop() (int, error)
	Len() int
}

type stack struct {
	arr []int
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(val int) {
	s.arr = append(s.arr, val)
}

func (s *stack) Pop() (int, error) {
	if len(s.arr) == 0 {
		return 0, errors.New("stack is empty")
	}

	top := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]

	return top, nil
}

func (s *stack) Len() int {
	return len(s.arr)
}
