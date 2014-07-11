package engine

func (g *Game) GoUp() {
  g.transpose()
  g.reflect()
  for _, val := range g.Board {
    g.shift(val)
  }
  g.reflect()
  g.transpose()
}

func (g *Game) GoDown() {
  g.transpose()
  for _, val := range g.Board {
    g.shift(val)
  }
  g.transpose()
}

func (g *Game) GoLeft() {
  g.reflect()
  for _, val := range g.Board {
    g.shift(val)
  }
  g.reflect()
}

func (g *Game) GoRight() {
  for _, val := range g.Board {
    g.shift(val)
  }
}

func (g *Game) shift(array []int) {
  for n := 0; n < g.Size - 1; n++ {
    for i := g.Size - 1; i > 0; i-- {
      if array[i] == 0 {
        array[i], array[i - 1] = array[i - 1], array[i]
      } else if array[i] == array[i - 1] {
        // Clear the collapsing cell
        array[i - 1] = 0
        // Increase cell value
        array[i] = array[i] + 1
        // Update score and moves
        g.Score = g.Score + array[i]
        g.Moves = g.Moves + 1
      }
    }
  }
}
