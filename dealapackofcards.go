package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	ind := 0
	for j := 0; j < 12; j += 3 {
		ind++
		fmt.Printf("Player %d: %d, %d, %d\n", ind, deck[j], deck[1+j], deck[2+j])
	}
}
