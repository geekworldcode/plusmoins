package game

import "math/rand/v2"

type Result int

const (
    TooLow  Result = 0
	TooHigh Result = 1
    Win     Result = 2
)

type Game struct {
	secret	 int
	Attempts int
}

func New(max int) *Game {
	return &Game{secret: rand.IntN(max) + 1}
}

func (g *Game) Guess(n int) Result {
	// À vous :
	g.Attempts += 1
	if n > g.secret {
		return TooHigh
	}
	if n < g.secret {
		return TooLow
	} else {
		return Win
	}
	// 1. compter une tentative de plus
	// 2. comparer n au secret
	// 3. retourner TooLow, TooHigh ou Win
}