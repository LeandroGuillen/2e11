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

func CreateGame() Game {
  Size := 4
  b := make([][]int, Size)
  for i := 0; i < Size; i++ {
    b[i] = make([]int, Size)   
  }
  g := Game{Board: b, Score: 0, Moves: 0, Size: Size, Valid: true}
  return g
}

func (g *Game) NewGame() {
  g.ClearBoard()
  // Start off with a couple of cells filled
  g.placeNewValue()
  g.placeNewValue()
}

func (g *Game) placeNewValue() {
  value := getNewValue()  
  x, y := getNewPos(g.Size)
  
  if g.Board[x][y] != 0 {
    for x, y := getNewPos(g.Size); g.Board[x][y] != 0; x, y = getNewPos(g.Size) {      
    }
  } else {
    g.Board[x][y] = value
  }  
}

func getNewPos(size int) (x, y int ) { 
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
  g = Game{Valid: false}
  return
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
