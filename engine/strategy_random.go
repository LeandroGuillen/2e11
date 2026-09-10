package engine

import "math/rand/v2"

type Random struct {
}

func (s Random) GetNextMove(g *Game) int {
	return rand.IntN(4)
}

func (s Random) Name() string {
	return "Random"
}
