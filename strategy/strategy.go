package strategy

import "math/rand"
import "time"

type Strategy interface {
    GetNextMove([][]int) int
}

type Random struct {
}

func (s Random) GetNextMove(board [][]int) int {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  return r.Intn(4)
}