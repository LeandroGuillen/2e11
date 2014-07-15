package engine

import "math/rand"
import "time"

func (g *Game) transpose() {
  N := g.Size
  for n := 0; n < N - 1; n++ {
    for m := n + 1; m < N; m++ {
      g.Board[n][m], g.Board[m][n] = g.Board[m][n], g.Board[n][m]
    }
  }
}

func (g *Game) reflect() {
  N := g.Size
  for n := 0; n < N; n++ {
    for m := 0; m < N / 2; m++ {
      g.Board[n][m], g.Board[n][N - 1 - m] = g.Board[n][N - 1 - m], g.Board[n][m]
    }
  }
}

func getNewPos(size int) (x, y int ) { 
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  x = r.Intn(size)
  y = r.Intn(size)
  return
}