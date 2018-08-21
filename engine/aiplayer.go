package engine

import "fmt"

// Player ...
type Player struct {
	Name  string
	state Stack
	Score int
}

// Init ...
func (p *Player) Init() {
	g := NewGame(4)
	p.state.Push(g)
}

// Play ...
func (p *Player) Play(strat Strategy, c chan int) {

	finish := false

	for !finish {
		// Generate next step (copy current state)
		current := p.state.Pop()
		next := current.Copy()
		p.state.Push(current)

		current.PrettyPrint()

		var end error
		var ok bool
		// Decide which way to go according to strategy
		nextMove := strat.GetNextMove(&next)
		switch nextMove {
		case 0:
			ok, end = next.GoLeft()
			fmt.Println("Going left")
		case 1:
			ok, end = next.GoRight()
			fmt.Println("Going right")
		case 2:
			ok, end = next.GoUp()
			fmt.Println("Going up")
		default:
			ok, end = next.GoDown()
			fmt.Println("Going down")
		}

		PlaceNewValue(&next)

		if next.IsGameOver() {
			finish = true
		}

		// I can't keep playing
		if !ok {

			// We save the step
			p.state.Push(next)
			p.Score = next.Score

			//next.PrettyPrint()
			switch end.Error() {
			case "movement not possible":
			case "try again":
				continue
			case "can't place another value, the game is over":
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
}
