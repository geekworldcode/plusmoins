package game

import "math/rand/v2"

type Result int

const (
	TooLow  Result = 0
	TooHigh Result = 1
	Win     Result = 2
)

type Game struct {
	secret   int
	Attempts int
}

func New(max int) *Game {
	return NewWithSecret(rand.IntN(max) + 1)
}

func NewWithSecret(secret int) *Game {
	return &Game{secret: secret}
}

func (g *Game) Guess(n int) Result {
	g.Attempts++
	switch {
	case n < g.secret:
		return TooLow
	case n > g.secret:
		return TooHigh
	default:
		return Win
	}
}
