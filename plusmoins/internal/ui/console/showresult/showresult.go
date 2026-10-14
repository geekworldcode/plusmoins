package ui

import (
	"plusmoins/internal/game"
	"plusmoins/internal/ui/console/color"
	"fmt"
)

func ShowResult(r game.Result, attempts int) {
	switch r {
	case game.TooLow:
		fmt.Println(color.Red + "C'est plus !" + color.Reset)
	case game.TooHigh:
		fmt.Println(color.Red + "C'est moins !" + color.Reset)
	case game.Win:
		fmt.Printf(color.Yellow + "Trouvé en %d essais !\n" + color.Reset, attempts)
	}
}