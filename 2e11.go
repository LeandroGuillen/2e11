package main

import (
	"fmt"
	"time"

	"github.com/leandroguillen/2e11/engine"
)

func main() {
	score := 0
	highest := 0
	lowest := 2000000
	average := 0
	games := 1

	c := make(chan int)

	s := engine.Recursive{}
	fmt.Println("Strategy:", s.Name())

	// for i := 0; i < games; i++ {
	// 	go func(c chan int) {
	p := engine.Player{Name: "Paco"}
	p.Init()
	p.Play(s, c)
	// 	}(c)
	// }

	for x := 0; x < games; x++ {
		select {
		case score = <-c:

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
