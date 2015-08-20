package strategy

import "github.com/leandroguillen/2e11/engine"

type Strategy interface {
	GetNextMove(*engine.Game) int
	Name() string
}
