package main

import (
	"errors"
	"flag"
	"plusmoins/internal/game"
	"plusmoins/internal/parser"
	askui "plusmoins/internal/ui/console/ask"
	errorui "plusmoins/internal/ui/console/showerror"
	resultui "plusmoins/internal/ui/console/showresult"
	welcomeui "plusmoins/internal/ui/console/welcome"
)

func main() {
	var max = flag.Int("max", 100, "la borne haute de la valeur à deviner (int)")
	flag.Parse()
	g := game.New(*max)
	welcomeui.Welcome(*max)

	for {
		text, err := askui.Ask()
		if err != nil {
			return // fin de l'entrée
		}

		n, err := parser.ParseGuess(text, *max)
		if errors.Is(err, parser.ErrQuit) {
			return
		}
		if err != nil {
			errorui.ShowError(err)
			continue
		}

		result := g.Guess(n)
		resultui.ShowResult(result, g.Attempts)

		if result == game.Win {
			return
		}
	}
}