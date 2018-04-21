package aiplayer

import "fmt"

import "github.com/leandroguillen/2e11/engine"
import "github.com/leandroguillen/2e11/strategy"

// Player ...
type Player struct {
	Name  string
	state Stack
	Score int
}

// Init ...
func (p *Player) Init() {
	g := engine.NewGame(4)
	p.state.Push(g)
}

// Play ...
func (p *Player) Play(strat strategy.Strategy, c chan int) {

	finish := false

	for !finish {
		// Generate next step (copy current state)
		current := p.state.Pop()
		next := current.Copy()
		p.state.Push(current)

		var end error
		var ok bool
		// Decide which way to go according to strategy
		nextMove := strat.GetNextMove(&next)
		switch nextMove {
		case 0:
			ok, end = next.GoLeft()
		case 1:
			ok, end = next.GoRight()
		case 2:
			ok, end = next.GoUp()
		default:
			ok, end = next.GoDown()
		}

		// I can't keep playing
		if !ok {

			// We save the step
			p.state.Push(next)
			p.Score = next.Score

			//next.PrettyPrint()
		}

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

	}
}
