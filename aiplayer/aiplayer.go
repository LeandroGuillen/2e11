package aiplayer

import "fmt"

import "github.com/LeandroGuillen/2e11/engine"
import "github.com/LeandroGuillen/2e11/strategy"

type Player struct {
  Name string
  state Stack
  Score int
}

func (p *Player) Init() {
  g := engine.CreateGame()
  g.NewGame()
  p.state.Push(g)
}

func (p *Player) Play(strat strategy.Strategy, c chan int) {
  
  finish := false
  
  for !finish {
    // Generate next step (copy current state)
    current := p.state.Pop()
    next := engine.Game{}
    next = current
    p.state.Push(current)
    
    var end error
    // Decide which way to go according to strategy
    nextMove := strat.GetNextMove(&next)
    switch(nextMove) {
      case 0:
        end = next.GoLeft()
        fmt.Println("left")
      case 1:
        end = next.GoRight()
        fmt.Println("right")
      case 2:
        end = next.GoUp()
        fmt.Println("up")
      default:
        end = next.GoDown()
        fmt.Println("down")
    }
    
    // I can't keep playing
    if end != nil {
      //fmt.Println(end.Error())
      switch end.Error() {
        case "Movement not possible!":
        case "Try again!":
          continue
        case "I can't place another value, the game is over!":
          //fmt.Println("Final score:", current.Score)
          //fmt.Println("Board:")
          //current.PrettyPrint()
          finish = true
          
          // Send score through channel
          c <- current.Score
          
        default:
          fmt.Println("Unexpected error, finishing game. Score so far:", current.Score)
          finish = true
      }
      
    // We save the step
    } else {
      p.state.Push(next)
      p.Score = next.Score
    
      //next.PrettyPrint()
      //fmt.Println("-------")
    }
  }
}
