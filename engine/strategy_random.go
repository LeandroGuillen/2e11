package engine

import (
	"math/rand"
	"time"
)

type Random struct {
}

func (s Random) GetNextMove(g *Game) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(4)
}

func (s Random) Name() string {
	return "Random"
}
