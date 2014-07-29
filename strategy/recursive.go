package strategy

import "math"
import "time"
import "fmt"
import "github.com/LeandroGuillen/2e11/engine"

type Recursive struct {
}

func (s Recursive) GetNextMove(g *engine.Game) int {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  
  left    := engine.Game{}
  left    = g
  right   := engine.Game{}
  right   = g
  up      := engine.Game{}
  up      = g
  down    := engine.Game{}
  down    = g
  
  left.GoLeft()
  right.GoLeft()
  up.GoLeft()
  down.GoLeft()
  
  fmt.Println("left.Score:", left.Score)
  fmt.Println("right.Score:", right.Score)
  fmt.Println("up.Score:", up.Score)
  fmt.Println("down.Score:", down.Score)
  
  max := math.Max(math.Max(float64(left.Score),
                           float64(right.Score)),
                  math.Max(float64(up.Score),
                           float64(down.Score)))
  
  fmt.Println("Max score:", max)
  
  return 0
}

func (s Recursive) Name() string {
  return "Recursive"
}