package engine

import "fmt"

// Game type
type Game struct {
	Board [][]int
	Score int
	Moves int
	Size  int
	Valid bool
}

// Equals returns true if two games are equal. This means that
// every cell on the board has the same value, regardless of
// score or movements.
func (g *Game) Equals(h *Game) bool {
	if g.Size != h.Size {
		return false
	}
	out := true
	g.traverse(func(i, j int) {
		if g.Board[i][j] != h.Board[i][j] {
			out = false
		}
	})
	return out
}

// Copy returns a copy of a given game. Ignores score and movements.
func (g Game) Copy() Game {
	n := g
	n.Board = make([][]int, g.Size)
	for i := 0; i < g.Size; i++ {
		n.Board[i] = make([]int, g.Size)
	}
	g.traverse(func(i, j int) {
		n.Board[i][j] = g.Board[i][j]
	})
	return n
}

// PrettyPrint prints game board on standard output
func (g Game) PrettyPrint() {
	fmt.Printf("Score: %4d  Moves: %4d  Size: %4d  Valid %t\n", g.Score, g.Moves, g.Size, g.Valid)
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			fmt.Printf("%5d", g.Board[i][j])
		}
		fmt.Println()
	}
}

// PrintBoard prints only the board
func (g Game) PrintBoard() {
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			fmt.Printf("%3d", g.Board[i][j])
		}
		fmt.Println()
	}
}

// ClearBoard sets all postions to zero
func (g *Game) ClearBoard() {
	g.traverse(func(i, j int) {
		g.Board[i][j] = 0
	})
}

// IsGameOver returns true if no more movements are possible
func (g *Game) IsGameOver() bool {
	// Any free cells?
	return !freeCells(g) && !twoAdjacentsCells(g)
}

func (g *Game) traverse(fn func(int, int)) {
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			fn(i, j)
		}
	}
}

// GetCell obtains the value of a cell
func (g *Game) GetCell(i, j int) int {
	return g.Board[i][j]
}

// SetCell changes the value of a cell
func (g *Game) SetCell(i, j, value int) {
	g.Board[i][j] = value
}

func (g *Game) reflect() {
	N := g.Size
	for n := 0; n < N; n++ {
		for m := 0; m < N/2; m++ {
			g.Board[n][m], g.Board[n][N-1-m] = g.Board[n][N-1-m], g.Board[n][m]
		}
	}
}

func (g *Game) transpose() {
	N := g.Size
	for n := 0; n < N-1; n++ {
		for m := n + 1; m < N; m++ {
			g.Board[n][m], g.Board[m][n] = g.Board[m][n], g.Board[n][m]
		}
	}
}
