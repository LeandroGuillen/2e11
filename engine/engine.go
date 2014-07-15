package engine

import "math/rand"
import "time"
import "errors"

func CreateGame() Game {
  Size := 4
  
  b := make([][]int, Size)
  for i := 0; i < Size; i++ {
    b[i] = make([]int, Size)   
  }
    
  g := Game{Board: b, Score: 0, Moves: 0, Size: Size, Valid: true}
  return g
}

func getNewValue() int {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  chance := r.Intn(10)
  value := 0
  if chance <= 1 {
    value = 4
  } else {
    value = 2
  }
  return value
}

func placeNewValue(g *Game) error {
  value := getNewValue()  
  x, y := getNewPos(g.Size)
  
  freeCells := g.freeCells() 
  adjCells := g.twoAdjacentsCells()
  
  if !freeCells {
    if !adjCells {
      return errors.New("I can't place another value, the game is over!")
    } 
    return errors.New("Try again!")
  } else {
    if g.Board[x][y] != 0 {
      for x, y = getNewPos(g.Size); g.Board[x][y] != 0; x, y = getNewPos(g.Size) {}
    } else {
      g.Board[x][y] = value
    } 
  }
  return nil
}

func NullGame() (g Game) {
  g = Game{Valid: false}
  return
}
