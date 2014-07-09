package main

import "fmt"
import "./engine"

func main() {
  
  
  g := engine.NewGame(4)
  
  
  g.PrettyPrint()
  g.GoDown()
  g.PrettyPrint()
  
  fmt.Println(g.Board)
//   fmt.Println("Score: ", g.Score)
}
