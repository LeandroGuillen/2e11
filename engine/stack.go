package engine

// Author:
// https://gist.github.com/bemasher/1777766

// Stack ...
type Stack struct {
	top  *Element
	size int
}

// Element ...
type Element struct {
	value Game
	next  *Element
}

// Len returns the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value Game) {
	s.top = &Element{value, s.top}
	s.size++
}

// Pop removes the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value Game) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return NullGame()
}
