package main

import (
	"errors"
	"flag"
	"os"
	"plusmoins/internal/game"
	"plusmoins/internal/parser"
	"plusmoins/internal/score"
	askui "plusmoins/internal/ui/console/ask"
	errorui "plusmoins/internal/ui/console/showerror"
	fatalui "plusmoins/internal/ui/console/showfatal"
	recordui "plusmoins/internal/ui/console/showrecord"
	resultui "plusmoins/internal/ui/console/showresult"
	welcomeui "plusmoins/internal/ui/console/welcome"
)

func main() {
	var max = flag.Int("max", 100, "la borne haute de la valeur à deviner (int)")
	var secret = flag.Int("secret", 0, "impose le nombre secret, pour les tests (0 : au hasard)")
	flag.Parse()
	best, err := score.Load(scoreFile)
	if err != nil {
		fatal(err)
	}
	g := game.New(*max)
	if *secret != 0 {
		g = game.NewWithSecret(*secret)
	}
	welcomeui.Welcome(*max)
	recordui.ShowRecord(best)

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
			if best == 0 || g.Attempts < best {
				if err := score.Save(scoreFile, g.Attempts); err != nil {
					fatal(err)
				}
			}
			return
		}
	}
}

const scoreFile = "record.txt"

func fatal(err error) {
	fatalui.ShowFatal(err)
	os.Exit(1)
}