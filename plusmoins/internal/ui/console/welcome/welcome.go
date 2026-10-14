package ui

import "fmt"

func Welcome(max int) {
	fmt.Printf("Bonjour vous jouez au jeu du plus ou moins ! Vous devez un nombre entre 0 et %d \n", max)
}