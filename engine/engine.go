package engine

import "fmt"

type Game struct {
  Board [][]int
  Score int
  Moves int
  Size int
}

type direction int

const (
  _  = iota
  Up direction = 1
  Down = 2
  Left = 3
  Right = 4
)

func NewGame(Size int) Game {
  b := make([][]int, Size)
  for i := 0; i < Size; i++ {
    b[i] = make([]int, Size)   
  }
  
  g := Game{Board: b, Score: 0, Moves: 0, Size: Size}
  return g
}

func (g *Game) GoUp() {
  g.transpose()
//   fmt.Println("Transposed!")
//   g.PrettyPrint()
  for _, val := range g.Board {
    g.shiftLeft(val)
  }
//   fmt.Println("Upped!")
  g.transpose()
//   fmt.Println("Transposed!")
}

func (g *Game) GoDown() {
  g.transpose()
//   fmt.Println("Transposed!")
//   g.PrettyPrint()
  for _, val := range g.Board {
    g.shiftRight(val)
  }
//   fmt.Println("Downed!")
//   g.PrettyPrint()
  g.transpose()
//   fmt.Println("Transposed!")
}

func (g *Game) GoLeft() {
  for _, val := range g.Board {
    g.shiftLeft(val)
  }
}

func (g *Game) GoRight() {
  for _, val := range g.Board {
    g.shiftRight(val)
  }
}

func (g *Game) transpose() {
  N := g.Size
  for n := 0; n < N - 1; n++ {
    for m := n + 1; m < N; m++ {
      g.Board[n][m], g.Board[m][n] = g.Board[m][n], g.Board[n][m]
    }
  }
}

func (g *Game) shiftRight(array []int) {
  for n := 0; n < g.Size - 1; n++ {
    for i := g.Size - 1; i > 0; i-- {
  //     fmt.Println("i=",i,", i-1=",(i-1))
      if array[i] == 0 {
        array[i], array[i - 1] = array[i - 1], array[i]
      }
    }
  }
}

func (g *Game) shiftLeft(array []int) {  
  for n := 0; n < g.Size - 1; n++ {
    for i := 0; i < g.Size - 1; i++ {
//       fmt.Println("i=",i,", i+1=",(i+1))
      if array[i] == 0 {
        array[i], array[i + 1] = array[i + 1], array[i]
      }
    }
  }
}

func (g *Game) PrettyPrint() {
  fmt.Println("Board:")
  for i := 0; i < g.Size; i++ {
      for j := 0; j < g.Size; j++ {
          fmt.Print(g.Board[i][j], " ")
      }
      fmt.Println()
  } 
//   fmt.Println("Score: ", g.Score, " Moves: ", g.Moves)
}