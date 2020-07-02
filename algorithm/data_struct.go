package algorithm

import "errors"

// Stack stack operate
type Stack []interface{}

// Len  stack length
func (s Stack) Len() int {
	return len(s)
}

// Cap stack cap
func (s Stack) Cap() int {
	return cap(s)
}

// IsEmpty is stack empty?
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Push add data to stack
func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

// Top get stack top data
func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s[len(s)-1], nil
}

// Pop pop stack data
func (s *Stack) Pop() (interface{}, error) {
	stk := *s
	if len(stk) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := stk[len(stk)-1]
	*s = stk[:len(stk)-1]
	return top, nil
}
