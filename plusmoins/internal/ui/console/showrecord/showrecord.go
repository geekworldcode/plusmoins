package ui

import "fmt"

func ShowRecord(best int) {
	if best == 0 {
		fmt.Println("Pas encore de record.")
		return
	}
	fmt.Printf("Record à battre : %d essais.\n", best)
}
