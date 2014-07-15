package engine

import "fmt"
// import "math"

type Game struct {
  Board [][]int
  Score int
  Moves int
  Size int
  Valid bool
}

func (g *Game) NewGame() {
  g.ClearBoard()
  // Start off with a couple of cells filled
  placeNewValue(g)
  placeNewValue(g)
}

func (g *Game) Equals(h *Game) bool {
  if g.Size != h.Size {
    return false
  }
  
  for i := 0; i < g.Size; i++ {
    for j := 0; j < g.Size; j++ {
      if g.Board[i][j] != h.Board[i][j] {
        return false
      }
    }
  } 
  return true
}

func (g *Game) PrettyPrint() {
  for i := 0; i < g.Size; i++ {
      for j := 0; j < g.Size; j++ {
          fmt.Printf("%5d", g.Board[i][j])
      }
      fmt.Println()
  } 
}

func (g *Game) ClearBoard() {
  for i := 0; i < g.Size; i++ {
      for j := 0; j < g.Size; j++ {
          g.Board[i][j] = 0
      }
  } 
}

func (g *Game) freeCells() bool {
  for i := 0; i < g.Size; i++ {
    for j := 0; j < g.Size; j++ {
        if g.Board[i][j] == 0 {
          return true
        }
    }
  }
  return false
}

func (g *Game) hasEqualAdjacent(x, y int) bool {
  
  v := g.Board[x][y]
  
  if x == 0 {
    if y == 0 {
      return v == g.Board[x][y + 1] || v == g.Board[x + 1][y]
    } else if y == g.Size - 1 {
      return v == g.Board[x - 1][y] || v == g.Board[x][y + 1]
    } else {
      return v == g.Board[x][y + 1] || v == g.Board[x][y - 1] || v == g.Board[x + 1][y]
    }
  } else if x == g.Size - 1 {
    if y == 0 {
      return v == g.Board[x][y + 1] || v == g.Board[x - 1][y]
    } else if y == g.Size - 1 {
      return v == g.Board[x - 1][y] || v == g.Board[x][y - 1]
    } else {
      return v == g.Board[x][y + 1] || v == g.Board[x][y - 1] || v == g.Board[x - 1][y]
    }
  } else {
    
    if y == 0 {
      return v == g.Board[x][y + 1] || v == g.Board[x - 1][y] || v == g.Board[x + 1][y]
    } else if y == g.Size - 1 {
      return v == g.Board[x - 1][y] || v == g.Board[x][y - 1] || v == g.Board[x + 1][y]
    } else {
      return v == g.Board[x][y + 1] || v == g.Board[x][y - 1] || v == g.Board[x - 1][y] || v == g.Board[x + 1][y]
    }
  }
}

func (g *Game) areThereAdjCells() bool {
  
  for i := 0; i < g.Size - 1; i++ {
    for j := 0; j < g.Size - 1; j++ {
      if g.hasEqualAdjacent(i, j) {
        return true
      }
    }
  }
  
  return false
}

func (g *Game) twoAdjacentsCells() bool {
  for i := 1; i < g.Size - 1; i++ {
    for j := 1; j < g.Size - 1; j++ {
      a := g.Board[i][j]
      b := g.Board[i][j + 1]
      c := g.Board[i][j - 1]
      d := g.Board[i + 1][j]
      e := g.Board[i - 1][j]
      
      if a == b || a == c || a == d || a == e   {
        return true
      }
    }
  }
  
  a := g.Board[0][0]
  b := g.Board[1][0]
  c := g.Board[2][0]
  d := g.Board[3][0]
  
  if b == a || b == c || c == d {
    return true
  }
  
  e := g.Board[0][3]
  f := g.Board[1][3]
  x := g.Board[2][3]
  y := g.Board[3][3]
     
  if f == e || f == x || x == y {
    return true
  }
  
  a = g.Board[0][0]
  b = g.Board[0][1]
  c = g.Board[0][2]
  d = g.Board[0][3]

  if b == a || b == c || c == d {
    return true
  }
  
  e = g.Board[3][0]
  f = g.Board[3][1]
  x = g.Board[3][2]
  y = g.Board[3][3]
     
  if f == e || f == x || x == y {
    return true
  }
  
  return false
}

func (g *Game) IsGameOver() bool {
  // Any free cells?
  return !g.freeCells() && !g.twoAdjacentsCells()
}