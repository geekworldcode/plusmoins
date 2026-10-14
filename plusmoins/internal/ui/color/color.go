package color

import (
	"fmt"
	"os"

	"plusmoins/internal/game"
	"plusmoins/internal/ui/console"
)

const (
	reset  = "\033[0m"
	bold   = "\033[1m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func Welcome(max int) {
	fmt.Printf(bold+"J'ai choisi un nombre entre 1 et %d."+reset+" Tape q pour quitter.\n", max)
}

func ShowRecord(best int) {
	fmt.Printf("Record actuel : "+green+"%d essais"+reset+".\n", best)
}

func Ask() (string, error) {
	return console.Ask()
}

func ShowResult(r game.Result, attempts int) {
	switch r {
	case game.TooLow:
		fmt.Println(yellow + "C'est plus !" + reset)
	case game.TooHigh:
		fmt.Println(yellow + "C'est moins !" + reset)
	case game.Win:
		fmt.Printf(green+bold+"Trouvé en %d essais !"+reset+"\n", attempts)
	}
}

func ShowError(err error) {
	fmt.Println(red + err.Error() + reset)
}

func ShowFatal(err error) {
	fmt.Fprintln(os.Stderr, red+"erreur : "+err.Error()+reset)
}
