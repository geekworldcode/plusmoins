package console

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"plusmoins/internal/game"
)

var in = bufio.NewScanner(os.Stdin)

func Welcome(max int) {
	fmt.Printf("J'ai choisi un nombre entre 1 et %d. Tape q pour quitter.\n", max)
}

func ShowRecord(best int) {
	fmt.Printf("Record actuel : %d essais.\n", best)
}

func Ask() (string, error) {
	fmt.Print("Ton nombre : ")
	if in.Scan() {
		return in.Text(), nil
	}
	if err := in.Err(); err != nil {
		return "", err
	}
	return "", io.EOF
}

func ShowResult(r game.Result, attempts int) {
	switch r {
	case game.TooLow:
		fmt.Println("C'est plus !")
	case game.TooHigh:
		fmt.Println("C'est moins !")
	case game.Win:
		fmt.Printf("Trouvé en %d essais !\n", attempts)
	}
}

func ShowError(err error) {
	fmt.Println(err)
}

func ShowFatal(err error) {
	fmt.Fprintln(os.Stderr, "erreur :", err)
}
