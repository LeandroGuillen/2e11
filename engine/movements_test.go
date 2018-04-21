package engine

import (
	"fmt"
	"testing"
)

func TestGoRight(t *testing.T) {
	a := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 0},
			[]int{0, 0},
		},
	}
	b := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 1},
			[]int{0, 0},
		},
	}

	if ok, err := a.GoRight(); !ok || !a.Equals(b) {
		fmt.Println("Expected:")
		a.PrintBoard()
		fmt.Println("Got:")
		b.PrintBoard()
		t.Error(err)
	}
}

func TestGoDown(t *testing.T) {
	a := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 0},
			[]int{0, 0},
		},
	}
	b := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 0},
			[]int{1, 0},
		},
	}

	if ok, err := a.GoDown(); !ok || !a.Equals(b) {
		fmt.Println("Expected:")
		a.PrintBoard()
		fmt.Println("Got:")
		b.PrintBoard()
		t.Error(err)
	}
}

func TestGoLeft(t *testing.T) {
	a := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 1},
			[]int{0, 0},
		},
	}
	b := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 0},
			[]int{0, 0},
		},
	}

	if ok, err := a.GoLeft(); !ok || !a.Equals(b) {
		fmt.Println("Expected:")
		a.PrintBoard()
		fmt.Println("Got:")
		b.PrintBoard()
		t.Error(err)
	}
}

func TestGoUp(t *testing.T) {
	a := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 0},
			[]int{0, 1},
		},
	}
	b := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 1},
			[]int{0, 0},
		},
	}

	if ok, err := a.GoUp(); !ok || !a.Equals(b) {
		fmt.Println("Expected:")
		a.PrintBoard()
		fmt.Println("Got:")
		b.PrintBoard()
		t.Error(err)
	}
}

// func TestGoRightNoCollapse(t *testing.T) {
// 	g := NewGame(4)
// 	g.ClearBoard()
// 	g.Board[0][0] = 1
// 	g.Board[0][1] = 2
// 	g.Board[0][2] = 3
// 	g.Board[1][0] = 4
// 	g.Board[1][3] = 6
// 	g.Board[2][2] = 5
// 	g.Board[2][3] = 7
// 	g.Board[3][1] = 8

// 	e := NewGame(4)
// 	e.ClearBoard()
// 	e.Board[0][0] = 0
// 	e.Board[0][1] = 1
// 	e.Board[0][2] = 2
// 	e.Board[0][3] = 3
// 	e.Board[1][0] = 0
// 	e.Board[1][1] = 0
// 	e.Board[1][2] = 4
// 	e.Board[1][3] = 6
// 	e.Board[2][0] = 0
// 	e.Board[2][1] = 0
// 	e.Board[2][2] = 5
// 	e.Board[2][3] = 7
// 	e.Board[3][0] = 0
// 	e.Board[3][1] = 0
// 	e.Board[3][2] = 0
// 	e.Board[3][3] = 8

// 	g.GoRight()

// 	if !g.Equals(&e) {
// 		fmt.Println("Expected:")
// 		e.PrettyPrint()
// 		fmt.Println("Got:")
// 		g.PrettyPrint()
// 		t.Error("Boards don't match")
// 	}
// }

// func TestGoRight(t *testing.T) {
// 	g := NewGame(4)
// 	g.ClearBoard()
// 	g.Board[0][0] = 1
// 	g.Board[0][1] = 1
// 	g.Board[0][2] = 3
// 	g.Board[1][0] = 4
// 	g.Board[1][3] = 6
// 	g.Board[2][2] = 7
// 	g.Board[2][3] = 7
// 	g.Board[3][1] = 8

// 	e := NewGame(4)
// 	e.ClearBoard()
// 	e.Board[0][0] = 0
// 	e.Board[0][1] = 0
// 	e.Board[0][2] = 2
// 	e.Board[0][3] = 3
// 	e.Board[1][0] = 0
// 	e.Board[1][1] = 0
// 	e.Board[1][2] = 4
// 	e.Board[1][3] = 6
// 	e.Board[2][0] = 0
// 	e.Board[2][1] = 0
// 	e.Board[2][2] = 0
// 	e.Board[2][3] = 14
// 	e.Board[3][0] = 0
// 	e.Board[3][1] = 0
// 	e.Board[3][2] = 0
// 	e.Board[3][3] = 8

// 	g.GoRight()

// 	if !g.Equals(&e) {
// 		fmt.Println("Expected:")
// 		e.PrettyPrint()
// 		fmt.Println("Got:")
// 		g.PrettyPrint()
// 		t.Error("Boards don't match")
// 	}
// }
