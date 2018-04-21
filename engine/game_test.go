package engine

import (
	"fmt"
	"testing"
)

func TestEquals(t *testing.T) {
	game1 := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 0},
			[]int{0, 0},
		},
	}

	game2 := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 0},
			[]int{0, 0},
		},
	}

	if !game1.Equals(game2) {
		t.Error("copied game does not equal source")
		fmt.Println("Expected:")
		game1.PrettyPrint()
		fmt.Println("Got:")
		game2.PrettyPrint()
	}
}
func TestCopyGame(t *testing.T) {
	aGame := Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 0},
			[]int{0, 0},
		},
	}
	copy1 := aGame.Copy()

	if !copy1.Equals(&aGame) {
		t.Error("copied game does not equal source")
		fmt.Println("Expected:")
		aGame.PrettyPrint()
		fmt.Println("Got:")
		copy1.PrettyPrint()
	}

	// modify aGame to check that copy1 and copy2 are still
	// the same and different than aGame now
	copy2 := aGame.Copy()
	aGame.GoRight()

	if !copy1.Equals(&copy2) || aGame.Equals(&copy2) || aGame.Equals(&copy1) {
		t.Error("copied game does not equal source")
	}
}

func TestClearBoard(t *testing.T) {
	game1 := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{1, 0},
			[]int{0, 1},
		},
	}

	zeroedBoardGame := &Game{
		Score: 0,
		Moves: 0,
		Size:  2,
		Valid: true,
		Board: [][]int{
			[]int{0, 0},
			[]int{0, 0},
		},
	}

	game1.ClearBoard()

	if !game1.Equals(zeroedBoardGame) {
		t.Error("copied game does not equal source")
		fmt.Println("Expected:")
		zeroedBoardGame.PrettyPrint()
		fmt.Println("Got:")
		game1.PrettyPrint()
	}
}
