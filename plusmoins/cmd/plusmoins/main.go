package main

import (
	"errors"
	"strconv"
	"strings"
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
			parser.ErrQuit
		}

		n, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			errorui.ShowError(errors.New("ce n'est pas un nombre"))
			continue
		}

		result := g.Guess(n)
		resultui.ShowResult(result, g.Attempts)

		if result == game.Win {
			parser.ErrQuit
		}
	}
}