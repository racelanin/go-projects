package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	for i := 1; i <= 4; i++ {
		fmt.Printf("Player %v: %v, %v, %v\n", i, deck[i*12/4-3], deck[i*12/4-2], deck[i*12/4-1])
	}
}
