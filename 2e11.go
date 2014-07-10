package main

// import "fmt"
import "github.com/LeandroGuillen/2e11/engine"

func main() {
	//   test()

	g := engine.NewGame(4)

	g.Board[0][0] = 1
	g.Board[1][3] = 6
	g.Board[2][2] = 5
	g.Board[2][3] = 7
	g.Board[3][1] = 8

	g.PrettyPrint()

	g.GoLeft()
	g.PrettyPrint()
	g.GoRight()
	g.PrettyPrint()

	g.GoDown()
	g.PrettyPrint()
	g.GoUp()
	g.PrettyPrint()

	//   fmt.Println(g.Board)
	//   fmt.Println("Score: ", g.Score)
}
