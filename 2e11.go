package main

import "fmt"
import "time"


import "github.com/LeandroGuillen/2e11/aiplayer"
import "github.com/LeandroGuillen/2e11/strategy"

func main() {
  score := 0
  highest := 0
  lowest := 2000000
  average := 0
  games := 1
  
  c := make(chan int)
  
  s := strategy.Random{}
  fmt.Println("Strategy:", s.Name(), "...")

  for i := 0; i < games; i++ {
    go func(c chan int) {
      p := aiplayer.Player{Name: "Paco"}
      p.Init()
      p.Play(s, c)
    }(c)
  }
 
  for x := 0; x < games; x++{
    select {
    case score = <- c:
        
      if score > highest {
        highest = score
      }
      
      if score < lowest {
        lowest = score
      }
      
      average = average + score
      
    case <-time.After(time.Second * 1):
      fmt.Println("I'm done")
      
      break
    }
  }
  
  
  average = average / games
  
  fmt.Println("Total games:", games)
  fmt.Println("Highest:", highest)
  fmt.Println("Lowest:", lowest)
  fmt.Println("Average:", average)
}
