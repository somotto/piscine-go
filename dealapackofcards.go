package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	numPlayers := 4
	cardsPerPlayer := len(deck) / numPlayers
	for i := 0; i < numPlayers; i++ {
		playerHand := deck[i*cardsPerPlayer : (i+1)*cardsPerPlayer]
		fmt.Printf("Player %d: %v\n", i+1, playerHand)
	}
}
