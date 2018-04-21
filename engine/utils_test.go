package engine

import "testing"
import "fmt"

func TestReflect(t *testing.T) {
	g := NewGame(4)
	g.ClearBoard()
	g.Board[0][0] = 1
	g.Board[1][3] = 6
	g.Board[2][2] = 5
	g.Board[2][3] = 7
	g.Board[3][1] = 8

	g.reflect()

	e := NewGame(4)
	e.ClearBoard()
	e.Board[0][0] = 0
	e.Board[0][1] = 0
	e.Board[0][2] = 0
	e.Board[0][3] = 1
	e.Board[1][0] = 6
	e.Board[1][1] = 0
	e.Board[1][2] = 0
	e.Board[1][3] = 0
	e.Board[2][0] = 7
	e.Board[2][1] = 5
	e.Board[2][2] = 0
	e.Board[2][3] = 0
	e.Board[3][0] = 0
	e.Board[3][1] = 0
	e.Board[3][2] = 8
	e.Board[3][3] = 0

	if !g.Equals(&e) {
		fmt.Println("Expected:")
		e.PrettyPrint()
		fmt.Println("Got:")
		g.PrettyPrint()
		t.Error("Boards don't match")
	}
}

func TestTranspose(t *testing.T) {
	g := NewGame(4)
	g.ClearBoard()
	g.Board[0][0] = 1
	g.Board[1][3] = 6
	g.Board[2][2] = 5
	g.Board[2][3] = 7
	g.Board[3][1] = 8

	g.transpose()

	e := NewGame(4)
	e.ClearBoard()
	e.Board[0][0] = 1
	e.Board[0][1] = 0
	e.Board[0][2] = 0
	e.Board[0][3] = 0
	e.Board[1][0] = 0
	e.Board[1][1] = 0
	e.Board[1][2] = 0
	e.Board[1][3] = 8
	e.Board[2][0] = 0
	e.Board[2][1] = 0
	e.Board[2][2] = 5
	e.Board[2][3] = 0
	e.Board[3][0] = 0
	e.Board[3][1] = 6
	e.Board[3][2] = 7
	e.Board[3][3] = 0

	if !g.Equals(&e) {
		fmt.Println("Expected:")
		e.PrettyPrint()
		fmt.Println("Got:")
		g.PrettyPrint()
		t.Error("Boards don't match")
	}
}

func TestFreeCells(t *testing.T) {
	a := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 1},
			[]int{1, 1},
		},
	}
	b := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 1},
			[]int{1, 1},
		},
	}
	c := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 1},
			[]int{1, 0},
		},
	}
	if freeCells(a) {
		t.Error("Expected: false, Got: true")
	}
	if !freeCells(b) {
		t.Error("Expected: true, Got: false")
	}
	if !freeCells(c) {
		t.Error("Expected: true, Got: false")
	}
}

func TestHasEqualAdjacent(t *testing.T) {
	a := &Game{
		Score: 0,
		Moves: 0,
		Size:  3,
		Valid: true,
		Board: [][]int{
			[]int{0, 0, 0},
			[]int{1, 1, 0},
			[]int{0, 0, 0},
		},
	}
	b := &Game{
		Score: 0,
		Moves: 0,
		Size:  3,
		Valid: true,
		Board: [][]int{
			[]int{0, 2, 0},
			[]int{1, 0, 1},
			[]int{0, 3, 0},
		},
	}
	c := &Game{
		Score: 0,
		Moves: 0,
		Size:  3,
		Valid: true,
		Board: [][]int{
			[]int{0, 2, 0},
			[]int{3, 5, 1},
			[]int{0, 0, 0},
		},
	}

	if !hasEqualAdjacent(a, 1, 1) {
		t.Error("Expected: false, Got: true")
	}
	if hasEqualAdjacent(b, 1, 1) {
		t.Error("Expected: false, Got: true")
	}
	if hasEqualAdjacent(c, 1, 1) {
		t.Error("Expected: false, Got: true")
	}
}

func TestAreThereAdjacentCells(t *testing.T) {
	a := &Game{
		Score: 0,
		Moves: 0,
		Size:  3,
		Valid: true,
		Board: [][]int{
			[]int{0, 0, 0},
			[]int{1, 1, 0},
			[]int{0, 0, 0},
		},
	}
	b := &Game{
		Score: 0,
		Moves: 0,
		Size:  3,
		Valid: true,
		Board: [][]int{
			[]int{0, 2, 0},
			[]int{1, 0, 1},
			[]int{0, 3, 1},
		},
	}
	c := &Game{
		Score: 0,
		Moves: 0,
		Size:  3,
		Valid: true,
		Board: [][]int{
			[]int{0, 2, 2},
			[]int{3, 5, 1},
			[]int{0, 0, 0},
		},
	}
	d := &Game{
		Score: 0,
		Moves: 0,
		Size:  3,
		Valid: true,
		Board: [][]int{
			[]int{0, 2, 7},
			[]int{3, 5, 1},
			[]int{2, 1, 0},
		},
	}

	if !areThereAdjacentCells(a) {
		t.Error("Expected: true, Got: false")
	}
	if areThereAdjacentCells(b) {
		t.Error("Expected: false, Got: true")
	}
	if !areThereAdjacentCells(c) {
		t.Error("Expected: true, Got: false")
	}
	if areThereAdjacentCells(d) {
		t.Error("Expected: false, Got: true")
	}
}
