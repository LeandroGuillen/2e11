package main

import "fmt"
import "github.com/LeandroGuillen/2e11/aiplayer"
import "github.com/LeandroGuillen/2e11/strategy"

func main() {
  p := aiplayer.Player{Name: "Paco"}

  score := 0
  highest := 0
  lowest := 10000
  average := 0
  games := 10
  
  for i := 0; i < games; i++ {
    p.Init()
    p.Play(strategy.Random{})
    
    score = p.Score
    if score > highest {
      highest = score
    }
    
    if score < lowest {
      lowest = score
    }
    
    average = average + score
  }
  
  average = average / games
  
  fmt.Println("No. games:", games)
  fmt.Println("Highest:", highest)
  fmt.Println("Lowest:", lowest)
  fmt.Println("Average:", average)
}
