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
	secret := flag.Int("secret", 0, "nombre secret imposé pour les tests (0 : tirage au hasard)")
	recordFile := flag.String("record", "record.txt", "fichier du record")
	flag.Parse()

	if *max < 1 {
		fatal(fmt.Errorf("-max doit valoir au moins 1, reçu %d", *max))
	}
	if *secret < 0 || *secret > *max {
		fatal(fmt.Errorf("-secret doit être compris entre 1 et %d, reçu %d", *max, *secret))
	}

	g := game.New(*max)
	if *secret > 0 {
		g = game.NewWithSecret(*secret)
	}

	if err := play(g, *max, *recordFile); err != nil && !errors.Is(err, parser.ErrQuit) {
		fatal(err)
	}
}

func play(g *game.Game, max int, recordFile string) error {
	best, err := score.Load(recordFile)
	if err != nil {
		return err
	}

	welcomeui.Welcome(max)
	if best > 0 {
		fmt.Printf("Record actuel : %d essais.\n", best)
	}

	for {
		text, err := askui.Ask()
		if errors.Is(err, io.EOF) {
			return parser.ErrQuit
		}
		if err != nil {
			return err
		}

		n, err := parser.ParseGuess(text, max)
		if errors.Is(err, parser.ErrQuit) {
			return err
		}
		if err != nil {
			errorui.ShowError(err)
			continue
		}

		result := g.Guess(n)
		resultui.ShowResult(result, g.Attempts)

		if result == game.Win {
			if best == 0 || g.Attempts < best {
				if err := score.Save(recordFile, g.Attempts); err != nil {
					return err
				}
			}
			return parser.ErrQuit
		}
	}
}

func fatal(err error) {
	fatalui.ShowFatal(err)
	os.Exit(1)
}
