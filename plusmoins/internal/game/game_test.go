package game

import "testing"

func TestGuess(t *testing.T) {
	g := NewWithSecret(42)
	cases := []struct {
		guess int
		want  Result
	}{
		{10, TooLow},
		{99, TooHigh},
		{42, Win},
	}
	for i, c := range cases {
		if got := g.Guess(c.guess); got != c.want {
			t.Errorf("Guess(%d) = %d, attendu %d", c.guess, got, c.want)
		}
		if g.Attempts != i+1 {
			t.Errorf("Attempts = %d, attendu %d", g.Attempts, i+1)
		}
	}
}

func TestNewSecretInRange(t *testing.T) {
	for i := 0; i < 1000; i++ {
		g := New(10)
		if g.secret < 1 || g.secret > 10 {
			t.Fatalf("secret %d hors de 1-10", g.secret)
		}
	}
}
