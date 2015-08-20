package aiplayer

// Author:
// https://gist.github.com/bemasher/1777766

import "github.com/leandroguillen/2e11/engine"

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value engine.Game
	next  *Element
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value engine.Game) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value engine.Game) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return engine.NullGame()
}

// func main() {
// 	stack := new(Stack)
//
// 	stack.Push("Things")
// 	stack.Push("and")
// 	stack.Push("Stuff")
//
// 	for stack.Len() > 0 {
// 		fmt.Printf("%s ", stack.Pop().(string))
// 	}
// 	fmt.Println()
// }
