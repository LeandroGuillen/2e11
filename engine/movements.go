package engine

import "errors"

// GoUp executes a game movement to the appropriate direction
func (g *Game) GoUp() (bool, error) {
	backup := Game{}
	backup.Board = g.Board

	g.transpose()
	g.reflect()
	for _, val := range g.Board {
		g.shift(val)
	}
	g.reflect()
	g.transpose()

	if backup.Equals(g) {
		return false, errors.New("movement not possible")
	}
	// return placeNewValue(g)
	return true, nil
}

// GoDown executes a game movement to the appropriate direction
func (g *Game) GoDown() (bool, error) {
	backup := Game{}
	backup.Board = g.Board

	g.transpose()
	for _, val := range g.Board {
		g.shift(val)
	}
	g.transpose()
	if backup.Equals(g) {
		return false, errors.New("movement not possible")
	}
	// return placeNewValue(g)
	return true, nil
}

// GoLeft executes a game movement to the appropriate direction
func (g *Game) GoLeft() (bool, error) {
	backup := g.Copy()

	g.reflect()
	for _, val := range g.Board {
		g.shift(val)
	}
	g.reflect()

	if backup.Equals(g) {
		return false, errors.New("movement not possible")
	}
	// return placeNewValue(g)
	return true, nil
}

// GoRight executes a game movement to the appropriate direction
func (g *Game) GoRight() (bool, error) {
	backup := Game{}
	backup.Board = g.Board

	for _, val := range g.Board {
		g.shift(val)
	}

	if backup.Equals(g) {
		return false, errors.New("movement not possible")
	}
	// return placeNewValue(g)
	return true, nil
}

func (g *Game) shift(array []int) {
	//for n := 0; n < g.Size - 1; n++ {

	temp := make([]int, g.Size)
	tempIndex := g.Size - 1

	// Push values to the right
	for i := g.Size - 1; i >= 0; i-- {
		if array[i] != 0 {
			temp[tempIndex] = array[i]
			tempIndex--
		}
	}

	// Update array
	for i := 0; i < g.Size; i++ {
		array[i] = temp[i]
	}

	needToPushAgain := false

	// Merge double occurrences
	for i := g.Size - 1; i > 0; i-- {
		if array[i] == array[i-1] {
			// Clear the collapsing cell
			array[i-1] = 0
			// Increase cell value
			array[i] = array[i] * 2
			// Update score and moves
			g.Score = g.Score + array[i]
			g.Moves = g.Moves + 1
			i--
			needToPushAgain = true
		}
	}

	// We may need to push right again
	if needToPushAgain {
		tempIndex = g.Size - 1

		// Reset temp
		for i := 0; i < g.Size; i++ {
			temp[i] = 0
		}

		// Push right
		for i := g.Size - 1; i >= 0; i-- {
			if array[i] != 0 {
				temp[tempIndex] = array[i]
				tempIndex--
			}
		}

		// Update array
		for i := 0; i < g.Size; i++ {
			array[i] = temp[i]
		}
	}
}
