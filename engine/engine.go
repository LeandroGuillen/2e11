package engine

import "math/rand"
import "time"
import "errors"

// NewGame returns a new initialized game
func NewGame(size int) Game {

	b := make([][]int, size)
	for i := 0; i < size; i++ {
		b[i] = make([]int, size)
	}

	g := Game{Board: b, Score: 0, Moves: 0, Size: size, Valid: true}
	g.ClearBoard()
	// Start off with a couple of cells filled
	PlaceNewValue(&g)
	PlaceNewValue(&g)
	return g
}

func getNewValue() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	chance := r.Intn(10)
	value := 0
	if chance <= 1 {
		value = 4
	} else {
		value = 2
	}
	return value
}

func PlaceNewValue(g *Game) error {
	value := getNewValue()
	x, y := getNewPos(g.Size)

	freeCells := freeCells(g)
	// adjCells := twoAdjacentsCells(*g)

	if !freeCells {
		// if !adjCells {
		// 	return errors.New("can't place another value, the game is over")
		// }
		return errors.New("try again")
	}

	if g.GetCell(x, y) != 0 {
		for x, y = getNewPos(g.Size); g.GetCell(x, y) != 0; x, y = getNewPos(g.Size) {
		}
	} else {
		g.SetCell(x, y, value)
	}

	return nil
}

// NullGame returns a non-valid game
func NullGame() (g Game) {
	g = Game{Valid: false}
	return
}
