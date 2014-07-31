package strategy

import "math"
import "fmt"
import "github.com/LeandroGuillen/2e11/engine"

type Recursive struct {
}

func (s Recursive) GetNextMove(g *engine.Game) int {  
  left    := engine.Game{}
  left    = *g
  right   := engine.Game{}
  right   = *g
  up      := engine.Game{}
  up      = *g
  down    := engine.Game{}
  down    = *g
  
  left.GoLeft()
  right.GoLeft()
  up.GoLeft()
  down.GoLeft()
  
  ret := 0
  
  maxLR := int(math.Max(float64(left.Score), float64(right.Score)))
  maxUD := int(math.Max(float64(up.Score),   float64(down.Score)))
  
  if left.Score == maxLR {
    ret = 0
  } else {
    ret = 1
  }
  if up.Score == maxUD && up.Score > maxLR {
    ret = 2
  } else if down.Score == maxUD && down.Score > maxLR {
    ret = 3
  }
  
  fmt.Println("maxLR:", maxLR)
  fmt.Println("maxUD:", maxUD)
  
  fmt.Println("left.Score:", left.Score)
  fmt.Println("right.Score:", right.Score)
  fmt.Println("up.Score:", up.Score)
  fmt.Println("down.Score:", down.Score)
  
  
  //fmt.Println("Max score:", max)
  
  return ret
}

func (s Recursive) Name() string {
  return "Recursive"
}