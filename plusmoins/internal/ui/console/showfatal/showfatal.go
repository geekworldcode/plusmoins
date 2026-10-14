package ui

import (
	"fmt"
	"os"
)

func ShowFatal(message error) {
	fmt.Fprintln(os.Stderr, "erreur :", message)
}
