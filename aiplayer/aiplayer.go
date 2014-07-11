package aiplayer

// import "fmt"
import "github.com/LeandroGuillen/2e11/engine"

type Player struct {
  name string
  state Stack
}

func (p *Player) Init(name string) {
  p.name = name
  p.state.Push(engine.NewGame())
}

// type Game struct {
//   Board [][]int
//   Score int
//   Moves int
//   Size int
// }
// 
// func NewGame(Size int) Game {
//   b := make([][]int, Size)
//   for i := 0; i < Size; i++ {
//     b[i] = make([]int, Size)   
//   }
//   
//   g := Game{Board: b, Score: 0, Moves: 0, Size: Size}
//   return g
// }
// 
// func (g *Game) PrettyPrint() {
//   for i := 0; i < g.Size; i++ {
//       for j := 0; j < g.Size; j++ {
//           fmt.Print(g.Board[i][j], " ")
//       }
//       fmt.Println()
//   } 
// }
// 
// func (g *Game) Equals(h *Game) (equal bool) {
//   equal = true
//   equal = equal && g.Size == h.Size
//   
//   for i := 0; i < g.Size; i++ {
//     for j := 0; j < g.Size; j++ {
//       equal = equal && g.Board[i][j] == h.Board[i][j]
//     }
//   } 
// 
//   return
// }