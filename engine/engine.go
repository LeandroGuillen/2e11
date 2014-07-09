package game

import "fmt"

type Game struct {
  board [][]int
  score int
  moves int
  size int
}

func NewGame(size int) Game {
  b := make([][]int, size)
  for i := 0; i < size; i++ {
      b[i] = make([]int, size)
  }
  
  g := Game{board: b, score: 0, moves: 0, size: size}
  
  return g
}

func (g *Game) PrettyPrint() {
  fmt.Println("Board:")
  for i := 0; i < g.size; i++ {
      
      for j := 0; j < g.size; j++ {
          fmt.Print(g.board[i][j], " ")
      }
      fmt.Println()
  } 
  fmt.Println("Score: ", g.score, " Moves: ", g.moves)
}