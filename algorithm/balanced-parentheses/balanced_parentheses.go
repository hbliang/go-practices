package balancedparentheses

import (
	"sync"
)

type Stack struct {
	mutex sync.Mutex
	top   *StackNode
}

type StackNode struct {
	next  *StackNode
	value string
}

func NewStack() *Stack {
	return &Stack{
		top: nil,
	}
}

func (s *Stack) Push(v string) {
	s.mutex.Lock()
	n := &StackNode{s.top, v}
	s.top = n
	s.mutex.Unlock()
}

func (s *Stack) Pop() string {
	r := ""
	s.mutex.Lock()
	if s.top != nil {
		r = s.top.value
		s.top = s.top.next
	}
	s.mutex.Unlock()

	return r
}

func IsParanthesisBalanced(s string) bool {
	stack := NewStack()
	for _, r := range s {
		v := string(r)
		if v == "(" || v == "[" || v == "{" {
			stack.Push(v)
		} else if v == ")" || v == "]" || v == "}" {
			pop := stack.Pop()
			if !match(pop, v) {
				return false
			}
		}
	}

	return true
}

func match(a string, b string) bool {

	if a == "(" && b == ")" {
		return true
	}

	if a == "[" && b == "]" {
		return true
	}

	if a == "{" && b == "}" {
		return true
	}

	return false
}
