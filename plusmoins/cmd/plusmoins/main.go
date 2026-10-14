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
	ui "plusmoins/internal/ui/console"
)

func main() {
	max := flag.Int("max", 100, "borne haute du nombre secret")
	secret := flag.Int("secret", 0, "nombre secret imposé, pour les tests (0 : tirage au hasard)")
	recordFile := flag.String("record", "record.txt", "fichier du record")
	flag.Parse()

	if *max < 1 {
		fatal(fmt.Errorf("-max doit valoir au moins 1, reçu %d", *max))
	}
	if *secret < 0 || *secret > *max {
		fatal(fmt.Errorf("-secret doit être compris entre 1 et %d, reçu %d", *max, *secret))
	}

	best, err := score.Load(*recordFile)
	if err != nil {
		fatal(err)
	}

	var g *game.Game
	if *secret > 0 {
		g = game.NewWithSecret(*secret)
	} else {
		g = game.New(*max)
	}

	ui.Welcome(*max)
	if best > 0 {
		ui.ShowRecord(best)
	}

	for {
		text, err := ui.Ask()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			fatal(err)
		}

		n, err := parser.ParseGuess(text, *max)
		if errors.Is(err, parser.ErrQuit) {
			return
		}
		if err != nil {
			ui.ShowError(err)
			continue
		}

		result := g.Guess(n)
		ui.ShowResult(result, g.Attempts)

		if result == game.Win {
			if best == 0 || g.Attempts < best {
				if err := score.Save(*recordFile, g.Attempts); err != nil {
					fatal(err)
				}
			}
			return
		}
	}
}

func fatal(err error) {
	ui.ShowFatal(err)
	os.Exit(1)
}
