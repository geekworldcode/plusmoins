package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"plusmoins/internal/game"
	"plusmoins/internal/parser"
	"plusmoins/internal/score"
	askui "plusmoins/internal/ui/console/ask"
	errorui "plusmoins/internal/ui/console/showerror"
	fatalui "plusmoins/internal/ui/console/showfatal"
	resultui "plusmoins/internal/ui/console/showresult"
	welcomeui "plusmoins/internal/ui/console/welcome"
)

func main() {
	var max = flag.Int("max", 100, "la borne haute de la valeur à deviner (int)")
	var secret = flag.Int("secret", 0, "nombre secret imposé pour les tests (0 : tirage au hasard)")
	var recordFile = flag.String("record", "record.txt", "fichier du record")
	flag.Parse()
	if *max < 1 || *secret < 0 || *secret > *max {
		quit(fmt.Errorf("options invalides : -max %d, -secret %d", *max, *secret))
	}
	g := game.New(*max)
	if *secret > 0 {
		g = game.NewWithSecret(*secret)
	}
	best, err := score.Load(*recordFile)
	if err != nil {
		quit(err)
	}
	welcomeui.Welcome(*max)
	if best > 0 {
		fmt.Printf("Record actuel : %d essais.\n", best)
	}

	for {
		text, err := askui.Ask()
		if err != nil {
			quit(err)
		}

		n, err := parser.ParseGuess(text, *max)
		if errors.Is(err, parser.ErrQuit) {
			quit(err)
		}
		if err != nil {
			errorui.ShowError(err)
			continue
		}

		result := g.Guess(n)
		resultui.ShowResult(result, g.Attempts)

		if result == game.Win {
			if best == 0 || g.Attempts < best {
				if err := score.Save(*recordFile, g.Attempts); err != nil {
					quit(err)
				}
			}
			quit(parser.ErrQuit)
		}
	}
}

func quit(err error) {
	if errors.Is(err, parser.ErrQuit) || errors.Is(err, io.EOF) {
		os.Exit(0)
	}
	fatalui.ShowFatal(err)
	os.Exit(1)
}
