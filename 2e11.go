package main

import "fmt"

type Game struct {
  board [][]int
  score int
  moves int
}

func main() {
  
  
  game := initalizeGame(4)
  
  fmt.Println("Game: ", game)
  
  prettyPrint(game, 4)
}

func initalizeGame(N int) Game {

  b := make([][]int, N)
  for i := 0; i < N; i++ {
      b[i] = make([]int, N)
  }
  
  game := Game{board: b, score: 0, moves: 0}
  
  return game
}

func prettyPrint(g Game, N int) {
  for i := 0; i < N; i++ {
      
      for j := 0; j < N; j++ {
          fmt.Print(g.board[i][j], " ")
      }
      fmt.Println()
  } 
}