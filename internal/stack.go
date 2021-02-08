package internal

import "sync"

// Stack is a basic Stack data structure implementation
type Stack struct {
	elements []Package
	sync.RWMutex
}

// Push inserts an element in the stack
func (s *Stack) Push(element Package) {
	s.Lock()
	s.elements = append(s.elements, element)
	s.Unlock()
}

// Pop remove the last element from the stack and returns it
func (s *Stack) Pop() Package {
	s.Lock()
	defer s.Unlock()
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return element
}

// IsEmpty returns true when the stack is empty
func (s *Stack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.elements) == 0
}
