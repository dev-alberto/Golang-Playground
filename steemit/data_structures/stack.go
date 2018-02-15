package data_structures

import (
	"github.com/pkg/errors"
)

type Stack struct {
	items []int
	top int
	size int
}

func NewStack(size int) *Stack {
	return &Stack{make([]int, size), -1, size}
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return s.top - 1 == s.size
}

func (s *Stack) Push(item int) error {
	if s.IsFull() {
		return errors.New("Cannot push into a full Stack")
	}
	s.top ++
	s.items[s.top] = item
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Cannot pop empty stack")
	}
	s.top --
	return s.items[s.top + 1], nil
}
