package strategy

import "github.com/LeandroGuillen/2e11/engine"

type Strategy interface {
    GetNextMove(*engine.Game) int
    Name() string 
}

