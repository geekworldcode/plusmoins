package ui

import "fmt"

func Welcome(max int) {
	fmt.Printf("Bonjour vous jouez au jeu du plus ou moins ! Vous devez trouver un nombre entre 1 et %d \n", max)
}