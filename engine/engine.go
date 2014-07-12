package engine

import "fmt"
import "math/rand"
import "time"

type Game struct {
  Board [][]int
  Score int
  Moves int
  Size int
  Valid bool
}

func NewGame() Game {
  Size := 4
  b := make([][]int, Size)
  for i := 0; i < Size; i++ {
    b[i] = make([]int, Size)   
  }
  
  cell1, cell2 := getNewValue(), getNewValue()
  
  x1, y1 := getNewPos(Size)
  x2, y2 := getNewPos(Size)
  
  b[x1][y1] = cell1
  b[x2][y2] = cell2
 
  g := Game{Board: b, Score: 0, Moves: 0, Size: Size, Valid: true}
  return g
}

func (g *Game) PrettyPrint() {
  for i := 0; i < g.Size; i++ {
      for j := 0; j < g.Size; j++ {
          fmt.Print(g.Board[i][j], " ")
      }
      fmt.Println()
  } 
}

func (g *Game) ClearBoard() {
  for i := 0; i < g.Size; i++ {
      for j := 0; j < g.Size; j++ {
          g.Board[i][j] = 0
      }
  } 
}


func getNewPos(size int) (x, y int) { 
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  x = r.Intn(size)
  y = r.Intn(size)
  return
}

func getNewValue() int {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  chance := r.Intn(10)
  value := 0
  if chance <= 2 {
    value = 4
  } else {
    value = 2
  }
  return value
}

func (g *Game) Equals(h *Game) (equal bool) {
  equal = true
  equal = equal && g.Size == h.Size
  
  for i := 0; i < g.Size; i++ {
    for j := 0; j < g.Size; j++ {
      equal = equal && g.Board[i][j] == h.Board[i][j]
    }
  } 
  return
}

func NullGame() (g Game) {
  g = NewGame()
  g.Valid = false
  return
}