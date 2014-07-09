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
  Down
  Left
  Right
)

func NewGame(Size int) Game {
  b := make([][]int, Size)
  for i := 0; i < Size; i++ {
    b[i] = make([]int, Size)   
  }
  
  b[0][0] = 1
  b[0][1] = 2
  b[0][2] = 3
  b[1][0] = 4
  
  g := Game{Board: b, Score: 0, Moves: 0, Size: Size}
  return g
}

func (g *Game) GoUp() {
  g.move(Up)
}

func (g *Game) GoDown() {
  g.move(Down)
}

func (g *Game) GoLeft() {
  g.move(Left)
}

func (g *Game) GoRight() {
  g.move(Right)
}

func (g *Game) boardColumn(columnIndex int) (column []int) {
    column = make([]int, 0)
    for _, row := range g.Board {
        column = append(column, row[columnIndex])
    }
    return
}

func (g *Game) moveColumn(columnIndex int, d direction) []int {
  // Move in the specified direction
  // Only if the cell moving into is INSIDE the board
  // and EMPTY
  
  
  column := g.boardColumn(columnIndex)
  fmt.Println("column BEFORE: ", column)
  
  // Repeat N times
  for r := 0; r < g.Size; r++ {
    // Every cell
    for i := 0; i < g.Size - 1; i++{
      // Move if its appropiate
      if column[i + 1] == 0 && i + 1 < g.Size {
        column[i + 1] = column[i]
        column[i] = 0
      }
    }
  }
  fmt.Println("column AFTER: ", column)
  
  return column
}

func (g *Game) move(d direction) {
  switch d {
    case Down:
       
      // Traverse by columns
      for j := 0; j < g.Size; j++ {
        g.moveColumn(j, d)
      }
    default:
      for i := 0; i < g.Size; i++ {
//         g.Board[:][i] = 9
//         for j := 0; j < g.Size; j++ {
//               g.Board[i][j]
//           }
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
  fmt.Println("Score: ", g.Score, " Moves: ", g.Moves)
}