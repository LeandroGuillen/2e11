package engine

import "testing"
import "fmt"

func TestGoRightNoCollapse(t *testing.T) {
  g := CreateGame()
  g.ClearBoard()
  g.Board[0][0] = 1
  g.Board[0][1] = 2
  g.Board[0][2] = 3
  g.Board[1][0] = 4
  g.Board[1][3] = 6
  g.Board[2][2] = 5
  g.Board[2][3] = 7
  g.Board[3][1] = 8
  
  
  e := CreateGame()
  e.ClearBoard()
  e.Board[0][0] = 0 
  e.Board[0][1] = 1 
  e.Board[0][2] = 2 
  e.Board[0][3] = 3 
  e.Board[1][0] = 0 
  e.Board[1][1] = 0 
  e.Board[1][2] = 4 
  e.Board[1][3] = 6 
  e.Board[2][0] = 0 
  e.Board[2][1] = 0 
  e.Board[2][2] = 5 
  e.Board[2][3] = 7 
  e.Board[3][0] = 0 
  e.Board[3][1] = 0 
  e.Board[3][2] = 0 
  e.Board[3][3] = 8
  
  g.GoRight()
    
  if !g.Equals(&e) {
    fmt.Println("Expected:")
    e.PrettyPrint()
    fmt.Println("Got:")
    g.PrettyPrint()
    t.Error("Boards don't match")
  }
}

func TestGoRight(t *testing.T) {
  g := CreateGame()
  g.ClearBoard()
  g.Board[0][0] = 1
  g.Board[0][1] = 1
  g.Board[0][2] = 3
  g.Board[1][0] = 4
  g.Board[1][3] = 6
  g.Board[2][2] = 7
  g.Board[2][3] = 7
  g.Board[3][1] = 8
  
  e := CreateGame()
  e.ClearBoard()
  e.Board[0][0] = 0 
  e.Board[0][1] = 0 
  e.Board[0][2] = 2 
  e.Board[0][3] = 3
  e.Board[1][0] = 0 
  e.Board[1][1] = 0 
  e.Board[1][2] = 4 
  e.Board[1][3] = 6 
  e.Board[2][0] = 0 
  e.Board[2][1] = 0 
  e.Board[2][2] = 0 
  e.Board[2][3] = 14
  e.Board[3][0] = 0 
  e.Board[3][1] = 0 
  e.Board[3][2] = 0 
  e.Board[3][3] = 8
  
  g.GoRight()
    
  if !g.Equals(&e) {
    fmt.Println("Expected:")
    e.PrettyPrint()
    fmt.Println("Got:")
    g.PrettyPrint()
    t.Error("Boards don't match")
  }
}