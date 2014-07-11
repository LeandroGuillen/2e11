package engine

import "testing"
import "fmt"

func TestReflect(t *testing.T) {
  g := NewGame(4)
  g.Board[0][0] = 1
  g.Board[1][3] = 6
  g.Board[2][2] = 5
  g.Board[2][3] = 7
  g.Board[3][1] = 8
  
  g.reflect()
  
  e := NewGame(4)
  e.Board[0][0] = 0
  e.Board[0][1] = 0 
  e.Board[0][2] = 0 
  e.Board[0][3] = 1 
  e.Board[1][0] = 6 
  e.Board[1][1] = 0 
  e.Board[1][2] = 0 
  e.Board[1][3] = 0 
  e.Board[2][0] = 7 
  e.Board[2][1] = 5 
  e.Board[2][2] = 0 
  e.Board[2][3] = 0 
  e.Board[3][0] = 0 
  e.Board[3][1] = 0 
  e.Board[3][2] = 8 
  e.Board[3][3] = 0
  
    
  if !g.Equals(&e) {
    fmt.Println("Expected:")
    e.PrettyPrint()
    fmt.Println("Got:")
    g.PrettyPrint()
    t.Error("Boards don't match")
  }
}

func TestTranspose(t *testing.T) {
  g := NewGame(4)
  g.Board[0][0] = 1
  g.Board[1][3] = 6
  g.Board[2][2] = 5
  g.Board[2][3] = 7
  g.Board[3][1] = 8
  
  g.transpose()
  
  e := NewGame(4)
  e.Board[0][0] = 1
  e.Board[0][1] = 0 
  e.Board[0][2] = 0 
  e.Board[0][3] = 0 
  e.Board[1][0] = 0 
  e.Board[1][1] = 0 
  e.Board[1][2] = 0 
  e.Board[1][3] = 8 
  e.Board[2][0] = 0 
  e.Board[2][1] = 0 
  e.Board[2][2] = 5 
  e.Board[2][3] = 0 
  e.Board[3][0] = 0 
  e.Board[3][1] = 6 
  e.Board[3][2] = 7 
  e.Board[3][3] = 0
  
    
  if !g.Equals(&e) {
    fmt.Println("Expected:")
    e.PrettyPrint()
    fmt.Println("Got:")
    g.PrettyPrint()
    t.Error("Boards don't match")
  }
}
