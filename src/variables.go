package main

import "fmt"

func basicVariable() {
	// These are some basics type in go
	// bool, int, string, float64
	// And we can initialize the variable in 4 different ways

	card := "Ace of spades"
	var card1 = "Ace of fire"
	var card2 string
	var card3 string = "Queen of ice"

	fmt.Println(card,card1,card2,card3)
}
