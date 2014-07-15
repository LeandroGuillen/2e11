package main

// import "fmt"
import "github.com/LeandroGuillen/2e11/aiplayer"

func main() {
	p := aiplayer.Player{Name: "Paco"}

	p.Init()

	p.Play()
// 	fmt.Println(p.Name, "got", p.Score, "points")
}
