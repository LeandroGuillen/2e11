package engine

import "math/rand"
import "time"

func getNewPos(size int) (x, y int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	x = r.Intn(size)
	y = r.Intn(size)
	return
}

// areThereAdjacentCells returns true if there is at least
// one cell whose value has an equal adjacent in the same line
func areThereAdjacentCells(g *Game) bool {
	for i := 0; i < g.Size-1; i++ {
		for j := 0; j < g.Size-1; j++ {
			if hasEqualAdjacent(g, i, j) {
				return true
			}
		}
	}
	return false
}

// twoAdjacentsCells returns true if there is at least one
// cell at north, south, east and west of any other that
// equal each other
func twoAdjacentsCells(g *Game) bool {
	if g.Size < 4 {
		// TODO implement case where board is smaller
		return false
	}

	for i := 1; i < g.Size-1; i++ {
		for j := 1; j < g.Size-1; j++ {
			a := g.Board[i][j]
			b := g.Board[i][j+1]
			c := g.Board[i][j-1]
			d := g.Board[i+1][j]
			e := g.Board[i-1][j]

			if a == b || a == c || a == d || a == e {
				return true
			}
		}
	}

	// border cases
	a := g.Board[0][0]
	b := g.Board[1][0]
	c := g.Board[2][0]
	d := g.Board[3][0]

	if b == a || b == c || c == d {
		return true
	}

	e := g.Board[0][3]
	f := g.Board[1][3]
	x := g.Board[2][3]
	y := g.Board[3][3]

	if f == e || f == x || x == y {
		return true
	}

	a = g.Board[0][0]
	b = g.Board[0][1]
	c = g.Board[0][2]
	d = g.Board[0][3]

	if b == a || b == c || c == d {
		return true
	}

	e = g.Board[3][0]
	f = g.Board[3][1]
	x = g.Board[3][2]
	y = g.Board[3][3]

	if f == e || f == x || x == y {
		return true
	}

	return false
}

// freeCells returns true if there is at least one cell
// where a new value can be placed (cell value equals 0)
func freeCells(g *Game) bool {
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			if g.Board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

// hasEqualAdjacent returns true if, given a value, there is at least
// one other cell either north, south, east or west whose value
// is the same.
func hasEqualAdjacent(g *Game, x, y int) bool {

	v := g.Board[x][y]

	if x == 0 {
		if y == 0 {
			return v == g.Board[x][y+1] || v == g.Board[x+1][y]
		} else if y == g.Size-1 {
			return v == g.Board[x-1][y] || v == g.Board[x][y+1]
		} else {
			return v == g.Board[x][y+1] || v == g.Board[x][y-1] || v == g.Board[x+1][y]
		}
	} else if x == g.Size-1 {
		if y == 0 {
			return v == g.Board[x][y+1] || v == g.Board[x-1][y]
		} else if y == g.Size-1 {
			return v == g.Board[x-1][y] || v == g.Board[x][y-1]
		} else {
			return v == g.Board[x][y+1] || v == g.Board[x][y-1] || v == g.Board[x-1][y]
		}
	} else {

		if y == 0 {
			return v == g.Board[x][y+1] || v == g.Board[x-1][y] || v == g.Board[x+1][y]
		} else if y == g.Size-1 {
			return v == g.Board[x-1][y] || v == g.Board[x][y-1] || v == g.Board[x+1][y]
		} else {
			return v == g.Board[x][y+1] || v == g.Board[x][y-1] || v == g.Board[x-1][y] || v == g.Board[x+1][y]
		}
	}
}
