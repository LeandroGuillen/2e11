package strategy

import "math/rand"
import "time"
import "github.com/LeandroGuillen/2e11/engine"

type Strategy interface {
    GetNextMove(*engine.Game) int
}

type Random struct {
}

func (s Random) GetNextMove(g *engine.Game) int {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  return r.Intn(4)
}