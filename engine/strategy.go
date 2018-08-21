package engine

type Strategy interface {
	GetNextMove(*Game) int
	Name() string
}
