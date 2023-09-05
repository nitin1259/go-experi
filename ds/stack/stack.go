package stack

import "sync"

type Stack[T any] struct {
	arr  []T
	lock sync.RWMutex
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(in T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.arr = append(s.arr, in)
}

func (s *Stack[T]) Pop() (out T) {
	if s.IsEmpty() {
		return
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	out = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return out
}

func (s *Stack[T]) Top() (out T) {
	if s.IsEmpty() {
		return
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	out = s.arr[len(s.arr)-1]
	return out
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.arr) == 0
}
