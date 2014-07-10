package main

// import "fmt"
import "./engine"

func main() {
  
  
  g := engine.NewGame(4)
  
  
  g.PrettyPrint()
  
   g.GoLeft()
   g.PrettyPrint()
   g.GoRight()
   g.PrettyPrint()
  
  g.GoDown()
  g.PrettyPrint()
  g.GoUp()
  g.PrettyPrint()
  
//   fmt.Println(g.Board)
//   fmt.Println("Score: ", g.Score)
}
