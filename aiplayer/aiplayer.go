package aiplayer

// import "fmt"
import "math/rand"
import "time"
import "github.com/LeandroGuillen/2e11/engine"

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

func (p *Player) Play() {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  randomValue := r.Intn(4)
  
  g := p.state.Pop()
  h := engine.CreateGame()
  h = g
  p.state.Push(g)
  
  switch(randomValue) {
    case 0:
      h.GoLeft()
    case 1:
      h.GoRight()
    case 2:
      h.GoUp()
    default:
      h.GoDown()
  }
  
  h.GoRight()
  h.GoUp()
  h.GoDown()
  h.GoLeft()
  h.GoRight()
  h.GoUp()
  h.GoDown()
  h.GoLeft()
  h.GoRight()
  h.GoUp()
  h.GoDown()
  h.GoLeft()
  
  p.state.Push(h)
  p.Score = h.Score
}